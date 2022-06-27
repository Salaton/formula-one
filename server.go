package main

import (
	"context"
	"log"

	"github.com/Salaton/formula-one/pkg/presentation"
)

const (
	defaultPort = 8080
)

func main() {
	ctx := context.Background()
	err := presentation.InitializeServer(ctx, defaultPort)
	if err != nil {
		log.Fatalf("an error occured: %v", err)
	}
}
