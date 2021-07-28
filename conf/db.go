package conf

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func init()  {
	var err error
	DB, err = gorm.Open(mysql.Open(viper.GetString("db")), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Fatal error open db: %s \n", err))
	}
}

func InitDB()*sql.DB {
	sqlDB, err := DB.DB()
	if err != nil {
		panic(fmt.Errorf("Fatal error set db: %s \n", err))
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return sqlDB
}