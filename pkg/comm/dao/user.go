package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	db "software_experiment/pkg/comm/database"
	"software_experiment/pkg/web/plugin"
)

type User struct {
	gorm.Model
	Username     string `gorm:"type:varchar(32);not null"`
	Name         string `gorm:"type:varchar(32);not null"`
	Telephone    string `gorm:"type:varchar(32);not null"`
	Email        string `gorm:"type:varchar(32)"`
	PasswordHash string `gorm:"type:varchar(256)"`
	IconUrl      string `gorm:"type:varchar(256)"`
	Role         string `gorm:"type:varchar(32);not null"`
	Address      string `gorm:"type:varchar(256);not null"`
	QqNumber     string `gorm:"type:varchar(24)"`
	Information  string `gorm:"type:Text"`
}

func init() {
	if !db.SqlDB.HasTable("users") {
		db.SqlDB.CreateTable(&User{})
	} else {
		db.SqlDB.AutoMigrate(&User{})
	}
}

func GetUserById(ctx context.Context, queryMap map[string][]string, unscoped bool) (*User, error) {
	var user User
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql, num, err := plugin.ProcessQuery(sql.Table("users"), queryMap, "users")
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
			Information: "user not found",
		}
	}
	if unscoped {
		err = sql.First(&user).Error
	} else {
		err = sql.Unscoped().First(&user).Error
	}
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return &user, nil
}

func GetUserByUsername(ctx context.Context, queryMap map[string][]string, unscoped bool) (*User, error) {
	var user User
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql, num, err := plugin.ProcessQuery(sql.Table("users"), queryMap, "users")
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
			Information: "user not found",
		}
	}
	if unscoped {
		err = sql.First(&user).Error
	} else {
		err = sql.Unscoped().First(&user).Error
	}
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return &user, nil
}

func QueryUser(ctx context.Context, queryMap map[string][]string) ([]User, int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	users := make([]User, 0)
	sql = sql.Table("users")
	sql, num, err := plugin.ProcessQuery(sql, queryMap, "users")
	if err != nil {
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	sql = sql.Find(&users)
	err = sql.Error
	if err != nil {
		return nil, 0, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return users, num, nil
}

func InsertUser(ctx context.Context, user *User) (*User, error) {
	ctxValue := ctx.Value("tx")
	var sql *gorm.DB
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Create(user)
	err := sql.Error
	if err != nil {
		sql.Rollback()
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return user, nil
}

func DeleteUser(ctx context.Context, user *User) (int64, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Model(&user).Delete(user)
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

func UpdateUser(ctx context.Context, user *User, updateMap map[string]interface{}) (*User, error) {
	var sql *gorm.DB
	ctxValue := ctx.Value("tx")
	switch ctxValue.(type) {
	case *gorm.DB:
		sql = ctxValue.(*gorm.DB)
	default:
		sql = db.SqlDB
	}
	sql = sql.Model(&user).Updates(updateMap)
	err := sql.Error
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  500,
			Information: err.Error(),
		}
	}
	return user, nil
}
