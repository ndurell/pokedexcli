package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for name, _ := range cfg.pokedex {
		fmt.Printf("  - %s\n", name)
	}
	return nil
}
