package datasource

import (
	"github.com/mdcaceres/doctest/domains"
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
		&domains.User{},
		&domains.Project{},
		&domains.Role{},
		&domains.Bug{},
		&domains.Media{},
		&domains.TestCase{},
		&domains.TestSuit{},
		&domains.TestStep{},
		&domains.TestResult{})
}
