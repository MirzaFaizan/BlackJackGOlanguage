package main
import (
	"os"
	"testing"
)
func TestNewDeck(t *testing.T){
	d := newDeck()
	if len(d)!= 16 {
		t.Errorf("Expecting deck legth of 16 but got %v",len(d))
	}
	if d[0] != "Ace of Spades"{
		t.Errorf("Expecting first Card to be Ace of Spades but instead got %v",d[0])
	}
	if d[len(d)-1]!="Four of Clubs"{
		t.Errorf("Expecting last Card to be Four of Clubs but instead got %v",d[len(d)-1])
	}
}

func TestSaveToDiskAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck:=newDeckFromFile("_decktesting")
	if len(loadedDeck)!=16{
		t.Errorf("Expecting first Card to be Ace of Spades but instead got %v",len(loadedDeck))
	}
	os.Remove("_decktesting")
}