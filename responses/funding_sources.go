package responses

import (
	"time"

	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/pagination"
)

type FundingSource struct {
	// Account name identifying the funding source. This may be `null`.
	AccountName string `json:"account_name"`
	// An RFC 3339 string representing when this funding source was added to the Lithic
	// account. This may be `null`. UTC time zone.
	Created time.Time `json:"created,required" format:"date-time"`
	// The last 4 digits of the account (e.g. bank account, debit card) associated with
	// this FundingAccount. This may be null.
	LastFour string `json:"last_four,required"`
	// The nickname given to the `FundingAccount` or `null` if it has no nickname.
	Nickname string `json:"nickname"`
	// State of funding source.
	//
	// Funding source states:
	//
	//   - `ENABLED` - The funding account is available to use for card creation and
	//     transactions.
	//   - `PENDING` - The funding account is still being verified e.g. bank
	//     micro-deposits verification.
	//   - `DELETED` - The founding account has been deleted.
	State FundingSourceState `json:"state,required"`
	// A globally unique identifier for this FundingAccount.
	Token string `json:"token,required" format:"uuid"`
	// Types of funding source:
	//
	// - `DEPOSITORY_CHECKING` - Bank checking account.
	// - `DEPOSITORY_SAVINGS` - Bank savings account.
	Type FundingSourceType `json:"type,required"`
	JSON FundingSourceJSON
}

type FundingSourceJSON struct {
	AccountName pjson.Metadata
	Created     pjson.Metadata
	LastFour    pjson.Metadata
	Nickname    pjson.Metadata
	State       pjson.Metadata
	Token       pjson.Metadata
	Type        pjson.Metadata
	Raw         []byte
	Extras      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into FundingSource using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *FundingSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type FundingSourceState string

const (
	FundingSourceStateEnabled FundingSourceState = "ENABLED"
	FundingSourceStatePending FundingSourceState = "PENDING"
	FundingSourceStateDeleted FundingSourceState = "DELETED"
)

type FundingSourceType string

const (
	FundingSourceTypeDepositoryChecking FundingSourceType = "DEPOSITORY_CHECKING"
	FundingSourceTypeDepositorySavings  FundingSourceType = "DEPOSITORY_SAVINGS"
)

type FundingSourceListResponse struct {
	Data []FundingSource `json:"data,required"`
	// Page number. Will always be 1.
	Page int64 `json:"page,required"`
	// Total number of entries.
	TotalEntries int64 `json:"total_entries,required"`
	// Total number of pages. Will always be 1.
	TotalPages int64 `json:"total_pages,required"`
	JSON       FundingSourceListResponseJSON
}

type FundingSourceListResponseJSON struct {
	Data         pjson.Metadata
	Page         pjson.Metadata
	TotalEntries pjson.Metadata
	TotalPages   pjson.Metadata
	Raw          []byte
	Extras       map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into FundingSourceListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *FundingSourceListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type FundingSourcesPage struct {
	*pagination.Page[FundingSource]
}

func (r *FundingSourcesPage) FundingSource() *FundingSource {
	return r.Current()
}

func (r *FundingSourcesPage) NextPage() (*FundingSourcesPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &FundingSourcesPage{page}, nil
	}
}
