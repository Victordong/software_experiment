package dao

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	db "software_experiment/pkg/comm/database"
	"software_experiment/pkg/web/plugin"
)

type ExhibitionComment struct {
	gorm.Model
	Username      string `gorm:"type:varchar(32);not null"`
	Content       string `gorm:"type:Text"`
	CommentedId   uint   `gorm:"not null"`
	CommentedName string `gorm:"type:varchar(16);not null"`
}

func init() {
	if !db.SqlDB.HasTable("exhibition_comments") {
		db.SqlDB.CreateTable(&ExhibitionComment{})
	} else {
		db.SqlDB.AutoMigrate(&ExhibitionComment{})
	}
}

func GetExhibitionCommentById(ctx context.Context, queryMap map[string][]string, unscoped bool) (*ExhibitionComment, error) {
	var exhibitionComment ExhibitionComment
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql, num, err := plugin.ProcessQuery(sql.Table("exhibition_collections"), queryMap, "exhibition_comments")
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
		err = sql.First(&exhibitionComment).Error
	} else {
		err = sql.Unscoped().First(&exhibitionComment).Error
	}
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return &exhibitionComment, nil
}

func QueryExhibitionComment(ctx context.Context, queryMap map[string][]string) ([]ExhibitionComment, int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	exhibitionComments := make([]ExhibitionComment, 0)
	fmt.Println(queryMap)
	sql, num, err := plugin.ProcessQuery(sql.Table("exhibition_collections"), queryMap, "exhibition_collections")
	if err != nil {
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	sql = sql.Find(&exhibitionComments)
	err = sql.Error
	if err != nil {
		println(err.Error())
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return exhibitionComments, num, nil
}

func InsertExhibitionComment(ctx context.Context, exhibitionComment *ExhibitionComment) (*ExhibitionComment, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Create(exhibitionComment)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return exhibitionComment, nil
}

func DeleteExhibitionComment(ctx context.Context, exhibitionComment *ExhibitionComment) (int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Delete(exhibitionComment)
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

func UpdateExhibitionComment(ctx context.Context, exhibitionComment *ExhibitionComment, updateMap map[string]interface{}) (*ExhibitionComment, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Model(exhibitionComment).Updates(updateMap)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return exhibitionComment, nil
}
