package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"wisdom/config"
	"wisdom/data/model"
	_ "wisdom/docs"
	"wisdom/global"
)

func init() {
	var err error
	if err = SetupSetting(); err != nil {
		log.Fatalf("文件初始化失败：%s", err)
	}
	if err = SetupDatabase(); err != nil {
		log.Fatalf("数据库初始化失败：%s", err)
	}
}
func main() {
	r := gin.New()

	// 创建路由组

	// 文档界面访问URL
	// http://127.0.0.1:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

func SetupSetting() error {
	var (
		err     error
		setting *config.Setting
	)

	setting, err = config.NewSetting()
	if err != nil {
		return err
	}
	if err = setting.ReadSection("server", &global.ServerSetting); err != nil {
		return err
	}
	if err = setting.ReadSection("mysql", &global.DatabaseSetting); err != nil {
		return err
	}
	return nil
}
func SetupDatabase() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
