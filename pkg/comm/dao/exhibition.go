package dao

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	db "software_experiment/pkg/comm/database"
	"software_experiment/pkg/web/plugin"
)

type Exhibition struct {
	gorm.Model
	Username string `gorm:"type:varchar(32);not null"`
	Content  string `gorm:"type:Text"`
	Theme    string `gorm:"type:Text"`
	Name     string `gorm:"type:varchar(32);not null;unique"`
	VisitNum int    `gorm:"default=0"`
}

func init() {
	if !db.SqlDB.HasTable("exhibitions") {
		db.SqlDB.CreateTable(&Exhibition{})
	} else {
		db.SqlDB.AutoMigrate(&Exhibition{})
	}
}

func GetExhibitionById(ctx context.Context, queryMap map[string][]string, unscoped bool) (*Exhibition, error) {
	var exhibition Exhibition
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql, num, err := plugin.ProcessQuery(sql.Table("exhibitions"), queryMap, "exhibitions")
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
			Information: "exhibitions not found",
		}
	}
	if unscoped {
		err = sql.First(&exhibition).Error
	} else {
		err = sql.Unscoped().First(&exhibition).Error
	}
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return &exhibition, nil
}

func QueryExhibition(ctx context.Context, queryMap map[string][]string) ([]Exhibition, int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	exhibitions := make([]Exhibition, 0)
	fmt.Println(queryMap)
	sql, num, err := plugin.ProcessQuery(sql.Table("exhibitions"), queryMap, "exhibitions")
	if err != nil {
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	sql = sql.Find(&exhibitions)
	err = sql.Error
	if err != nil {
		println(err.Error())
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return exhibitions, num, nil
}

func InsertExhibition(ctx context.Context, exhibition *Exhibition) (*Exhibition, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Create(exhibition)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return exhibition, nil
}

func DeleteExhibition(ctx context.Context, exhibition *Exhibition) (int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Delete(exhibition)
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

func UpdateExhibition(ctx context.Context, exhibition *Exhibition, updateMap map[string]interface{}) (*Exhibition, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Model(exhibition).Updates(updateMap)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return exhibition, nil
}
