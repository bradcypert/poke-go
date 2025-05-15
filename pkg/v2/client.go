package v2

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	internal "github.com/bradcypert/poke-go/internal/v2"
)

const PRODUCTION_POKEAPI_URL = "https://pokeapi.co/api"

type pokeClient struct {
	BaseURL    string
	APIVersion string
	HTTPClient *http.Client
}

type PokeClientPagination struct {
	Limit  int
	Offset int
}

func NewClient() *pokeClient {
	return &pokeClient{
		BaseURL:    PRODUCTION_POKEAPI_URL,
		APIVersion: "v2",
		HTTPClient: &http.Client{},
	}
}

func (c *pokeClient) SetBaseURL(baseURL string) {
	c.BaseURL = baseURL
}

func (c *pokeClient) SetAPIVersion(apiVersion string) {
	c.APIVersion = apiVersion
}

func (c *pokeClient) GetPokemon(idOrName string) (*internal.Pokemon, error) {
	u, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.BaseURL, c.APIVersion, "pokemon", idOrName))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}
	resp, err := c.HTTPClient.Get(u.String())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	pokemon := &internal.Pokemon{}
	err = json.Unmarshal(respData, pokemon)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

func (c *pokeClient) GetAllPokemon(pagination PokeClientPagination) (*internal.ResultSet, error) {
	u, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.BaseURL, c.APIVersion, "pokemon"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}

	if pagination.Limit > 0 || pagination.Offset > 0 {
		query := u.Query()
		query.Add("limit", fmt.Sprintf("%d", pagination.Limit))
		query.Add("offset", fmt.Sprintf("%d", pagination.Offset))
		u.RawQuery = query.Encode()
	}

	resp, err := c.HTTPClient.Get(u.String())

	if err != nil {
		return nil, fmt.Errorf("failed to make GET request[%s]: %v", u.String(), err)
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	pokemon := internal.ResultSet{}
	err = json.Unmarshal(respData, &pokemon)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response data: %v", err)
	}
	return &pokemon, nil
}

func (c *pokeClient) GetGeneration(idOrName string) (*internal.Generation, error) {
	u, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.BaseURL, c.APIVersion, "generation", idOrName))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}
	resp, err := c.HTTPClient.Get(u.String())

	if err != nil {
		return nil, fmt.Errorf("failed to make GET request[%s]: %v", u.String(), err)
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	generation := &internal.Generation{}
	err = json.Unmarshal(respData, generation)
	if err != nil {
		return nil, err
	}
	return generation, nil
}

func (c *pokeClient) GetGenerations(pagination PokeClientPagination) (*internal.ResultSet, error) {
	u, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.BaseURL, c.APIVersion, "generation"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}

	if pagination.Limit > 0 || pagination.Offset > 0 {
		query := u.Query()
		query.Add("limit", fmt.Sprintf("%d", pagination.Limit))
		query.Add("offset", fmt.Sprintf("%d", pagination.Offset))
		u.RawQuery = query.Encode()
	}

	resp, err := c.HTTPClient.Get(u.String())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	generations := internal.ResultSet{}
	err = json.Unmarshal(respData, &generations)
	if err != nil {
		return nil, err
	}
	return &generations, nil
}
