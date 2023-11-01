package service

import (
	"context"
	"log"

	"br.dev.optimus/hermes/pb"
	"br.dev.optimus/hermes/repository"
	"gorm.io/gorm"
)

type HermesService struct {
	pb.UnimplementedHermesServer
	repository repository.HermesRepository
}

func NewHermesService(reader *gorm.DB, writer *gorm.DB) *HermesService {
	return &HermesService{repository: repository.NewHermesRepositoryDB(reader, writer)}
}

func (s *HermesService) DocumentStore(ctx context.Context, in *pb.DocumentRequest) (*pb.DocumentReply, error) {
	log.Printf("receive document::strore::%s\n", in.Code)
	reply, err := s.repository.DocumentStore(ctx, in)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *HermesService) DocumentImageStore(stream pb.Hermes_DocumentImageStoreServer) error {
	log.Printf("receive document::image::strore")
	if err := s.repository.DocumentImageStore(context.Background(), stream); err != nil {
		return err
	}
	return nil
}

func (s *HermesService) DocumentList(ctx context.Context, in *pb.ListRequest) (*pb.ListDocument, error) {
	log.Printf("receive Department::All\n")
	reply, err := s.repository.DocumentList(ctx, in)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *HermesService) DepartmentList(ctx context.Context, in *pb.ListRequest) (*pb.ListDepartment, error) {
	log.Printf("receive Department::All\n")
	reply, err := s.repository.DepartmentList(ctx, in)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *HermesService) DocumentTypeList(ctx context.Context, in *pb.ListRequest) (*pb.ListDocumentType, error) {
	log.Printf("receive Department::All\n")
	reply, err := s.repository.DocumentTypeList(ctx, in)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
