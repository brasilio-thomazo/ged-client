package model

import "github.com/google/uuid"

type DocumentImage struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" db:"id" json:"id"`
	DocumentID  uuid.UUID `db:"document_id" json:"document_id"`
	Filename    string    `db:"filename" json:"filename"`
	Page        uint32    `db:"page" json:"page"`
	StorageType int       `db:"storage_type" json:"storage_type"`
	CreatedAt   uint64    `gorm:"autoCreateTime:milli" db:"created_at" json:"created_at"`
	UpdatedAt   uint64    `gorm:"autoUpdateTime:milli" db:"updated_at" json:"updated_at"`
	DeletedAt   *uint64   `db:"deleted_at" json:"deleted_at"`
}
