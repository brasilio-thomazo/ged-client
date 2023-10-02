package model

type Department struct {
	ID        uint64  `gorm:"primaryKey;autoIncrement" db:"id" json:"id"`
	Name      string  `gorm:"unique" db:"name" json:"name"`
	CreatedAt uint64  `gorm:"autoCreateTime:milli" db:"created_at" json:"created_at"`
	UpdatedAt uint64  `gorm:"autoUpdateTime:milli" db:"updated_at" json:"updated_at"`
	DeletedAt *uint64 `db:"deleted_at" json:"deleted_at"`
}
