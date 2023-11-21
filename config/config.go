package config

import (
	"fmt"
	"go-web-echo-app/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	ServerPORT int
}

var DB *gorm.DB

func LoadEnvVariables() {

}

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("loading env vars %s", err.Error()))
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSL_MODE")
	sslRootCert := os.Getenv("SSL_ROOT_CERT")

	fmt.Println("DB IS CONNECTED YEAH???? 0.1")

	// DSN for PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s sslrootcert=%s", dbHost, dbUser, dbPass, dbName, dbPort, sslMode, sslRootCert)
	fmt.Println("dsn: ", dsn)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Connecting with dsn string failed: %s", err.Error()))
	}

	fmt.Println("DB IS CONNECTED YEAH? 1")

	db.AutoMigrate(&models.User{})
	if err != nil {
		panic(fmt.Sprintf("Automigrate failed: %s", err.Error()))
	}

	fmt.Println("DB IS CONNECTED YEAH!2")
	DB = db

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("Closing db failed: %s", err.Error()))
	}
	dbSQL.Close()
}
