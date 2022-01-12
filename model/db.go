package model

import (
	"GinProject/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

//全局变量
var db *gorm.DB
var err error

func InitDB() {
	//2022年1月6日版本
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DBUser, utils.DBPassWord, utils.DBHost, utils.DBPort, utils.DBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("连接数据库失败，失败原因为：", err)
		return
	}

	err = db.AutoMigrate(&User{}, &Category{}, &Article{})
	if err != nil {
		fmt.Println("无法迁移数据库，原因是：", err)
		return
	}

	//维护连接池
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("连接池设置错误，错误原因为：", err)
		return
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Minute)
}
