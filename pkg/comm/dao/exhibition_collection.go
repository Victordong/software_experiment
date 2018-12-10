package dao

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	db "software_experiment/pkg/comm/database"
	"software_experiment/pkg/web/plugin"
)

type ExhibitionCollection struct {
	gorm.Model
	CollectedName string `gorm:"type:varchar(32);not null"`
	Username      string `gorm:"type:varchar(32);not null"`
	CollectedId   uint   `gorm:"not null"`
}

func init() {
	if !db.SqlDB.HasTable("exhibition_collections") {
		db.SqlDB.CreateTable(&ExhibitionCollection{})
	} else {
		db.SqlDB.AutoMigrate(&ExhibitionCollection{})
	}
}

func GetExhibitionCollectionById(ctx context.Context, queryMap map[string][]string, unscoped bool) (*ExhibitionCollection, error) {
	var exhibitionCollection ExhibitionCollection
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql, num, err := plugin.ProcessQuery(sql.Table("exhibition_collections"), queryMap, "exhibition_collections")
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
			Information: "exhibition collections not found",
		}
	}
	if unscoped {
		err = sql.First(&exhibitionCollection).Error
	} else {
		err = sql.Unscoped().First(&exhibitionCollection).Error
	}
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return &exhibitionCollection, nil
}

func QueryExhibitionCollection(ctx context.Context, queryMap map[string][]string) ([]ExhibitionCollection, int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	exhibitionCollections := make([]ExhibitionCollection, 0)
	fmt.Println(queryMap)
	sql, num, err := plugin.ProcessQuery(sql.Table("exhibition_collections"), queryMap, "exhibition_collections")
	if err != nil {
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	sql = sql.Find(&exhibitionCollections)
	err = sql.Error
	if err != nil {
		println(err.Error())
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return exhibitionCollections, num, nil
}

func InsertExhibitionCollection(ctx context.Context, exhibitionCollection *ExhibitionCollection) (*ExhibitionCollection, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Create(exhibitionCollection)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return exhibitionCollection, nil
}

func DeleteExhibitionCollection(ctx context.Context, exhibitionCollection *ExhibitionCollection) (int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Delete(exhibitionCollection)
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

func UpdateExhibitionCollection(ctx context.Context, exhibitionCollection *ExhibitionCollection, updateMap map[string]interface{}) (*ExhibitionCollection, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Model(exhibitionCollection).Updates(updateMap)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return exhibitionCollection, nil
}
