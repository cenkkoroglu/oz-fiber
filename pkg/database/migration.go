package database

import "github.com/cenkkoroglu/oz-fiber/app/models/entities"

func Migrate() error {
	return DBConn.AutoMigrate(
		&entities.User{},
	)
}
