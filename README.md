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