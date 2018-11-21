package main

import (
	"fmt"
	"life.kian/flash_card/model"
)

func SaveCard(card *model.Card) (int, error) {
	fmt.Printf("Save Card: %+v", card)
	return 1, nil
}

func main() {
	fmt.Println("vim-go")
}
