package database

import (
	"fmt"
	"github.com/cenkkoroglu/oz-fiber/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DBConn *gorm.DB
	err    error
)

func Init() error {
	config := config.GetConfig()
	driver := config.Database.Driver

	switch driver {
	case "postgres":
		DBConn, err = gorm.Open(postgres.New(postgres.Config{
			DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
				config.Database.Host,
				config.Database.Username,
				config.Database.Password,
				config.Database.DbName,
				config.Database.Port,
				config.Database.SslMode),
		}), &gorm.Config{
			Logger: LogMode(config.Database.LogLevel),
		})
		break
	case "mysql":
		DBConn, err = gorm.Open(mysql.New(mysql.Config{
			DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
				config.Database.Username,
				config.Database.Password,
				config.Database.Host,
				config.Database.Port,
				config.Database.DbName),
		}), &gorm.Config{
			Logger: LogMode(config.Database.LogLevel),
		})
		break
	}

	if err != nil {
		return err
	}

	return nil
}

func LogMode(logLevel int) logger.Interface {
	return logger.Default.LogMode(logger.LogLevel(logLevel))
}
