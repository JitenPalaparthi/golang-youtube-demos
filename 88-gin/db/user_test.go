package db

import (
	"demo/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserCreate(t *testing.T) {
	sqlDb, _ := gorm.Open(sqlite.Open(":memory:"))
	sqlDb.AutoMigrate(&models.User{}) // it creates a table
	//Init(sqlDb)
	userdb := NewUserDB(sqlDb)
	user := new(models.User)
	user.Name = "Jiten"
	user.Email = "JitenP@outlook.Com"
	user.Status = "active"
	user.LastModified = time.Now().Unix()
	// userExpt := user
	// userExpt.Id = 2
	_, err := userdb.Create(user)
	// fmt.Println(userExpt)
	// fmt.Println(userActual)

	assert.NoError(t, err)
	//assert.Equal(t, userExpt, userActual)
}
