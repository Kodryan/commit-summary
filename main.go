package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/kodryan/commit-summary/application/parser"
	"github.com/kodryan/commit-summary/infrastructure/openapi/client"
	"github.com/kodryan/commit-summary/resources"
)

const (
	requestTimeout = 10 * time.Second
)

func main() {
	resources, err := resources.Get()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := client.NewClient(http.DefaultClient, client.ConnectionConfig{
		APIKey:  resources.Env.GetOpenAIAPIKey(),
		Timeout: requestTimeout,
	})

	diff, err := parser.ParseDiff()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
		return
	}

	summary, err := client.GetSummary(context.Background(), diff)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
		return
	}

	fmt.Print(summary)
	os.Stdout.Sync()
}
