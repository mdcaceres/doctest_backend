package datasource

import (
	"github.com/mdcaceres/doctest/models"
	"github.com/mdcaceres/doctest/models/test"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("root:admin@/doctest_db?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("could not connect to doctest data base")
	}

	DB = conn

	conn.AutoMigrate(
		&models.User{},
		&models.Project{},
		&models.Role{},
		&models.Bug{},
		&models.Media{},
		&test.Case{},
		&test.Suit{},
		&test.Step{},
		&test.Result{})
}
