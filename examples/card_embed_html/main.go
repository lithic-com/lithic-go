package main

import (
	"context"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

func main() {
	client := lithic.NewLithic(options.WithEnvironmentSandbox())

	cards, err := client.Cards.List(context.TODO(), &requests.CardListParams{PageSize: lithic.Int(2)})
	println("Listing Cards")
	var card *responses.Card
	for cards.Next() {
		card = cards.Card()
		println(card.Token)
	}
	if cards.Err() != nil {
		panic(cards.Err().Error())
	}

	result, err := client.Cards.GetEmbedHTML(context.TODO(), &requests.EmbedRequestParams{Token: lithic.F(card.Token)})
	if err != nil {
		panic(err.Error())
	}

	println(string(result))
}
