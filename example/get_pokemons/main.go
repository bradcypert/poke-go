package main

import (
	"encoding/json"
	"fmt"

	v2 "github.com/bradcypert/poke-go/pkg/v2"
)

func main() {
	client := v2.NewClient()
	pokemon, err := client.GetAllPokemon(5, 46)
	if err != nil {
		panic(err)
	}

	pokemonJSON, err := json.MarshalIndent(pokemon, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(pokemonJSON))
}
