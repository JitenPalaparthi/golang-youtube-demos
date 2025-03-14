package db

import (
	"demo/models"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type userDb struct {
	DB *gorm.DB
}

type IUserDB interface {
	Create(*models.User) (*models.User, error)
	Update(*models.User) (*models.User, error)
	GetByID(id string) (*models.User, error)
	DeleteByID(id string) error
	GetAll() ([]models.User, error)
	GetAllBy(limit, offset int) ([]models.User, error)
}

func NewUserDB(db *gorm.DB) IUserDB {
	return &userDb{DB: db}
}

func (u *userDb) Create(user *models.User) (*models.User, error) {
	tx := u.DB.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *userDb) Update(user *models.User) (*models.User, error) { // it updates all the records.
	user.LastModified = time.Now().Unix()
	tx := u.DB.Save(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *userDb) GetByID(id string) (*models.User, error) {
	user := new(models.User)
	tx := u.DB.First(user, id)
	log.Println(tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return nil, errors.New("no data")
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *userDb) GetAll() ([]models.User, error) {
	//users := make([]models.User, 10)
	var users []models.User
	tx := u.DB.Find(&users)
	//tx := u.DB.First(user, id)
	log.Println(tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return nil, errors.New("no data")
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (u *userDb) GetAllBy(limit, offset int) ([]models.User, error) {
	//users := make([]models.User, 10)
	var users []models.User
	tx := u.DB.Limit(limit).Offset(offset).Find(&users)
	//tx := u.DB.First(user, id)
	log.Println(tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return nil, errors.New("no data")
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (u *userDb) DeleteByID(id string) error {
	tx := u.DB.Delete(&models.User{}, id)
	log.Println(tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return errors.New("no data to delete")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
