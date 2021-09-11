package seed

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/amr9mohamed/mainApp/api/models"
	"gorm.io/gorm"
)

var (
	users        []models.User
	code2country = map[string]string{
		"237": "Cameroon",
		"251": "Ethiopia",
		"212": "Morocoo",
		"258": "Mozambique",
		"256": "Uganda ",
	}
	code2regex = map[string]string{
		"237": "237 ?[2368]\\d{7,8}$",
		"251": "251 ?[1-59]\\d{8}$",
		"212": "212 ?[5-9]\\d{8}$",
		"258": "258 ?[28]\\d{7,8}$",
		"256": "256 ?\\d{9}$",
	}
)

func getCode(num string) string {
	if len(num) > 3 {
		return num[:3]
	}
	return ""
}

func seedDB() {
	fp, err := os.Open("/root/test_file.csv")
	if err != nil {
		log.Fatal("Error while opening file", err)
	}
	defer fp.Close()

	lines, err := csv.NewReader(fp).ReadAll()
	if err != nil {
		log.Fatal("Error in parsing csv file", err)
	}

	for _, line := range lines[1:] {
		id, _ := strconv.ParseUint(strings.Trim(line[0], " "), 10, 64)
		email := strings.Trim(line[1], " ")
		pn := strings.Trim(line[2], " ")
		w, _ := strconv.ParseFloat(strings.Trim(line[3], " "), 64)
		code := getCode(pn)
		if code == "" {
			fmt.Println("Error, phone format not supported", pn)
			continue
		}
		if country, ok := code2country[code]; ok {
			if match, _ := regexp.MatchString(code2regex[code], pn); match {
				u := models.User{
					ID:           id,
					Email:        email,
					Country:      country,
					PhoneNumber:  pn,
					ParcelWeight: w,
				}
				users = append(users, u)
			}
		}
	}
}

func Load(db *gorm.DB) {
	err := db.Debug().Migrator().DropTable(&models.User{})
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	seedDB()
	users = users[:250]
	err = db.Debug().Model(&models.User{}).Create(&users).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
}
