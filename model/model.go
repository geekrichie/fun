package model

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var v *viper.Viper

func GetDb() *gorm.DB {
	return db
}

func ReadConfig() {
	v = viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("config/")
	v.SetConfigType("ini")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatal("config file not found")
		} else {
			log.Fatal("config file error: ", err)
			// Config file was found but another error was produced
		}
	}
}

func init() {
	ReadConfig()
	var err error
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		v.GetString("mysql.user"),
		v.GetString("mysql.password"),
		v.GetString("mysql.ip"),
		v.GetInt("mysql.port"),
		v.GetString("mysql.dbname"))
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("database connect error: ", err)
	}

	sqlDB, _ := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 可以通过Set设置附加参数，下面设置表的存储引擎为InnoDB
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
}
