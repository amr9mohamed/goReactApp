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

func Paginate(pageNumber int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := pageNumber
		if page == 0 {
			page = 1
		}

		pageSize := pageSize
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
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

func (u *User) GetUsersByCountry(db *gorm.DB, country string, pageNumber int, pageSize int) (*[]User, error) {
	users := []User{}
	err := db.Debug().Scopes(Paginate(pageNumber, pageSize)).Where("country = ?", country).Find(&users).Error
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

type CountryFrequency struct {
	Country   string `json:"country"`
	Frequency int64  `json:"frequency"`
}

func (u *User) GetCountyFrequency(db *gorm.DB) (*[]CountryFrequency, error) {
	result := []CountryFrequency{}
	db.Table("users").Select("country, count(country) as frequency").Group("country").Order("country").Scan(&result)
	return &result, nil
}
