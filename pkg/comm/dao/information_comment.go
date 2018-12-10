package dao

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	db "software_experiment/pkg/comm/database"
	"software_experiment/pkg/web/plugin"
)

type InformationComment struct {
	gorm.Model
	Username      string `gorm:"type:varchar(32);not null"`
	Content       string `gorm:"type:Text"`
	CommentedId   uint   `gorm:"not null"`
	CommentedName string `gorm:"type:varchar(16);not null"`
}

func init() {
	if !db.SqlDB.HasTable("information_comments") {
		db.SqlDB.CreateTable(&InformationComment{})
	} else {
		db.SqlDB.AutoMigrate(&InformationComment{})
	}
}

func GetInformationCommentById(ctx context.Context, queryMap map[string][]string, unscoped bool) (*InformationComment, error) {
	var informationComment InformationComment
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql, num, err := plugin.ProcessQuery(sql.Table("information_comments"), queryMap, "information_comments")
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
			Information: "comment not found",
		}
	}
	if unscoped {
		err = sql.First(&informationComment).Error
	} else {
		err = sql.Unscoped().First(&informationComment).Error
	}
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return &informationComment, nil
}

func QueryInformationComment(ctx context.Context, queryMap map[string][]string) ([]InformationComment, int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	informationComments := make([]InformationComment, 0)
	fmt.Println(queryMap)
	sql, num, err := plugin.ProcessQuery(sql.Table("information_comments"), queryMap, "information_comments")
	if err != nil {
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	sql = sql.Find(&informationComments)
	err = sql.Error
	if err != nil {
		println(err.Error())
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return informationComments, num, nil
}

func InsertInformationComment(ctx context.Context, informationComment *InformationComment) (*InformationComment, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Create(informationComment)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return informationComment, nil
}

func DeleteInformationComment(ctx context.Context, informationComment *InformationComment) (int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Delete(informationComment)
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

func UpdateInformationComment(ctx context.Context, informationComment *InformationComment, updateMap map[string]interface{}) (*InformationComment, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Model(informationComment).Updates(updateMap)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return informationComment, nil
}
