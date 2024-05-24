package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) == 0 {
		return errors.New("Must provide the pokemon name.")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}
	winChance := determineWinChance(pokemonResp.BaseExperience)
	roll := rand.Intn(100)
	if roll > winChance {
		fmt.Printf("%s escaped!\n", args[0])
	} else {
		fmt.Printf("%s was caught!\n", args[0])
		cfg.pokedex[args[0]] = pokemonResp
	}

	return nil
}

func determineWinChance(baseXP int) int {
	fmt.Printf("Base Xp: %d\n", baseXP)
	if baseXP > 100 {
		return 10
	} else if baseXP > 75 {
		return 20
	}

	return 50
}
