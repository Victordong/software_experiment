package dao

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	db "software_experiment/pkg/comm/database"
	"software_experiment/pkg/web/plugin"
)

type SupplyCollection struct {
	gorm.Model
	CollectedName string `gorm:"type:varchar(32);not null"`
	Username      string `gorm:"type:varchar(32);not null"`
	CollectedId   uint   `gorm:"not null"`
}

func init() {
	if !db.SqlDB.HasTable("supply_collections") {
		db.SqlDB.CreateTable(&SupplyCollection{})
	} else {
		db.SqlDB.AutoMigrate(&SupplyCollection{})
	}
}

func GetSupplyCollectionById(ctx context.Context, queryMap map[string][]string, unscoped bool) (*SupplyCollection, error) {
	var supplyCollection SupplyCollection
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql, num, err := plugin.ProcessQuery(sql.Table("supply_collections"), queryMap, "supply_collections")
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
			Information: "supply collections not found",
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

func QuerySupplyCollection(ctx context.Context, queryMap map[string][]string) ([]SupplyCollection, int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	supplyCollections := make([]SupplyCollection, 0)
	fmt.Println(queryMap)
	sql, num, err := plugin.ProcessQuery(sql.Table("supply_collections"), queryMap, "supply_collections")
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

func InsertSupplyCollection(ctx context.Context, supplyCollection *SupplyCollection) (*SupplyCollection, error) {
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

func DeleteSupplyCollection(ctx context.Context, supplyCollection *SupplyCollection) (int64, error) {
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

func UpdateSupplyCollection(ctx context.Context, supplyCollection *SupplyCollection, updateMap map[string]interface{}) (*SupplyCollection, error) {
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
