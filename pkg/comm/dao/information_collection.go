package dao

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	db "software_experiment/pkg/comm/database"
	"software_experiment/pkg/web/plugin"
)

type InformationCollection struct {
	gorm.Model
	CollectedName string `gorm:"type:varchar(32);not null"`
	Username      string `gorm:"type:varchar(32);not null"`
	CollectedId   uint   `gorm:"not null"`
}

func init() {
	if !db.SqlDB.HasTable("information_collections") {
		db.SqlDB.CreateTable(&InformationCollection{})
	} else {
		db.SqlDB.AutoMigrate(&InformationCollection{})
	}
}

func GetInformationCollectionById(ctx context.Context, queryMap map[string][]string, unscoped bool) (*InformationCollection, error) {
	var informationCollection InformationCollection
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql, num, err := plugin.ProcessQuery(sql.Table("information_collections"), queryMap, "information_collections")
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
			Information: "information collections not found",
		}
	}
	if unscoped {
		err = sql.First(&informationCollection).Error
	} else {
		err = sql.Unscoped().First(&informationCollection).Error
	}
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return &informationCollection, nil
}

func QueryInformationCollection(ctx context.Context, queryMap map[string][]string) ([]InformationCollection, int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	informationCollections := make([]InformationCollection, 0)
	fmt.Println(queryMap)
	sql, num, err := plugin.ProcessQuery(sql.Table("information_collections"), queryMap, "information_collections")
	if err != nil {
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	sql = sql.Find(&informationCollections)
	err = sql.Error
	if err != nil {
		println(err.Error())
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return informationCollections, num, nil
}

func InsertInformationCollection(ctx context.Context, informationCollection *InformationCollection) (*InformationCollection, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Create(informationCollection)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return informationCollection, nil
}

func DeleteInformationCollection(ctx context.Context, informationCollection *InformationCollection) (int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Delete(informationCollection)
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

func UpdateInformationCollection(ctx context.Context, informationCollection *InformationCollection, updateMap map[string]interface{}) (*InformationCollection, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Model(informationCollection).Updates(updateMap)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return informationCollection, nil
}
