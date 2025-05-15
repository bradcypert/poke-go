package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	internal "github.com/bradcypert/poke-go/internal/v2"
)

const PRODUCTION_POKEAPI_URL = "https://pokeapi.co/api"

type pokeClient struct {
	BaseURL    string
	APIVersion string
	HTTPClient *http.Client
}

type PokeClientPagination = internal.PokeClientPagination

// NewClient creates a new PokeClient with default settings.
// The default base URL is the production PokeAPI URL, and the default API version is "v2".
func NewClient() *pokeClient {
	return &pokeClient{
		BaseURL:    PRODUCTION_POKEAPI_URL,
		APIVersion: "v2",
		HTTPClient: &http.Client{},
	}
}

// SetBaseURL sets the base URL for the PokeClient.
// This can be used to set a custom base URL for testing or other if hosting your own POKEAPI.
func (c *pokeClient) SetBaseURL(baseURL string) {
	c.BaseURL = baseURL
}

// SetAPIVersion sets the API version for the PokeClient.
// This can be used to set a custom API version if needed.
// The default API version is "v2".
func (c *pokeClient) SetAPIVersion(apiVersion string) {
	c.APIVersion = apiVersion
}

// Name returns a Key for a Pokemon by its name.
// This is used to retrieve a specific Pokemon from the PokeAPI.
func Name(name string) Key {
	return NameKey{
		Name: name,
	}
}

// ID returns a Key for a Pokemon by its ID.
// This is used to retrieve a specific Pokemon from the PokeAPI.
// The ID is an integer that represents the Pokemon's unique identifier.
func ID(id int) Key {
	return IDKey{
		ID: id,
	}
}

// getPokeAPIUrl constructs the full URL for a given resource.
// It combines the base URL, API version, and resource path to create the full URL.
// The resource path is provided as a variadic argument, allowing for multiple segments.
// It returns a parsed URL or an error if the URL could not be constructed.
func (c *pokeClient) getPokeAPIUrl(resource ...string) (*url.URL, error) {
	fullPath := fmt.Sprintf("%s/%s/%s", c.BaseURL, c.APIVersion, path.Join(resource...))
	return url.Parse(fullPath)
}

// GetPokemon retrieves a Pokemon by its key (name or ID).
func (c *pokeClient) GetPokemon(context context.Context, key Key) (*internal.Pokemon, error) {
	u, err := c.getPokeAPIUrl("pokemon", key.getResourceKey())
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}

	respData, err := c.makeRequest(context, u)
	if err != nil {
		return nil, fmt.Errorf("failed to make request[%s]: %v", u.String(), err)
	}

	pokemon := &internal.Pokemon{}
	err = json.Unmarshal(respData, pokemon)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

// GetAllPokemon retrieves all Pokemon with optional pagination.
// It returns a ResultSet containing the list of Pokemon.
// Pagination can be used to limit the number of results and set an offset.
// The pagination parameters are passed as a PokeClientPagination struct.
func (c *pokeClient) GetAllPokemon(context context.Context, pagination PokeClientPagination) (*internal.ResultSet, error) {
	u, err := c.getPokeAPIUrl("pokemon")
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}

	if pagination.Limit > 0 || pagination.Offset > 0 {
		internal.AddPaginationToURL(u, pagination)
	}

	respData, err := c.makeRequest(context, u)
	if err != nil {
		return nil, fmt.Errorf("failed to make request[%s]: %v", u.String(), err)
	}

	pokemon := internal.ResultSet{}
	err = json.Unmarshal(respData, &pokemon)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response data: %v", err)
	}
	return &pokemon, nil
}

// GetGeneration retrieves a Generation by its key (name or ID).
func (c *pokeClient) GetGeneration(context context.Context, key Key) (*internal.Generation, error) {
	u, err := c.getPokeAPIUrl("generation", key.getResourceKey())
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}

	respData, err := c.makeRequest(context, u)
	if err != nil {
		return nil, fmt.Errorf("failed to make request[%s]: %v", u.String(), err)
	}

	generation := &internal.Generation{}
	err = json.Unmarshal(respData, generation)
	if err != nil {
		return nil, err
	}
	return generation, nil
}

// GetGenerations retrieves all Generations with optional pagination.
// It returns a ResultSet containing the list of Generations.
// Pagination can be used to limit the number of results and set an offset.
// The pagination parameters are passed as a PokeClientPagination struct.
func (c *pokeClient) GetGenerations(context context.Context, pagination PokeClientPagination) (*internal.ResultSet, error) {
	u, err := c.getPokeAPIUrl("generation")
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}

	if pagination.Limit > 0 || pagination.Offset > 0 {
		internal.AddPaginationToURL(u, pagination)
	}

	respData, err := c.makeRequest(context, u)

	if err != nil {
		return nil, fmt.Errorf("failed to make request[%s]: %v", u.String(), err)
	}

	generations := internal.ResultSet{}
	err = json.Unmarshal(respData, &generations)
	if err != nil {
		return nil, err
	}
	return &generations, nil
}

// makeRequest makes an HTTP GET request to the given URL and returns the response body as a byte slice.
// Uses the provided context to allow for cancellation and timeout.
func (c *pokeClient) makeRequest(context context.Context, u *url.URL) ([]byte, error) {
	req, err := http.NewRequestWithContext(context, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %v", err)
	}
	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
