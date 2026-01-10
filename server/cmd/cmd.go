package cmd

import (
	"checkme/config"
	"checkme/internal/api/handler"
	"checkme/internal/api/router"
	"checkme/internal/repository"
	"checkme/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Start() {
	//读取配置文件
	var conf *config.Config
	var err error
	if conf, err = config.Load(); err != nil {
		panic(fmt.Sprintf("读取配置文件失败：%v", err))
	}

	//连接数据库
	var db *gorm.DB
	db, err = config.InitMysql(conf)

	if err != nil {
		panic(fmt.Sprintf("数据库连接失败：%v", err))
	}

	//创建仓库层
	recordRepo := repository.NewUserRepository(db)
	//创建服务层
	recordService := service.NewRecoderService(recordRepo, conf)
	//创建处理器
	h := handler.NewHandler(recordService)

	//设置Gin模式
	gin.SetMode(conf.Server.Mode)

	//创建路由
	r := gin.Default()

	router.Setup(r, h, conf)

	r.Run(fmt.Sprintf(":%d", conf.Server.Port))

}
