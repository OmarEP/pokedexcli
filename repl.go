package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/OmarEP/pokedexcli/internal/pokeapi"
	"github.com/OmarEP/pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient	pokeapi.Client
	pokecache		pokecache.Cache
	nextLocationURL *string 
	prevLocationURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type clicCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]clicCommand {
	return map[string]clicCommand {
		"help": {
			name:			"help",
			description: 	"Displays a help message",
			callback:       commandHelp,	
		},
		"exit": {
			name:			"exit",
			description:   	"Exit the Pokedex",
			callback:    	commandExit,
		},
		"map": {
			name: 			"map",
			description: 	"Displays the names of 20 location areas in the Pokemon world",
			callback: 		commandMap,
		},
		"mapb": {
			name:			"mapb",
			description: 	"Displays the previous 20 locations",
			callback:   	commandMapb,	
		},
		"explore": {
			name:			"explore",
			description:  	"explore <location_name>",
			callback: 		commandExplore,	
		},
		"catch": {
			name: 			"catch",
			description:  	"catch <pokemon_name>",
			callback:  		commandCatch,	
		},
		"inspect": {
			name:			"inspect",
			description: 	"inspect <pokemon_name>",
			callback: 		commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See all the pokemon you've caught",
			callback:    commandPokedex,
		},
	}
}

