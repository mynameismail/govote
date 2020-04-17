package models

import (
	"log"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Vote struct {
	gorm.Model
	Name  string `gorm:"column:name;not null;unique" json:"name"`
	Votes int    `gorm:"column:votes;not null;default:0" json:"votes"`
}

type Voter struct {
	gorm.Model
	Code   string `gorm:"column:code;not null;unique" json:"code"`
	Status string `gorm:"column:status;not null;default:'init'" json:"status"` // ["init", "logged", "voted"]
}

var Conn *gorm.DB

func Init() {
	db, err := gorm.Open("sqlite3", "govote.db")
	if err != nil {
		log.Println("Failed to connect database")
	}

	db.AutoMigrate(&Vote{}, &Voter{})
	// defer db.Close()

	Conn = db

	rand.Seed(time.Now().UnixNano())
}

const bytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = bytes[rand.Intn(len(bytes))]
	}

	return string(b)
}
