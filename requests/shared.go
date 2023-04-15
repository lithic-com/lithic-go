package requests

import (
	"github.com/lithic-com/lithic-go/core/field"
	pjson "github.com/lithic-com/lithic-go/core/json"
)

type Address struct {
	// Valid deliverable address (no PO boxes).
	Address1 field.Field[string] `json:"address1,required"`
	// Unit or apartment number (if applicable).
	Address2 field.Field[string] `json:"address2"`
	// Name of city.
	City field.Field[string] `json:"city,required"`
	// Valid country code. Only USA is currently supported, entered in uppercase ISO
	// 3166-1 alpha-3 three-character format.
	Country field.Field[string] `json:"country,required"`
	// Valid postal code. Only USA ZIP codes are currently supported, entered as a
	// five-digit ZIP or nine-digit ZIP+4.
	PostalCode field.Field[string] `json:"postal_code,required"`
	// Valid state code. Only USA state codes are currently supported, entered in
	// uppercase ISO 3166-2 two-character format.
	State field.Field[string] `json:"state,required"`
}

// MarshalJSON serializes Address into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r Address) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type ShippingAddress struct {
	// Customer's first name. This will be the first name printed on the physical card.
	FirstName field.Field[string] `json:"first_name,required"`
	// Customer's surname (family name). This will be the last name printed on the
	// physical card.
	LastName field.Field[string] `json:"last_name,required"`
	// Text to be printed on line two of the physical card. Use of this field requires
	// additional permissions.
	Line2Text field.Field[string] `json:"line2_text"`
	// Valid USPS routable address.
	Address1 field.Field[string] `json:"address1,required"`
	// Unit number (if applicable).
	Address2 field.Field[string] `json:"address2"`
	// City
	City field.Field[string] `json:"city,required"`
	// Uppercase ISO 3166-2 two character abbreviation for US and CA. Optional with a
	// limit of 24 characters for other countries.
	State field.Field[string] `json:"state,required"`
	// Postal code (formerly zipcode). For US addresses, either five-digit zipcode or
	// nine-digit "ZIP+4".
	PostalCode field.Field[string] `json:"postal_code,required"`
	// Uppercase ISO 3166-1 alpha-3 three character abbreviation.
	Country field.Field[string] `json:"country,required"`
	// Email address to be contacted for expedited shipping process purposes. Required
	// if `shipping_method` is `EXPEDITED`.
	Email field.Field[string] `json:"email"`
	// Cardholder's phone number in E.164 format to be contacted for expedited shipping
	// process purposes. Required if `shipping_method` is `EXPEDITED`.
	PhoneNumber field.Field[string] `json:"phone_number"`
}

// MarshalJSON serializes ShippingAddress into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r ShippingAddress) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}
