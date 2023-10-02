package service

import (
	"context"
	"log"

	"br.dev.optimus/duat/pb"
	"br.dev.optimus/duat/repository"
	"gorm.io/gorm"
)

type DocumentImageService struct {
	pb.UnimplementedDocumentImageServiceServer
	repository repository.DocumentImageRepository
}

func NewDocumentImageService(reader *gorm.DB, writer *gorm.DB) *DocumentImageService {
	return &DocumentImageService{repository: repository.NewDocumentImageRepositoryDB(reader, writer)}
}

func (s *DocumentImageService) Store(stream pb.DocumentImageService_StoreServer) error {
	log.Printf("receive DocumentImage::Store\n")
	if err := s.repository.Store(context.Background(), stream); err != nil {
		return err
	}
	return nil
}
