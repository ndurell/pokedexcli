package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Must provide an area to explore.")
	}

	areaResponse, err := cfg.pokeapiClient.ExploreArea(args[0])
	if err != nil {
		return err
	}

	for _, encounter := range areaResponse.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
