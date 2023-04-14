package responses

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/option"
)

type Page[T any] struct {
	Data []T `json:"data,required"`
	// Page number.
	Page int64 `json:"page,required"`
	// Total number of entries.
	TotalEntries int64 `json:"total_entries,required"`
	// Total number of pages.
	TotalPages int64 `json:"total_pages,required"`
	JSON       PageJSON
	cfg        *option.RequestConfig
	res        *http.Response
}

type PageJSON struct {
	Data         pjson.Metadata
	Page         pjson.Metadata
	TotalEntries pjson.Metadata
	TotalPages   pjson.Metadata
	Raw          []byte
	Extras       map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Page[T] using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Page[T]) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *Page[T]) GetNextPage() (res *Page[T], err error) {
	currentPage := r.Page
	if currentPage >= r.TotalPages {
		return nil, nil
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *Page[T]) SetPageConfig(cfg *option.RequestConfig, res *http.Response) {
	r.cfg = cfg
	r.res = res
}

type PageAutoPager[T any] struct {
	page *Page[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPageAutoPager[T any](page *Page[T], err error) *PageAutoPager[T] {
	return &PageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageAutoPager[T]) Next() bool {
	if len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageAutoPager[T]) Err() error {
	return r.err
}

func (r *PageAutoPager[T]) Index() int {
	return r.run
}

type CursorPage[T any] struct {
	Data    []T  `json:"data,required"`
	HasMore bool `json:"has_more,required"`
	JSON    CursorPageJSON
	cfg     *option.RequestConfig
	res     *http.Response
}

type CursorPageJSON struct {
	Data    pjson.Metadata
	HasMore pjson.Metadata
	Raw     []byte
	Extras  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CursorPage[T] using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CursorPage[T]) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *CursorPage[T]) GetNextPage() (res *CursorPage[T], err error) {
	items := r.Data
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("Token")
	cfg.Apply(option.WithQuery("starting_after", field.Interface().(string)))
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorPage[T]) SetPageConfig(cfg *option.RequestConfig, res *http.Response) {
	r.cfg = cfg
	r.res = res
}

type CursorPageAutoPager[T any] struct {
	page *CursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewCursorPageAutoPager[T any](page *CursorPage[T], err error) *CursorPageAutoPager[T] {
	return &CursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorPageAutoPager[T]) Next() bool {
	if len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorPageAutoPager[T]) Index() int {
	return r.run
}

type Address struct {
	// Valid deliverable address (no PO boxes).
	Address1 string `json:"address1,required"`
	// Unit or apartment number (if applicable).
	Address2 string `json:"address2"`
	// Name of city.
	City string `json:"city,required"`
	// Valid country code. Only USA is currently supported, entered in uppercase ISO
	// 3166-1 alpha-3 three-character format.
	Country string `json:"country,required"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode string `json:"postal_code,required"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State string `json:"state,required"`
	JSON  AddressJSON
}

type AddressJSON struct {
	Address1   pjson.Metadata
	Address2   pjson.Metadata
	City       pjson.Metadata
	Country    pjson.Metadata
	PostalCode pjson.Metadata
	State      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Address using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Address) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ShippingAddress struct {
	// Customer's first name. This will be the first name printed on the physical card.
	FirstName string `json:"first_name,required"`
	// Customer's surname (family name). This will be the last name printed on the
	// physical card.
	LastName string `json:"last_name,required"`
	// Text to be printed on line two of the physical card. Use of this field requires
	// additional permissions.
	Line2Text string `json:"line2_text"`
	// Valid USPS routable address.
	Address1 string `json:"address1,required"`
	// Unit number (if applicable).
	Address2 string `json:"address2"`
	// City
	City string `json:"city,required"`
	// Uppercase ISO 3166-2 two character abbreviation for US and CA. Optional with a
	// limit of 24 characters for other countries.
	State string `json:"state,required"`
	// Postal code (formerly zipcode). For US addresses, either five-digit zipcode or
	// nine-digit "ZIP+4".
	PostalCode string `json:"postal_code,required"`
	// Uppercase ISO 3166-1 alpha-3 three character abbreviation.
	Country string `json:"country,required"`
	// Email address to be contacted for expedited shipping process purposes. Required
	// if `shipping_method` is `EXPEDITED`.
	Email string `json:"email"`
	// Cardholder's phone number in E.164 format to be contacted for expedited shipping
	// process purposes. Required if `shipping_method` is `EXPEDITED`.
	PhoneNumber string `json:"phone_number"`
	JSON        ShippingAddressJSON
}

type ShippingAddressJSON struct {
	FirstName   pjson.Metadata
	LastName    pjson.Metadata
	Line2Text   pjson.Metadata
	Address1    pjson.Metadata
	Address2    pjson.Metadata
	City        pjson.Metadata
	State       pjson.Metadata
	PostalCode  pjson.Metadata
	Country     pjson.Metadata
	Email       pjson.Metadata
	PhoneNumber pjson.Metadata
	Raw         []byte
	Extras      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ShippingAddress using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ShippingAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
