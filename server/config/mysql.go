package config

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(conf *Config) (*gorm.DB, error) {
	cfg := conf.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, errors.New(fmt.Sprintf("MySql连接失败:%v", err))
	}

	return db, nil
}
