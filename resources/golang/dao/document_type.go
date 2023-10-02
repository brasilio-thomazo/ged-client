package dao

import (
	"context"

	"br.dev.optimus/duat/model"
	"gorm.io/gorm"
)

type DocumentTypeDAO struct {
	reader *gorm.DB
	writer *gorm.DB
}

func NewDocumentTypeDAO(reader *gorm.DB, writer *gorm.DB) *DocumentTypeDAO {
	return &DocumentTypeDAO{reader: reader, writer: writer}
}

func (r *DocumentTypeDAO) FindById(ctx context.Context, ID uint64) (*model.DocumentType, error) {
	data := &model.DocumentType{}
	result := r.reader.WithContext(ctx).Where("id = ?", ID).First(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}
