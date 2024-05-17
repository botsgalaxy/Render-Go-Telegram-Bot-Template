package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	newLogger := logger.New(
		log.New(nil, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	var err error
	dsn := os.Getenv("DSN")
	if dsn == "" {
		panic("DSN Environment variable is empty")
	}
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connect with database")
	}
	sqlDB, _ := DB.DB()

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(50)

	sqlDB.SetConnMaxLifetime(time.Hour)
	DB.AutoMigrate(&TelegramUser{})

}
