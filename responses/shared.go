package responses

import (
	pjson "github.com/lithic-com/lithic-go/core/json"
)

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
