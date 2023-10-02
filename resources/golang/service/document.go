package service

import (
	"context"
	"log"

	"br.dev.optimus/duat/pb"
	"br.dev.optimus/duat/repository"
	"gorm.io/gorm"
)

type DocumentService struct {
	pb.UnimplementedDocumentServiceServer
	repository repository.DocumentRepository
}

func NewDocumentService(reader *gorm.DB, writer *gorm.DB) *DocumentService {
	return &DocumentService{repository: repository.NewDocumentRepositoryDB(reader, writer)}
}

func (s *DocumentService) Store(ctx context.Context, in *pb.DocumentRequest) (*pb.DocumentReply, error) {
	log.Printf("receive Document::Store\n")
	reply, err := s.repository.Store(ctx, in)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
