package main
import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
	"math/rand"
)
//create a new type deck which is slice of string

type deck []string

//receiver fucntions binded with deck
func (d deck) print(){
	for i, card:= range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle(){
	for i:= range d{
		newPosition:=rand.Intn(len(d)-1)
		d[i],d[newPosition]=d[newPosition],d[i]
	}
}

func (d deck) toString() string{
	return strings.Join([]string(d),",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename,[]byte(d.toString()),0666)
}

//non reciever functions

func newDeckFromFile(filename string) deck{
	bs,err := ioutil.ReadFile(filename)
	if err!= nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs),",")) 
}


func newDeck() deck{
	cards:= deck{}
	cardSuits:= []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}


	for  _,suit :=range cardSuits{
		for  _,value :=range cardValues{
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func Deal(d deck, handSize int) (deck,deck){
	return d[:handSize],d[handSize:]
}