package model

import "github.com/google/uuid"

type Document struct {
	ID             uuid.UUID `gorm:"primaryKey" db:"id" json:"id"`
	DocumentTypeID uint64    `db:"document_type_id" json:"document_type_id"`
	DepartmentID   uint64    `db:"department_id" json:"department_id"`
	Code           string    `gorm:"unique" db:"code" json:"code"`
	Identity       string    `db:"identity" json:"identity"`
	Name           string    `db:"name" json:"name"`
	Comment        *string   `db:"comment" json:"comment"`
	Storage        *string   `db:"storage" json:"storage"`
	DateDocument   string    `db:"date_document" json:"date_document"`
	CreatedAt      uint64    `gorm:"autoCreateTime:milli" db:"created_at" json:"created_at"`
	UpdatedAt      uint64    `gorm:"autoUpdateTime:milli" db:"updated_at" json:"updated_at"`
	DeletedAt      *uint64   `db:"deleted_at" json:"deleted_at"`
}
