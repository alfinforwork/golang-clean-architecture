package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/teten-nugraha/golang-crud/domain"
	"os"
	"path/filepath"
)

func InitDB() *gorm.DB  {
	processENV()

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")


	db, err := gorm.Open("mysql", dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	if err != nil {
		logrus.Error("Cannot Connect to MySQL DB")
		panic(err)
	}

	migrateDDL(db)

	return db
}

func migrateDDL(db *gorm.DB) {
	// migrasi domain ke tabel di mysql

	db.AutoMigrate(&domain.Mahasiswa{})
}

func processENV() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(".env"); err != nil {
			logrus.Warn("Error loading env file: " + err.Error())
		}
	}
}
