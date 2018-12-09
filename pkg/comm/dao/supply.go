package dao

import (
	db "auto_fertilizer_back/pkg/comm/database"
	"github.com/jinzhu/gorm"
	"time"
)

type Supply struct {
	gorm.Model
	CreateBy   string    `gorm:"type:varchar(32);not null"`
	Name       string    `gorm:"type:varchar(32);not null"`
	Content    string    `gorm:"type:Text"`
	VisitNum   int       `gorm:"default=0"`
	Type       int       `gorm:"not null"`
	ExpiryDate time.Time `gorm:"type:varchar(16);not null"`
}

func init() {
	if !db.SqlDB.HasTable("supplies") {
		db.SqlDB.CreateTable(&Supply{})
	} else {
		db.SqlDB.AutoMigrate(&Supply{})
	}
}
