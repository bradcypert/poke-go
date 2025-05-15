package main

import (
	"context"
	"encoding/json"
	"fmt"

	v2 "github.com/bradcypert/poke-go/pkg/v2"
)

func main() {
	client := v2.NewClient()
	pokemon, err := client.GetAllPokemon(context.Background(), v2.PokeClientPagination{Limit: 1, Offset: 1})
	if err != nil {
		panic(err)
	}

	pokemonJSON, err := json.MarshalIndent(pokemon, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(pokemonJSON))
}
