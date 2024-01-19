package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"ID"` // Primary key ID
	CreatedAt time.Time      // Creation time
	UpdatedAt time.Time      // Update time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Deletion time
}
