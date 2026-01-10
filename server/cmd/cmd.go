package cmd

import (
	"checkme/config"
	"fmt"

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

}
