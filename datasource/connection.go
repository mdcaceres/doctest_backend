package datasource

import (
	"github.com/mdcaceres/doctest/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("root:rootroot@/doctest_db?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("could not connect to doctest data base")
	}

	DB = conn

	conn.AutoMigrate(
		&models.User{},
		&models.Project{},
		&models.Role{},
		&models.Bug{},
		&models.Suite{},
		&models.Invitation{},
		&models.Case{},
		&models.Step{},
		&models.Priority{})
}

func GetDB() *gorm.DB {
	return DB
}
