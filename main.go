package main

import (
	"fmt"
	"go_web/goweb-gin-blog/dao"
	"go_web/goweb-gin-blog/loggers"
	"go_web/goweb-gin-blog/models"
	"go_web/goweb-gin-blog/routers"
	"go_web/goweb-gin-blog/settings"
)

func main() {

	// 加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Printf("init config failed, err: %v\n", err)
		return
	}

	// 日志初始化
	if err := loggers.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	// 连接数据库
	if err := dao.InitMySQL(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close()

	// 模型绑定
	dao.DB.AutoMigrate(new(models.Category), new(models.Comment), new(models.Config), new(models.Post), new(models.User))

	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
