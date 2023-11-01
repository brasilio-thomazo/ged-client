package model

import "github.com/google/uuid"

type DocumentImage struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" db:"id" json:"id"`
	DocumentID uuid.UUID `db:"document_id" json:"document_id"`
	Disk       string    `db:"disk" json:"disk"`
	Filename   string    `db:"filename" json:"filename"`
	Pages      uint32    `db:"pages" json:"pages"`
	CreatedAt  uint64    `gorm:"autoCreateTime:milli" db:"created_at" json:"created_at"`
	UpdatedAt  uint64    `gorm:"autoUpdateTime:milli" db:"updated_at" json:"updated_at"`
	DeletedAt  *uint64   `db:"deleted_at" json:"deleted_at"`
}
