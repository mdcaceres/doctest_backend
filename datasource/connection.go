package datasource

import (
	"fmt"
	"github.com/mdcaceres/doctest/models"
	"github.com/mdcaceres/doctest/models/execution/SuiteExecution"
	"github.com/mdcaceres/doctest/models/execution/TestExecution"
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

	err = conn.AutoMigrate(
		&models.User{},
		&models.Project{},
		&models.Role{},
		&models.Bug{},
		&models.Suite{},
		&models.Invitation{},
		&models.Case{},
		&models.Step{},
		&models.Priority{},
		&models.ProjectClient{},
		&TestExecution.TestExecution{},
		&TestExecution.ExecutionStep{},
		&SuiteExecution.SuiteExecution{},
		&models.TestComment{},
		&models.BugComment{},
		&models.Post{})
	if err != nil {
		fmt.Println("could not migrate models")
		fmt.Println(err)
		return
	}
}

func GetDB() *gorm.DB {
	return DB

}
