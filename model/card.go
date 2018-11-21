package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Card struct {
	gorm.Model
	Type    int
	Title   string
	Content string
}

func getDBConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	db.AutoMigrate(&Card{})

	return db
}

func CreateCard(card *Card) (uint, error) {
	db := getDBConnection()
	db.Create(card)

	return card.ID, nil
}

func QueryCardById(id int64) (*Card, error) {
	db := getDBConnection()

	var card Card
	db.First(&card, id)

	return &card, nil
}
