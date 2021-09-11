package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Email        string    `json:"email" gorm:"size:255;unique;not null"`
	PhoneNumber  string    `json:"phoneNumber" gorm:"size:255;unique;not null"`
	Country      string    `json:"country" gorm:"size:255;not null"`
	ParcelWeight float64   `json:"parcelWeight" gorm:"size:255"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	err := db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) GetUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}
