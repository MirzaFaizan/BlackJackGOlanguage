package main

func main(){
	//var card string = "Ace of 2"

	//USE colon equal only when declaring
	cards := newDeck()

	hand,remainingDeck := Deal(cards,5)
	hand.print()
	remainingDeck.print()
}
