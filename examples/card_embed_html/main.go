package main

import (
	"context"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

func main() {
	client := lithic.NewClient(option.WithEnvironmentSandbox())

	pager := client.Cards.ListAutoPaging(context.TODO(), requests.CardListParams{PageSize: lithic.Int(2)})
	println("Listing Cards")
	var card responses.Card
	for pager.Next() {
		card = pager.Current()
		println(card.Token)
	}
	if err := pager.Err(); err != nil {
		panic(err.Error())
	}

	result, err := client.Cards.GetEmbedHTML(context.TODO(), requests.EmbedRequestParams{Token: lithic.F(card.Token)})
	if err != nil {
		panic(err.Error())
	}

	println(string(result))
}
