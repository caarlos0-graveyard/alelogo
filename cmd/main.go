package main

import (
	"log"
	"os"

	"github.com/caarlos0/alelogo"
)

func main() {
	cpf := os.Args[1]
	pwd := os.Args[2]
	client, err := alelogo.New(cpf, pwd)
	if err != nil {
		log.Fatalln(err.Error())
	}
	cards, err := client.Cards()
	if err != nil {
		log.Fatalln(err.Error())
	}
	for _, card := range cards {
		result, err := client.Details(card)
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(result)
	}
}
