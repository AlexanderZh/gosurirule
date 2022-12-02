package database

import (
    "os"
	"github.com/AlexanderZh/gosurirule/go/model"	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB (){
	db_uri, ok := os.LookupEnv("DB_URI")
	if !ok {
		db_uri = "host=localhost port=5432 user=postgres dbname=stage sslmode=disable password=secret"
	}
	print(db_uri)
	var err error
	DB, err = gorm.Open(postgres.Open(db_uri), &gorm.Config{})
	if err != nil {
		panic(err.Error() + " - failed to connect database")
	}
	//automigrate tables
	err = DB.AutoMigrate(model.Ruleset{})
	if err != nil {
		panic(err.Error() + " - Ruleset automigration failed")
	}
	err = DB.AutoMigrate(model.Rule{})
	if err != nil {
		panic(err.Error() + " - Rules automigration failed")
	}
}
