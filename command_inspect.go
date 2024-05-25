package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args[0]) == 0 {
		fmt.Println("Please provide a name to inspect.")
		return nil
	}
	pokemon, ok := cfg.pokedex[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", args[0])
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, ability := range pokemon.Abilities {
		if ability.IsHidden {
			continue
		}
		fmt.Printf(" - %s\n", ability.Ability.Name)
	}
	return nil
}
