package paste

import (
	"gorm.io/gorm"
)

type Paste struct {
	gorm.Model
	ID   string `gorm:"primaryKey"`
	Text string
}
