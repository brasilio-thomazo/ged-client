package repository

import (
	"context"
	"fmt"
	"io"
	"os"

	"br.dev.optimus/duat/dao"
	"br.dev.optimus/duat/model"
	"br.dev.optimus/duat/pb"
	"br.dev.optimus/duat/utils"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type DocumentImageRepository interface {
	Store(context.Context, pb.DocumentImageService_StoreServer) error
}

type DocumentImageRepositoryDB struct {
	DocumentImageRepository
	dataDAO     *dao.DocumentImageDAO
	documentDAO *dao.DocumentDAO
}

func NewDocumentImageRepositoryDB(reader *gorm.DB, writer *gorm.DB) *DocumentImageRepositoryDB {
	return &DocumentImageRepositoryDB{
		dataDAO:     dao.NewDocumentImageDAO(reader, writer),
		documentDAO: dao.NewDocumentDAO(reader, writer),
	}
}

func (r *DocumentImageRepositoryDB) Store(ctx context.Context, stream pb.DocumentImageService_StoreServer) error {
	in, err := stream.Recv()
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	documentID, err := uuid.Parse(in.DocumentId)
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := r.documentDAO.FindById(ctx, documentID); err != nil {
		return status.Error(codes.InvalidArgument, "document id not exists")
	}

	data := &model.DocumentImage{}
	if err := copier.Copy(data, in); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	uploadPath := fmt.Sprintf("%s/%s", os.Getenv("UPLOAD_IMAGE"), in.DocumentId)
	file, err := utils.NewFile(fmt.Sprintf("%s%s", id.String(), in.ImageExt), uploadPath)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if err := file.Create(); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	_, err = writeData(file, in.Data)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	image := &model.DocumentImage{
		DocumentID:  documentID,
		Filename:    file.Filename,
		Page:        in.Page,
		StorageType: int(in.StorageType),
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if _, err := writeData(file, in.Data); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	result, err := r.dataDAO.Store(ctx, image)
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	reply := &pb.DocumentImageReply{}
	if err := copier.Copy(reply, result); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if err := stream.SendAndClose(reply); err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func writeData(file *utils.File, data []byte) (int, error) {
	write, err := file.File.Write(data)
	if err != nil {
		return 0, err
	}
	return write, nil
}
