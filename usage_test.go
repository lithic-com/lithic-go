package lithic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
)

func TestUsage(t *testing.T) {
	client := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	card, err := client.Cards.New(context.TODO(), lithic.CardNewParams{
		Type: lithic.F(lithic.CardNewParamsTypeVirtual),
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", card)
}
