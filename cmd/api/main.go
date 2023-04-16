package main

import (
	"github.com/victorrenop/raven/internal"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	wrappedRouter := internal.WrappedRouter{}
	wrappedRouter.Run(port)
}
