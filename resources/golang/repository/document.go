package repository

import (
	"context"

	"br.dev.optimus/duat/dao"
	"br.dev.optimus/duat/model"
	"br.dev.optimus/duat/pb"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type DocumentRepository interface {
	Store(ctx context.Context, in *pb.DocumentRequest) (*pb.DocumentReply, error)
}

type DocumentRepositoryDB struct {
	DocumentRepository
	dataDAO         *dao.DocumentDAO
	departmentDAO   *dao.DepartmentDAO
	documentTypeDAO *dao.DocumentTypeDAO
}

func NewDocumentRepositoryDB(reader *gorm.DB, writer *gorm.DB) *DocumentRepositoryDB {
	return &DocumentRepositoryDB{
		dataDAO:         dao.NewDocumentDAO(reader, writer),
		departmentDAO:   dao.NewDepartmentDAO(reader, writer),
		documentTypeDAO: dao.NewDocumentTypeDAO(reader, writer),
	}
}

func (r *DocumentRepositoryDB) Store(ctx context.Context, in *pb.DocumentRequest) (*pb.DocumentReply, error) {
	if _, err := r.departmentDAO.FindById(ctx, in.DepartmentId); err != nil {
		return nil, status.Error(codes.InvalidArgument, "department id not exists")
	}

	if _, err := r.documentTypeDAO.FindById(ctx, in.DocumentTypeId); err != nil {
		return nil, status.Error(codes.InvalidArgument, "document type id not exists")
	}

	data := &model.Document{}
	if err := copier.Copy(data, in); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data.ID = id
	result, err := r.dataDAO.Store(ctx, data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	reply := &pb.DocumentReply{}
	if err := copier.Copy(reply, result); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return reply, nil
}
