package repository

import (
	"context"

	"br.dev.optimus/hermes/dao"
	"br.dev.optimus/hermes/model"
	"br.dev.optimus/hermes/pb"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *HermesRepositoryDB) DepartmentStore(ctx context.Context, in *pb.DepartmentRequest) (*pb.DepartmentReply, error) {
	departmentDAO := dao.NewDepartmentDAO(r.Reader, r.Writer)
	if _, err := departmentDAO.FindByName(ctx, in.Name); err == nil {
		return nil, status.Error(codes.AlreadyExists, "department name exists")
	}

	data := &model.Department{}
	if err := copier.Copy(data, in); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result, err := departmentDAO.Store(ctx, data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	reply := &pb.DepartmentReply{}
	if err := copier.Copy(reply, result); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return reply, nil
}

func (r *HermesRepositoryDB) DepartmentList(ctx context.Context, in *pb.ListRequest) (*pb.ListDepartment, error) {
	var reply []*pb.DepartmentReply
	departmentDAO := dao.NewDepartmentDAO(r.Reader, r.Writer)
	list, err := departmentDAO.GetAll(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := copier.Copy(&reply, list); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ListDepartment{Data: reply}, nil
}
