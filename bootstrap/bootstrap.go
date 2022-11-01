package bootstrap

import (
	"blog/base"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() {
	initViper()
	initGorm()
}

func initViper() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("application")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initGorm() {
	mysqlMap := viper.GetStringMapString("mysql")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", mysqlMap["username"], mysqlMap["password"], mysqlMap["path"], mysqlMap["db-name"])

	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector)
	if err != nil {
		panic(err)
	}
	base.Db = db
}
