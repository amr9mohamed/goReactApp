package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Email        string    `json:"email" gorm:"size:255;unique;not null"`
	PhoneNumber  string    `json:"phoneNumber" gorm:"size:255;unique;not null"`
	Country      string    `json:"country" gorm:"index:,option:CONCURRENTLY;size:255;not null"`
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
	users := []User{}
	err := db.Debug().Model(&User{}).Limit(250).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) GetDistinctCountries(db *gorm.DB) (*[]string, error) {
	var countries []string
	err := db.Debug().Order("country").Model(&User{}).Distinct().Pluck("country", &countries).Error
	if err != nil {
		return &[]string{}, err
	}
	return &countries, err
}

type countryFrequency struct {
	Country   string
	Frequency int64
}

func (u *User) GetCountyFrequency(db *gorm.DB) (*[]countryFrequency, error) {
	result := []countryFrequency{}
	db.Table("users").Select("country, count(country) as frequency").Group("country").Order("country").Scan(&result)
	return &result, nil
}
