# Poké GO

A small SDK for working with Poké-API. 

Several examples exist in the [example](./example/) directory. Check them out!

## Installation

In your Go project, run `go get github.com/bradcypert/poke-go` then import the package like so:

```go
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
```

## Key Concepts

When working with this SDK, you'll start by creating a new client. The client can be configured to point to a different URL or API Version (though v2 is the only supported version at the moment). This is useful for testing (see examples in `client_test.go` for inspiration) but also useful if you're hosting your own PokeAPI server.

When making requests for resources, we use a Key interface. Key's can be generated using the `Name` function or the `ID` function.

```go
Name("pikachu")
// or
ID(151) // Mew!
```

You'll pass a Key value to the methods for fetching resources like so:

```go
generation, err := client.GetGeneration(context.Background(), v2.ID(1))
```

## Contexts

As illustrated in the previous example, this SDK supports the use of Contexts, allowing support of cancellations, timeouts, and other context-related behaviors. While the context in the example above is the Background context, you can specify a timeout with the following (as an example):

```go
ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
defer cancel()  // releases resources if slowOperation completes before timeout elapses

generation, err := client.GetGeneration(context.Background(), v2.ID(1))
```

## Testing

This library has both unit tests and integration tests. To run the tests, run `go test ./...`. The examples do not have tests, and internal does not (yet) have tests, but you should see output like so:

```
➜ go test ./...
?       github.com/bradcypert/poke-go/example/get_generation    [no test files]
?       github.com/bradcypert/poke-go/example/get_generations   [no test files]
?       github.com/bradcypert/poke-go/example/get_pikachu       [no test files]
?       github.com/bradcypert/poke-go/example/get_pokemons      [no test files]
?       github.com/bradcypert/poke-go/internal/v2       [no test files]
ok      github.com/bradcypert/poke-go/pkg/v2    0.508s
```

## Tools Used

This repo uses [JSON to Go](https://mholt.github.io/json-to-go/) to help generate the struct tree to model these API endpoints. To add support for new endpoints, go to the documentation for that endpoint, grab the JSON structure and paste it into JSON-To-Go. The types that are generated for you will be close but several may be off, especially around nullabilty. If you see an `any` type, we need to further clarify what that type truly is. Again, refering to the documentation can help make this more clear.

## Design Decisions

I've split this library up into versions that match the supported version of the PokeAPI. To support further versions, we would likely want to refactor this to encapsulate all of the version 2 functionality into a struct, and then do the same with version 1 or version 3 (or so on). When creating a client, the user would then specify the version and recieve a struct specifically for working with that API version.