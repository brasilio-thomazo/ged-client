package repository

import (
	"context"

	"br.dev.optimus/hermes/dao"
	"br.dev.optimus/hermes/model"
	"br.dev.optimus/hermes/pb"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *HermesRepositoryDB) DocumentStore(ctx context.Context, in *pb.DocumentRequest) (*pb.DocumentReply, error) {
	documentTypeDAO := dao.NewDocumentTypeDAO(r.Reader, r.Writer)
	departmentDAO := dao.NewDepartmentDAO(r.Reader, r.Writer)
	documentDAO := dao.NewDocumentDAO(r.Reader, r.Writer)

	if _, err := departmentDAO.FindById(ctx, in.DepartmentId); err != nil {
		return nil, status.Error(codes.InvalidArgument, "department id not exists")
	}

	if _, err := documentTypeDAO.FindById(ctx, in.DocumentTypeId); err != nil {
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
	result, err := documentDAO.Store(ctx, data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	reply := &pb.DocumentReply{}
	if err := copier.Copy(reply, result); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return reply, nil
}

func (r *HermesRepositoryDB) DocumentList(ctx context.Context, in *pb.ListRequest) (*pb.ListDocument, error) {
	var reply []*pb.DocumentReply
	documentDAO := dao.NewDocumentDAO(r.Reader, r.Writer)
	list, err := documentDAO.GetAll(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := copier.Copy(&reply, list); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ListDocument{Data: reply}, nil
}
