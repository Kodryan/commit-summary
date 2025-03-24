package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kodryan/commit-summary/application/parser"
	"github.com/kodryan/commit-summary/infrastructure/openapi/client"
	"github.com/kodryan/commit-summary/resources"
)

func main() {
	resources, err := resources.Get()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := client.NewClient(http.DefaultClient, client.ConnectionConfig{
		APIKey: resources.Env.GetOpenAIAPIKey(),
	})

	diff, err := parser.ParseDiff()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Diff:", diff)
	summary, err := client.GetSummary(context.Background(), diff)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Summary:", summary)
}
