// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lithic

import (
	"github.com/lithic-com/lithic-go/option"
)

// TransactionMonitoringService contains methods and other services that help with
// interacting with the lithic API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionMonitoringService] method instead.
type TransactionMonitoringService struct {
	Options []option.RequestOption
	Cases   *TransactionMonitoringCaseService
	Queues  *TransactionMonitoringQueueService
}

// NewTransactionMonitoringService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransactionMonitoringService(opts ...option.RequestOption) (r *TransactionMonitoringService) {
	r = &TransactionMonitoringService{}
	r.Options = opts
	r.Cases = NewTransactionMonitoringCaseService(opts...)
	r.Queues = NewTransactionMonitoringQueueService(opts...)
	return
}
