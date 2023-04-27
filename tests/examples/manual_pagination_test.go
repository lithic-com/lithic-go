package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
)

func TestManualPagination(t *testing.T) {
	client := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	page, err := client.Cards.List(context.TODO(), requests.CardListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, card := range page.Data {
		fmt.Printf("%+v\n", card)
	}
	// Prism mock isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, card := range page.Data {
			fmt.Printf("%+v\n", card)
		}
	}
}
