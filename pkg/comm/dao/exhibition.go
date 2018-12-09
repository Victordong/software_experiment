package dao

import (
	db "auto_fertilizer_back/pkg/comm/database"
	"github.com/jinzhu/gorm"
)

type Exhibition struct {
	gorm.Model
	Username string `gorm:"type:varchar(32);not null"`
	Content  string `gorm:"type:Text"`
	Theme    string `gorm:"type:Text"`
	Name     string `gorm:"type:varchar(32);not null;unique"`
	Type     int    `gorm:"not null"`
	VisitNum int    `gorm:"default=0"`
}

func init() {
	if !db.SqlDB.HasTable("exhibition") {
		db.SqlDB.CreateTable(&Exhibition{})
	} else {
		db.SqlDB.AutoMigrate(&Exhibition{})
	}
}
