package main

import (
	"context"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
)

func main() {
	client := lithic.NewClient(option.WithEnvironmentSandbox())

	// The request body is an interface that can be fulfilled by
	// KYC, KYCExempt, and KYB
	accountholder, err := client.AccountHolders.New(context.TODO(), lithic.KYCParam{
		Individual: lithic.F(lithic.KYCIndividualParam{
			Address: lithic.F(lithic.AddressParam{
				Address1:   lithic.String("address1"),
				Address2:   lithic.String("address2"),
				City:       lithic.String("New York City"),
				Country:    lithic.String("USA"),
				PostalCode: lithic.String("10101"),
				State:      lithic.String("NY"),
			}),
			Dob:          lithic.String("1999-01-09"),
			Email:        lithic.String("example@example.com"),
			FirstName:    lithic.String("John"),
			LastName:     lithic.String("Doe"),
			GovernmentID: lithic.String("123-45-6789"),
			PhoneNumber:  lithic.String("123-456-7890"),
		}),
		TosTimestamp: lithic.String("2023-04-13T10:40:20.23Z"),
		Workflow:     lithic.F(lithic.KYCWorkflowKYCBasic),
	})
	if err != nil {
		panic(err.Error())
	}

	println(accountholder.Token)
}
