package dao

import (
	"context"

	"br.dev.optimus/hermes/model"
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

func (r *DocumentTypeDAO) Store(ctx context.Context, data *model.DocumentType) (*model.DocumentType, error) {
	result := r.writer.WithContext(ctx).Save(data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func (r *DocumentTypeDAO) FindByName(ctx context.Context, name string) (*model.DocumentType, error) {
	data := &model.DocumentType{}
	result := r.reader.WithContext(ctx).Where("name = ?", name).First(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func (r *DocumentTypeDAO) GetAll(ctx context.Context) ([]*model.DocumentType, error) {
	var list []*model.DocumentType
	result := r.reader.WithContext(ctx).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
