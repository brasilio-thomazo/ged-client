package dao

import (
	"context"

	"br.dev.optimus/hermes/model"
	"gorm.io/gorm"
)

type DocumentImageDAO struct {
	reader *gorm.DB
	writer *gorm.DB
}

func NewDocumentImageDAO(reader *gorm.DB, writer *gorm.DB) *DocumentImageDAO {
	return &DocumentImageDAO{reader: reader, writer: writer}
}

func (r *DocumentImageDAO) FindById(ctx context.Context, ID uint64) (*model.DocumentImage, error) {
	data := &model.DocumentImage{}
	result := r.reader.WithContext(ctx).Where("id = ?", ID).First(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func (r *DocumentImageDAO) Store(ctx context.Context, data *model.DocumentImage) (*model.DocumentImage, error) {
	result := r.writer.WithContext(ctx).Save(data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil

}
