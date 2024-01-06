package main

import "github.com/OmarEP/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client 
	nextLocationAreaURL *string 
	prevLocationAreaURL *string
}

func main() {

	cfg := config {
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}