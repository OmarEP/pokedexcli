package main

import (
	"fmt"
	"errors"
)

func commandMap(cfg *config, name ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationURL)
	if err != nil {
		return err 
	}

	

	cfg.nextLocationURL = locationsResp.Next
	cfg.prevLocationURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, name ...string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}
	
	locationsResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationURL)
	if err != nil {
		return err 
	}

	cfg.nextLocationURL = locationsResp.Next
	cfg.prevLocationURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}