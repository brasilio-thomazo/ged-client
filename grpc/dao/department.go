package dao

import (
	"context"

	"br.dev.optimus/hermes/model"
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

func (r *DepartmentDAO) FindByName(ctx context.Context, name string) (*model.Department, error) {
	data := &model.Department{}
	result := r.reader.WithContext(ctx).Where("name = ?", name).First(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func (r *DepartmentDAO) Store(ctx context.Context, data *model.Department) (*model.Department, error) {
	result := r.writer.WithContext(ctx).Save(data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func (r *DepartmentDAO) GetAll(ctx context.Context) ([]*model.Department, error) {
	var list []*model.Department
	result := r.reader.WithContext(ctx).Where("name != ?", "System").Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
