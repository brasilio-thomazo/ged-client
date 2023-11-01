package dao

import (
	"context"

	"br.dev.optimus/hermes/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DocumentDAO struct {
	reader *gorm.DB
	writer *gorm.DB
}

func NewDocumentDAO(reader *gorm.DB, writer *gorm.DB) *DocumentDAO {
	return &DocumentDAO{reader: reader, writer: writer}
}

func (r *DocumentDAO) FindById(ctx context.Context, ID uuid.UUID) (*model.Document, error) {
	data := &model.Document{}
	result := r.reader.WithContext(ctx).Where("id = ?", ID).First(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func (r *DocumentDAO) Store(ctx context.Context, data *model.Document) (*model.Document, error) {
	result := r.writer.WithContext(ctx).Save(data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func (r *DocumentDAO) GetAll(ctx context.Context) ([]*model.Document, error) {
	var list []*model.Document
	result := r.reader.WithContext(ctx).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
