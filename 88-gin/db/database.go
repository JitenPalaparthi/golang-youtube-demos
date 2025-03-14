package db

import (
	"demo/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	RETRY_COUNT    = 5
	RETRY_DURATION = time.Second * 5
)

func GetConnection(dsn string) (*gorm.DB, error) {
	count := 0
retry:
	count++
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("trying to connect to the database-->", count)
		time.Sleep(RETRY_DURATION)
		if count < RETRY_COUNT {
			goto retry
		} else {
			return nil, err
		}
	}
	Init(db)
	return db, nil
}

func Init(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}

//dsn := "host=localhost user=admin password=admin123 dbname=usersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
