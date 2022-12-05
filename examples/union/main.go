package main

import (
	"context"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/fields"
	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/requests"
)

func main() {
	client := lithic.NewLithic(options.WithEnvironmentSandbox())

	res, err := client.FundingSources.New(context.TODO(), &requests.FundingSourceNewParams{
		ValidationMethodType: "BANK",
		Bank: requests.Bank{
			AccountNumber: fields.F("account_number"),
			RoutingNumber: fields.F("routing_number"),
		},
	})
	if err != nil {
		panic(err)
	}

	println(res.Token)
}
