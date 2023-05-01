package lithic_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/option"
)

func TestDisputeNewWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.New(context.TODO(), lithic.DisputeNewParams{Amount: lithic.F(int64(0)), CustomerFiledDate: lithic.F(time.Now()), Reason: lithic.F(lithic.DisputeNewParamsReasonAtmCashMisdispense), TransactionToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), CustomerNote: lithic.F("string")})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDisputeGet(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDisputeUpdateWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.DisputeUpdateParams{Amount: lithic.F(int64(0)), CustomerFiledDate: lithic.F(time.Now()), CustomerNote: lithic.F("string"), Reason: lithic.F(lithic.DisputeUpdateParamsReasonAtmCashMisdispense)},
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDisputeListWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.List(context.TODO(), lithic.DisputeListParams{TransactionToken: lithic.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Status: lithic.F(lithic.DisputeListParamsStatusNew), PageSize: lithic.F(int64(1)), Begin: lithic.F(time.Now()), End: lithic.F(time.Now()), StartingAfter: lithic.F("string"), EndingBefore: lithic.F("string")})
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDisputeDelete(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDisputeDeleteEvidence(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.DeleteEvidence(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDisputeInitiateEvidenceUpload(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.InitiateEvidenceUpload(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDisputeListEvidencesWithOptionalParams(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.ListEvidences(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lithic.DisputeListEvidencesParams{PageSize: lithic.F(int64(1)), Begin: lithic.F(time.Now()), End: lithic.F(time.Now()), StartingAfter: lithic.F("string"), EndingBefore: lithic.F("string")},
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDisputeGetEvidence(t *testing.T) {
	c := lithic.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Disputes.GetEvidence(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *lithic.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDisputesUploadEvidence(t *testing.T) {
	t.Skip("TODO")
}
