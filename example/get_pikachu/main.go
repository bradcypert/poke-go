package main

import (
	"context"
	"encoding/json"
	"fmt"

	v2 "github.com/bradcypert/poke-go/pkg/v2"
)

func main() {
	client := v2.NewClient()
	pokemon, err := client.GetPokemon(context.Background(), v2.Name("pikachu"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	pokemonJSON, err := json.MarshalIndent(pokemon, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println(string(pokemonJSON))
}
