package repository

import (
	"context"

	"br.dev.optimus/hermes/pb"
	"gorm.io/gorm"
)

type HermesRepository interface {
	DepartmentStore(context.Context, *pb.DepartmentRequest) (*pb.DepartmentReply, error)
	DepartmentList(context.Context, *pb.ListRequest) (*pb.ListDepartment, error)
	DocumentTypeStore(context.Context, *pb.DocumentTypeRequest) (*pb.DocumentTypeReply, error)
	DocumentTypeList(context.Context, *pb.ListRequest) (*pb.ListDocumentType, error)
	DocumentStore(context.Context, *pb.DocumentRequest) (*pb.DocumentReply, error)
	DocumentList(context.Context, *pb.ListRequest) (*pb.ListDocument, error)
	DocumentImageStore(context.Context, pb.Hermes_DocumentImageStoreServer) error
}

type HermesRepositoryDB struct {
	HermesRepository
	Reader *gorm.DB
	Writer *gorm.DB
}

func NewHermesRepositoryDB(reader *gorm.DB, writer *gorm.DB) *HermesRepositoryDB {
	return &HermesRepositoryDB{Reader: reader, Writer: writer}
}
