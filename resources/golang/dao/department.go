package dao

import (
	"context"

	"br.dev.optimus/duat/model"
	"gorm.io/gorm"
)

type DepartmentDAO struct {
	reader *gorm.DB
	writer *gorm.DB
}

func NewDepartmentDAO(reader *gorm.DB, writer *gorm.DB) *DepartmentDAO {
	return &DepartmentDAO{reader: reader, writer: writer}
}

func (r *DepartmentDAO) FindById(ctx context.Context, ID uint64) (*model.Department, error) {
	data := &model.Department{}
	result := r.reader.WithContext(ctx).Where("id = ?", ID).First(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}
