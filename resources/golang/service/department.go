package service

import (
	"context"
	"log"

	"br.dev.optimus/duat/pb"
	"br.dev.optimus/duat/repository"
	"gorm.io/gorm"
)

type DepartmentService struct {
	pb.UnimplementedDepartmentServiceServer
	repository repository.DepartmentRepository
}

func NewDepartmentService(reader *gorm.DB, writer *gorm.DB) *DepartmentService {
	return &DepartmentService{repository: repository.NewDepartmentRepositoryDB(reader, writer)}
}

func (s *DepartmentService) Store(ctx context.Context, in *pb.DepartmentRequest) (*pb.DepartmentReply, error) {
	log.Printf("receive Department::Store\n")
	reply, err := s.repository.Store(ctx, in)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *DepartmentService) All(ctx context.Context, in *pb.EmptyRequest) (*pb.DepartmentListReply, error) {
	log.Printf("receive Department::All\n")
	reply, err := s.repository.All(ctx, in)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
