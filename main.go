package main

import (
	"fmt"
)

var log = fmt.Println

func main() {

	cards := newDeckFromFile("my_ards")

	log(cards)

}
