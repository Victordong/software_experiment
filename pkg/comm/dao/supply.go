package dao

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	db "software_experiment/pkg/comm/database"
	"software_experiment/pkg/web/plugin"
	"time"
)

type Supply struct {
	gorm.Model
	Username   string    `gorm:"type:varchar(32);not null"`
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

func GetSupplyById(ctx context.Context, queryMap map[string][]string, unscoped bool) (*Supply, error) {
	var supply Supply
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql, num, err := plugin.ProcessQuery(sql.Table("supplies"), queryMap, "supplies")
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	if num == 0 {
		return nil, plugin.CustomErr{
			Code:        404,
			StatusCode:  404,
			Information: "supplies not found",
		}
	}
	if unscoped {
		err = sql.First(&supply).Error
	} else {
		err = sql.Unscoped().First(&supply).Error
	}
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return &supply, nil
}

func QuerySupply(ctx context.Context, queryMap map[string][]string) ([]Supply, int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	supplys := make([]Supply, 0)
	fmt.Println(queryMap)
	sql, num, err := plugin.ProcessQuery(sql.Table("supplies"), queryMap, "supplies")
	if err != nil {
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	sql = sql.Find(&supplys)
	err = sql.Error
	if err != nil {
		println(err.Error())
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return supplys, num, nil
}

func InsertSupply(ctx context.Context, supply *Supply) (*Supply, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Create(supply)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return supply, nil
}

func DeleteSupply(ctx context.Context, supply *Supply) (int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Delete(supply)
	err := sql.Error
	if err != nil {
		return 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	num := sql.RowsAffected
	return num, nil
}

func UpdateSupply(ctx context.Context, supply *Supply, updateMap map[string]interface{}) (*Supply, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Model(supply).Updates(updateMap)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return supply, nil
}
