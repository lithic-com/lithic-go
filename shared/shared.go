// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/param"
)

type Address struct {
	// Valid deliverable address (no PO boxes).
	Address1 string `json:"address1,required"`
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
	// Unit or apartment number (if applicable).
	Address2 string      `json:"address2"`
	JSON     addressJSON `json:"-"`
}

// addressJSON contains the JSON metadata for the struct [Address]
type addressJSON struct {
	Address1    apijson.Field
	City        apijson.Field
	Country     apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Address2    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Address) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressJSON) RawJSON() string {
	return r.raw
}

type AddressParam struct {
	// Valid deliverable address (no PO boxes).
	Address1 param.Field[string] `json:"address1,required"`
	// Name of city.
	City param.Field[string] `json:"city,required"`
	// Valid country code. Only USA is currently supported, entered in uppercase ISO
	// 3166-1 alpha-3 three-character format.
	Country param.Field[string] `json:"country,required"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State param.Field[string] `json:"state,required"`
	// Unit or apartment number (if applicable).
	Address2 param.Field[string] `json:"address2"`
}

func (r AddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CarrierParam struct {
	// QR code url to display on the card carrier
	QrCodeURL param.Field[string] `json:"qr_code_url"`
}

func (r CarrierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ShippingAddressParam struct {
	// Valid USPS routable address.
	Address1 param.Field[string] `json:"address1,required"`
	// City
	City param.Field[string] `json:"city,required"`
	// Uppercase ISO 3166-1 alpha-3 three character abbreviation.
	Country param.Field[string] `json:"country,required"`
	// Customer's first name. This will be the first name printed on the physical card.
	// The combined length of `first_name` and `last_name` may not exceed 25
	// characters.
	FirstName param.Field[string] `json:"first_name,required"`
	// Customer's surname (family name). This will be the last name printed on the
	// physical card. The combined length of `first_name` and `last_name` may not
	// exceed 25 characters.
	LastName param.Field[string] `json:"last_name,required"`
	// Postal code (formerly zipcode). For US addresses, either five-digit zipcode or
	// nine-digit "ZIP+4".
	PostalCode param.Field[string] `json:"postal_code,required"`
	// Uppercase ISO 3166-2 two character abbreviation for US and CA. Optional with a
	// limit of 24 characters for other countries.
	State param.Field[string] `json:"state,required"`
	// Unit number (if applicable).
	Address2 param.Field[string] `json:"address2"`
	// Email address to be contacted for expedited shipping process purposes. Required
	// if `shipping_method` is `EXPEDITED`.
	Email param.Field[string] `json:"email"`
	// Text to be printed on line two of the physical card. Use of this field requires
	// additional permissions.
	Line2Text param.Field[string] `json:"line2_text"`
	// Cardholder's phone number in E.164 format to be contacted for expedited shipping
	// process purposes. Required if `shipping_method` is `EXPEDITED`.
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r ShippingAddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
