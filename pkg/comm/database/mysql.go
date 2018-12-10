package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	_ "software_experiment/pkg/comm/utils"
)

var SqlDB *gorm.DB

func init() {
	var err error
	SqlDB, err = gorm.Open(viper.GetString("db.type"), viper.GetString("db.url"))
	SqlDB.LogMode(true)
	if err != nil {
		log.Fatal(err.Error())
	}
	SqlDB = SqlDB.Debug()
}
