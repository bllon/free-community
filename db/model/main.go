package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
	"time"
)

var (
	once sync.Once

	DB *gorm.DB
)

func InitDatabase() {
	once.Do(func() {
		initDB()
	})
}

func initDB() {
	host := viper.GetString("Server.DB.Mysql8.Host")
	port := viper.GetString("Server.DB.Mysql8.Port")
	pw := viper.GetString("Server.DB.Mysql8.Passwd")
	dbname := viper.GetString("Server.DB.Mysql8.DBName")
	user := viper.GetString("Server.DB.Mysql8.User")
	minCon := viper.GetInt("Server.DB.Mysql8.MinConn")
	maxCon := viper.GetInt("Server.DB.Mysql8.MaxConn")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pw, host, port, dbname)
	DB = newDB(dsn, minCon, maxCon)
}

func newDB(dsn string, minCon, maxCon int) *gorm.DB {
	dbLogger := logger.Default.LogMode(logger.Silent)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   dbLogger,
	})
	if err != nil {
		log.Panicf("cannot connect to mysql %v", err)
	}
	dbConn, err := db.DB()
	if err != nil {
		log.Panicf("cannot get mysql %v", err)
	}
	// 设置连接池
	dbConn.SetMaxIdleConns(minCon)
	dbConn.SetMaxOpenConns(maxCon)
	dbConn.SetConnMaxLifetime(time.Second * 30)
	return db
}

func Migration() {
	fmt.Println("Begin migrate")

	if err := DB.AutoMigrate(
		&User{},
	); err != nil {
		log.Panicf("AutoMigrate DB %v", err)
	}
	fmt.Println("End migrate")
}
