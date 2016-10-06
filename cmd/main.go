package main

import (
	"fmt"
	"os"

	"github.com/caarlos0/alelogo"
)

func main() {
	if len(os.Args) != 3 {
		panic("You must pass the CPF and password as args 1 and 2")
	}
	client, err := alelogo.New(os.Args[1], os.Args[2])
	if err != nil {
		panic(err)
	}
	cards, err := client.Balance()
	if err != nil {
		panic(err)
	}
	for _, card := range cards {
		fmt.Println("Card " + card.Title + " balance is " + card.Balance)
	}
}
