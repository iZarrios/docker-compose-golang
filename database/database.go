package database

import (
	"fmt"
	"os"

	"github.com/iZarrios/docker-compose-go-psql/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

// connect to database postgres
func Connect() (*gorm.DB, error) {

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	// migrate the fact model
	db.AutoMigrate(&models.Fact{})
	DB = DbInstance{Db: db}

	return db, nil

}
