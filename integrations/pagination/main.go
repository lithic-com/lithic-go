package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
)

func main() {
	client := lithic.NewClient(option.WithEnvironmentSandbox())

	pager := client.Transactions.ListAutoPaging(context.TODO(), lithic.TransactionListParams{})

	tokens := make(map[string]int)

	var transaction lithic.Transaction
	for pager.Next() {
		transaction = pager.Current()
		_, exists := tokens[transaction.Token]
		if !exists {
			tokens[transaction.Token] = 0
		}
		tokens[transaction.Token] += 1
	}
	if err := pager.Err(); err != nil {
		panic(err.Error())
	}

	if len(tokens) < 0 {
		log.Fatalf("Expected at least one entry")
	}

	duplicates := make(map[string]int)
	for token, count := range tokens {
		if count > 1 {
			duplicates[token] = count
		}
	}
	if len(duplicates) > 0 {
		duplicatesJSON, err := json.MarshalIndent(duplicates, "", "  ")
		if err != nil {
			log.Fatal("Error marshaling duplicates:", err)
		}
		fmt.Println(string(duplicatesJSON))
		log.Fatalf("Found %d duplicate entries!", len(duplicates))
	}

	fmt.Println("Success!")
}
