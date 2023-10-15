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

type DepartmentRepository interface {
	Store(ctx context.Context, in *pb.DepartmentRequest) (*pb.DepartmentReply, error)
	All(ctx context.Context, in *pb.EmptyRequest) (*pb.DepartmentListReply, error)
}

type DepartmentRepositoryDB struct {
	DepartmentRepository
	dataDAO *dao.DepartmentDAO
}

func NewDepartmentRepositoryDB(reader *gorm.DB, writer *gorm.DB) *DepartmentRepositoryDB {
	return &DepartmentRepositoryDB{
		dataDAO: dao.NewDepartmentDAO(reader, writer),
	}
}

func (r *DepartmentRepositoryDB) Store(ctx context.Context, in *pb.DepartmentRequest) (*pb.DepartmentReply, error) {
	if _, err := r.dataDAO.FindByName(ctx, in.Name); err == nil {
		return nil, status.Error(codes.AlreadyExists, "department name exists")
	}

	data := &model.Department{}
	if err := copier.Copy(data, in); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result, err := r.dataDAO.Store(ctx, data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	reply := &pb.DepartmentReply{}
	if err := copier.Copy(reply, result); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return reply, nil
}

func (r *DepartmentRepositoryDB) All(ctx context.Context, in *pb.EmptyRequest) (*pb.DepartmentListReply, error) {
	list, err := r.dataDAO.GetAll(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	reply := []*pb.DepartmentReply{}
	if err := copier.Copy(reply, &list); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.DepartmentListReply{List: reply}, nil
}
