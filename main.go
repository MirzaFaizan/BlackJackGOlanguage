package main

import (
	"fmt"
)

func main(){
	//var card string = "Ace of 2"

	//USE colon equal only when declaring
	cards := [] string{
		newCard(),
		newCard(),
	}
	cards= append(cards,newCard())

	for i, card:= range cards {
		fmt.Println(i, card)
	}

}

func newCard () string{
	return "Five of diamonds"
}