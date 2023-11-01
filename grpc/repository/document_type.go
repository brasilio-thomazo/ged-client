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

func (r *HermesRepositoryDB) DocumentTypeStore(ctx context.Context, in *pb.DocumentTypeRequest) (*pb.DocumentTypeReply, error) {
	documentTypeDAO := dao.NewDocumentTypeDAO(r.Reader, r.Writer)
	if _, err := documentTypeDAO.FindByName(ctx, in.Name); err == nil {
		return nil, status.Error(codes.AlreadyExists, "department name exists")
	}

	data := &model.DocumentType{}
	if err := copier.Copy(data, in); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result, err := documentTypeDAO.Store(ctx, data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	reply := &pb.DocumentTypeReply{}
	if err := copier.Copy(reply, result); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return reply, nil
}

func (r *HermesRepositoryDB) DocumentTypeList(ctx context.Context, in *pb.ListRequest) (*pb.ListDocumentType, error) {
	var reply []*pb.DocumentTypeReply
	documentTypeDAO := dao.NewDocumentTypeDAO(r.Reader, r.Writer)
	list, err := documentTypeDAO.GetAll(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := copier.Copy(&reply, list); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ListDocumentType{Data: reply}, nil
}
