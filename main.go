package main

import (
	"fmt"
	"free-community/db/model"
	"free-community/router"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	engine.Use(gin.Recovery())

	//读取配置
	viper.SetConfigName("local_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("fatal error config file: %w", err)
	}

	//初始化数据库
	model.InitDatabase()
	model.Migration()

	//初始化路由
	router.InitRouter(engine)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", viper.GetString("Server.Service.Port")),
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("http server start error: %v", err)
	}
	return
}
