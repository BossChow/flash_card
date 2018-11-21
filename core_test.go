package main

import (
	"life.kian/flash_card/model"
	"testing"
)

func TestSaveCard(t *testing.T) {
	card := &model.Card {
		Title: "Card Title",
		Content: "Card Content",
	}

	SaveCard(card)
}
