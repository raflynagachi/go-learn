package app

import (
	"fmt"
	"log"

	"github.com/raflynagachi/go-store/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	DB     string
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBname string
}

func (dbConfig *DBConfig) setupEnv() {
	dbConfig.DB = helpers.GetEnv("DATABASE", "mysql")
	dbConfig.DBHost = helpers.GetEnv("DATABASE_HOST", "localhost")
	dbConfig.DBPort = helpers.GetEnv("DATABASE_PORT", "3306")
	dbConfig.DBUser = helpers.GetEnv("DATABASE_USERNAME", "admin")
	dbConfig.DBPass = helpers.GetEnv("DATABASE_PASSWORD", "password")
	dbConfig.DBname = helpers.GetEnv("DATABASE_NAME", "go-store")

}

func (s *Server) InitializeDB(dbConfig DBConfig) {
	var dsn string
	var err error

	if dbConfig.DB == "mysql" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.DBUser,
			dbConfig.DBPass,
			dbConfig.DBHost,
			dbConfig.DBPort,
			dbConfig.DBname,
		)
		s.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	} else if dbConfig.DB == "postgres" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			dbConfig.DBHost,
			dbConfig.DBUser,
			dbConfig.DBPass,
			dbConfig.DBname,
			dbConfig.DBPort,
		)
		s.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic("Failed on connecting to the database server")
	}
	fmt.Printf("Connected to %s at %s:%s\n", dbConfig.DB, dbConfig.DBHost, dbConfig.DBPort)
}

func (s *Server) MigrateDB() {
	for _, model := range RegisterModels() {
		err := s.DB.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal("err")
		}
	}

	fmt.Println("Database migrated successfully")
}
