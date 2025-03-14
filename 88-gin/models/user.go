package models

import "errors"

type User struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	Contact      string `json:"contact" gorm:"unique"`
	LinkedIn     string `json:"linkedin" gorm:"column:linkedin"`
	Status       string `json:"status" gorm:"default:active"`
	LastModified int64  `json:"last_modified" gorm:"autoCreateTime"` // unixtime
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("invalid name")
	}
	if u.Email == "" {
		return errors.New("invalid email")
	}
	return nil
}
