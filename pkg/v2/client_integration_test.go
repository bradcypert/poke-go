package v2

import (
	"context"
	"testing"
)

func TestClientIntegration_GetPokemon(t *testing.T) {
	client := NewClient()
	pokemonName := "pikachu"

	pokemon, err := client.GetPokemon(context.Background(), Name(pokemonName))
	if err != nil {
		t.Fatalf("failed to get pokemon: %v", err)
	}
	if pokemon.Name != pokemonName {
		t.Errorf("expected pokemon name to be %s, got %s", pokemonName, pokemon.Name)
	}
	if pokemon.ID == 0 {
		t.Errorf("expected pokemon ID to be greater than 0, got %d", pokemon.ID)
	}
	if len(pokemon.Types) == 0 {
		t.Errorf("expected pokemon to have types, got none")
	}
	if len(pokemon.Abilities) == 0 {
		t.Errorf("expected pokemon to have abilities, got none")
	}
	if len(pokemon.Moves) == 0 {
		t.Errorf("expected pokemon to have moves, got none")
	}
	if len(pokemon.Stats) == 0 {
		t.Errorf("expected pokemon to have stats, got none")
	}
}

func TestClientIntegration_GetAllPokemon(t *testing.T) {
	client := NewClient()
	pagination := PokeClientPagination{
		Limit:  1,
		Offset: 0,
	}

	pokemons, err := client.GetAllPokemon(context.Background(), pagination)
	if err != nil {
		t.Fatalf("failed to get all pokemons: %v", err)
	}
	if len(pokemons.Results) == 0 {
		t.Errorf("expected to get at least one pokemon, got none")
	}
	if pokemons.Count == 0 {
		t.Errorf("expected total count of pokemons to be greater than 0, got %d", pokemons.Count)
	}
}
func TestClientIntegration_GetGeneration(t *testing.T) {
	client := NewClient()
	generationID := 1

	generation, err := client.GetGeneration(context.Background(), ID(generationID))
	if err != nil {
		t.Fatalf("failed to get generation: %v", err)
	}
	if generation.ID != generationID {
		t.Errorf("expected generation ID to be %d, got %d", generationID, generation.ID)
	}
	if len(generation.PokemonSpecies) == 0 {
		t.Errorf("expected generation to have pokemon species, got none")
	}
}

func TestClientIntegration_GetGenerations(t *testing.T) {
	client := NewClient()
	pagination := PokeClientPagination{
		Limit:  1,
		Offset: 0,
	}

	generations, err := client.GetGenerations(context.Background(), pagination)
	if err != nil {
		t.Fatalf("failed to get all generations: %v", err)
	}
	if len(generations.Results) == 0 {
		t.Errorf("expected to get at least one generation, got none")
	}
	if generations.Count == 0 {
		t.Errorf("expected total count of generations to be greater than 0, got %d", generations.Count)
	}
}
