package dao

import (
	db "auto_fertilizer_back/pkg/comm/database"
	"github.com/jinzhu/gorm"
	"time"
)

type Information struct {
	gorm.Model
	CreateBy   string    `gorm:"type:varchar(32);not null"`
	Content    string    `gorm:"type:Text"`
	Name       string    `gorm:"type:varchar(32);not null;unique"`
	VisitNum   int       `gorm:"default=0"`
	ExpiryDate time.Time `gorm:"type:varchar(16);not null"`
}

func init() {
	if !db.SqlDB.HasTable("information") {
		db.SqlDB.CreateTable(&Information{})
	} else {
		db.SqlDB.AutoMigrate(&Information{})
	}
}
