package main

import (
	"bytes"
	"context"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
)

func main() {
	client := lithic.NewClient(option.WithEnvironmentSandbox())

	pager := client.Disputes.ListAutoPaging(context.TODO(), lithic.DisputeListParams{})

	println("Listing disputes")
	var dispute lithic.Dispute
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
