package v2

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	if client.BaseURL != PRODUCTION_POKEAPI_URL {
		t.Errorf("expected BaseURL to be %s, got %s", PRODUCTION_POKEAPI_URL, client.BaseURL)
	}
	if client.APIVersion != "v2" {
		t.Errorf("expected APIVersion to be v2, got %s", client.APIVersion)
	}
	if client.HTTPClient == nil {
		t.Error("expected HTTPClient to be initialized, got nil")
	}
}

func TestGetPokemon(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/pokemon/pikachu" {
			t.Errorf("unexpected URL path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name": "pikachu"}`))
	}))
	defer server.Close()

	client := NewClient()
	client.SetBaseURL(server.URL)
	pokemon, err := client.GetPokemon("pikachu")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pokemon.Name != "pikachu" {
		t.Errorf("expected pokemon name to be pikachu, got %s", pokemon.Name)
	}
}

func TestGetAllPokemon(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/pokemon" {
			t.Errorf("unexpected URL path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"results": [{"name": "pikachu"}, {"name": "bulbasaur"}]}`))
	}))
	defer server.Close()

	client := NewClient()
	client.SetBaseURL(server.URL)
	pokemons, err := client.GetAllPokemon(PokeClientPagination{
		Limit:  2,
		Offset: 0,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	results := pokemons.Results
	if len(results) != 2 {
		t.Errorf("expected 2 pokemons, got %d", len(results))
	}
	if results[0].Name != "pikachu" {
		t.Errorf("expected first pokemon to be pikachu, got %s", results[0].Name)
	}
}
func TestGetGeneration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/generation/generation-i" {
			t.Errorf("unexpected URL path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name": "generation-i"}`))
	}))
	defer server.Close()

	client := NewClient()
	client.SetBaseURL(server.URL)
	generation, err := client.GetGeneration("generation-i")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if generation.Name != "generation-i" {
		t.Errorf("expected generation name to be generation-i, got %s", generation.Name)
	}
}

func TestGetGenerations(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/generation" {
			t.Errorf("unexpected URL path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"results": [{"name": "generation-i"}, {"name": "generation-ii"}]}`))
	}))
	defer server.Close()

	client := NewClient()
	client.SetBaseURL(server.URL)
	generations, err := client.GetGenerations(PokeClientPagination{
		Limit:  2,
		Offset: 0,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	results := generations.Results
	if len(results) != 2 {
		t.Errorf("expected 2 generations, got %d", len(results))
	}
	if results[0].Name != "generation-i" {
		t.Errorf("expected first generation to be generation-i, got %s", results[0].Name)
	}
}
