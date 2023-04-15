package main

import (
	"context"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
)

func main() {
	client := lithic.NewLithic(option.WithEnvironmentSandbox())

	// The request body is an interface that can be fulfilled by
	// KYC, KYCExempt, and KYB
	accountholder, err := client.AccountHolders.New(context.TODO(), requests.KYC{
		Individual: lithic.F(requests.KYCIndividual{
			Address: lithic.F(requests.Address{
				Address1:   lithic.Str("address1"),
				Address2:   lithic.Str("address2"),
				City:       lithic.Str("New York City"),
				Country:    lithic.Str("USA"),
				PostalCode: lithic.Str("10101"),
				State:      lithic.Str("NY"),
			}),
			Dob:          lithic.Str("1999-01-09"),
			Email:        lithic.Str("example@example.com"),
			FirstName:    lithic.Str("John"),
			LastName:     lithic.Str("Doe"),
			GovernmentID: lithic.Str("123-45-6789"),
			PhoneNumber:  lithic.Str("123-456-7890"),
		}),
		TosTimestamp: lithic.Str("2023-04-13T10:40:20.23Z"),
		Workflow:     lithic.F(requests.KYCWorkflowKYCBasic),
	})
	if err != nil {
		panic(err.Error())
	}

	println(accountholder.Token)
}
