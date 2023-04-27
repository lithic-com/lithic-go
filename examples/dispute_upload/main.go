package main

import (
	"bytes"
	"context"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

func main() {
	client := lithic.NewClient(option.WithEnvironmentSandbox())

	pager := client.Disputes.ListAutoPaging(context.TODO(), requests.DisputeListParams{})

	println("Listing disputes")
	var dispute responses.Dispute
	for pager.Next() {
		dispute = pager.Current()
		println(dispute.Token)
	}
	if err := pager.Err(); err != nil {
		panic(err.Error())
	}

	err := client.Disputes.UploadEvidence(context.TODO(), dispute.Token, bytes.NewBuffer([]byte("some file contents")))
	if err != nil {
		panic(err.Error())
	}
}
