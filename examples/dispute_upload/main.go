package main

import (
	"bytes"
	"context"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

func main() {
	client := lithic.NewLithic(options.WithEnvironmentSandbox())

	page, err := client.Disputes.List(context.TODO(), &requests.DisputeListParams{})
	if err != nil {
		panic(err.Error())
	}

	var dispute *responses.Dispute
	println("Listing disputes")
	for page.Next() {
		dispute = page.Dispute()
		println(dispute.Token)
	}

	err = client.Disputes.UploadEvidence(context.TODO(), dispute.Token, bytes.NewBuffer([]byte("some file contents")))
	if err != nil {
		panic(err.Error())
	}
}
