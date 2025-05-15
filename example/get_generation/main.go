package main

import (
	"fmt"
	"log"

	v2 "github.com/bradcypert/poke-go/pkg/v2"
	// Replace with the actual import path of your SDK
)

func main() {
	client := v2.NewClient() // Initialize the SDK client

	generation, err := client.GetGeneration("1")
	if err != nil {
		log.Fatalf("Error fetching generation: %v", err)
	}

	// Print details about the generation
	fmt.Printf("Main Region: %s\n", generation.MainRegion.Name)
	fmt.Printf("Pokémon Species Count: %d\n", len(generation.PokemonSpecies))
}
