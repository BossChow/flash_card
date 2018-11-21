package model

import (
	"encoding/json"
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCreateCard(t *testing.T) {
	card := &Card{
		Type:    0,
		Title:   "Card Title",
		Content: "This is card",
	}

	id, err := CreateCard(card)
	assert.Equal(t, err, nil)
	assert.Equal(t, true, id > 0)
}

func TestQueryCardById(t *testing.T) {
	card, err := QueryCardById(4)
	assert.Equal(t, err, nil)
	jsonStr, err := json.Marshal(card)
	fmt.Printf("Card: %+v", string(jsonStr))
}
