package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	db "software_experiment/pkg/comm/database"
	"software_experiment/pkg/web/plugin"
)

type Information struct {
	gorm.Model
	Username string `gorm:"type:varchar(32);not null"`
	Content  string `gorm:"type:Text"`
	Name     string `gorm:"type:varchar(32);not null;unique"`
	VisitNum int    `gorm:"default=0"`
}

func init() {
	if !db.SqlDB.HasTable("informations") {
		db.SqlDB.CreateTable(&Information{})
	} else {
		db.SqlDB.AutoMigrate(&Information{})
	}
}

func GetInformationById(ctx context.Context, queryMap map[string][]string, unscoped bool) (*Information, error) {
	var supplyCollection Information
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql, num, err := plugin.ProcessQuery(sql.Table("informations"), queryMap, "informations")
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
			Information: "informations not found",
		}
	}
	if unscoped {
		err = sql.First(&supplyCollection).Error
	} else {
		err = sql.Unscoped().First(&supplyCollection).Error
	}
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return &supplyCollection, nil
}

func QueryInformation(ctx context.Context, queryMap map[string][]string) ([]Information, int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	supplyCollections := make([]Information, 0)
	sql, num, err := plugin.ProcessQuery(sql.Table("informations"), queryMap, "informations")
	if err != nil {
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	sql = sql.Find(&supplyCollections)
	err = sql.Error
	if err != nil {
		println(err.Error())
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return supplyCollections, num, nil
}

func InsertInformation(ctx context.Context, supplyCollection *Information) (*Information, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Create(supplyCollection)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return supplyCollection, nil
}

func DeleteInformation(ctx context.Context, supplyCollection *Information) (int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Delete(supplyCollection)
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

func UpdateInformation(ctx context.Context, supplyCollection *Information, updateMap map[string]interface{}) (*Information, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Model(supplyCollection).Updates(updateMap)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return supplyCollection, nil
}
