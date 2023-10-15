package service

import (
	"context"
	"log"

	"br.dev.optimus/duat/pb"
	"br.dev.optimus/duat/repository"
	"gorm.io/gorm"
)

type DocumentTypeService struct {
	pb.UnimplementedDocumentTypeServiceServer
	repository repository.DocumentTypeRepository
}

func NewDocumentTypeService(reader *gorm.DB, writer *gorm.DB) *DocumentTypeService {
	return &DocumentTypeService{repository: repository.NewDocumentTypeRepositoryDB(reader, writer)}
}

func (s *DocumentTypeService) Store(ctx context.Context, in *pb.DocumentTypeRequest) (*pb.DocumentTypeReply, error) {
	log.Printf("receive DocumentType::Store\n")
	reply, err := s.repository.Store(ctx, in)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *DocumentTypeService) All(ctx context.Context, in *pb.EmptyRequest) (*pb.DocumentTypeListReply, error) {
	log.Printf("receive DocumentType::All\n")
	reply, err := s.repository.All(ctx, in)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
