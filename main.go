package main

import (
	commandline "github.com/suchaos/learn-go-with-tdd/command-line"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := commandline.NewPlayerServer(&InMemoryPlayerStore{})
	log.Fatal(http.ListenAndServe(":5000", server))
}
