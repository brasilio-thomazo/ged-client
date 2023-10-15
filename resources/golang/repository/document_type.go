package repository

import (
	"context"

	"br.dev.optimus/duat/dao"
	"br.dev.optimus/duat/model"
	"br.dev.optimus/duat/pb"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type DocumentTypeRepository interface {
	Store(ctx context.Context, in *pb.DocumentTypeRequest) (*pb.DocumentTypeReply, error)
	All(ctx context.Context, in *pb.EmptyRequest) (*pb.DocumentTypeListReply, error)
}

type DocumentTypeRepositoryDB struct {
	DocumentTypeRepository
	dataDAO *dao.DocumentTypeDAO
}

func NewDocumentTypeRepositoryDB(reader *gorm.DB, writer *gorm.DB) *DocumentTypeRepositoryDB {
	return &DocumentTypeRepositoryDB{
		dataDAO: dao.NewDocumentTypeDAO(reader, writer),
	}
}

func (r *DocumentTypeRepositoryDB) Store(ctx context.Context, in *pb.DocumentTypeRequest) (*pb.DocumentTypeReply, error) {
	if _, err := r.dataDAO.FindByName(ctx, in.Name); err == nil {
		return nil, status.Error(codes.AlreadyExists, "department name exists")
	}

	data := &model.DocumentType{}
	if err := copier.Copy(data, in); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result, err := r.dataDAO.Store(ctx, data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	reply := &pb.DocumentTypeReply{}
	if err := copier.Copy(reply, result); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return reply, nil
}

func (r *DocumentTypeRepositoryDB) All(ctx context.Context, in *pb.EmptyRequest) (*pb.DocumentTypeListReply, error) {
	list, err := r.dataDAO.GetAll(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	reply := []*pb.DocumentTypeReply{}
	if err := copier.Copy(reply, &list); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.DocumentTypeListReply{List: reply}, nil
}
