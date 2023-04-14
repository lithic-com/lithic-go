package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
)

func TestUsage(t *testing.T) {
	client := lithic.NewLithic(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	card, err := client.Cards.New(context.TODO(), &requests.CardNewParams{
		Type: lithic.F(requests.CardNewParamsTypeVirtual),
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", card)
}
