package dao

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	db "software_experiment/pkg/comm/database"
	"software_experiment/pkg/web/plugin"
)

type SupplyComment struct {
	gorm.Model
	Username      string `gorm:"type:varchar(32);not null"`
	Content       string `gorm:"type:Text"`
	CommentedId   uint   `gorm:"not null"`
	CommentedName string `gorm:"type:varchar(16);not null"`
}

func init() {
	if !db.SqlDB.HasTable("supply_comments") {
		db.SqlDB.CreateTable(&SupplyComment{})
	} else {
		db.SqlDB.AutoMigrate(&SupplyComment{})
	}
}

func GetSupplyCommentById(ctx context.Context, queryMap map[string][]string, unscoped bool) (*SupplyComment, error) {
	var supplyComment SupplyComment
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql, num, err := plugin.ProcessQuery(sql.Table("supply_comments"), queryMap, "supply_comments")
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
			Information: "supply comments not found",
		}
	}
	if unscoped {
		err = sql.First(&supplyComment).Error
	} else {
		err = sql.Unscoped().First(&supplyComment).Error
	}
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return &supplyComment, nil
}

func QuerySupplyComment(ctx context.Context, queryMap map[string][]string) ([]SupplyComment, int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	supplyComments := make([]SupplyComment, 0)
	fmt.Println(queryMap)
	sql, num, err := plugin.ProcessQuery(sql.Table("supply_comments"), queryMap, "supply_comments")
	if err != nil {
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	sql = sql.Find(&supplyComments)
	err = sql.Error
	if err != nil {
		println(err.Error())
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return supplyComments, num, nil
}

func InsertSupplyComment(ctx context.Context, supplyComment *SupplyComment) (*SupplyComment, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Create(supplyComment)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return supplyComment, nil
}

func DeleteSupplyComment(ctx context.Context, supplyComment *SupplyComment) (int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Delete(supplyComment)
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

func UpdateSupplyComment(ctx context.Context, supplyComment *SupplyComment, updateMap map[string]interface{}) (*SupplyComment, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Model(supplyComment).Updates(updateMap)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return supplyComment, nil
}
