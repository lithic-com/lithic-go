# Lithic Go API Library

<a href="https://pkg.go.dev/github.com/lithic-com/lithic-go"><img src="https://pkg.go.dev/badge/github.com/lithic-com/lithic-go.svg" alt="Go Reference"></a>

The Lithic Go library provides convenient access to the Lithic REST
API from applications written in Go.

## Installation

Within a Go module, you can just import this package and let the Go compiler
resolve this module.

```go
import (
	"github.com/lithic-com/lithic-go" // imported as lithic
)
```

Or, explicitly import this package with

```
go get -u 'github.com/lithic-com/lithic-go'
```

## Documentation

The API documentation can be found [here](https://docs.lithic.com).

## Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/lithic-com/lithic-go"
	"github.com/lithic-com/lithic-go/core/field"
	"github.com/lithic-com/lithic-go/option"
	"github.com/lithic-com/lithic-go/requests"
)

func main() {
	client := lithic.NewLithic()
	res, err := client.Cards.New(context.TODO(), &requests.CardNewParams{
		Type: lithic.F(requests.CardNewParamsTypeVirtual),
	})
	if err != nil {
		panic(err)
	}
	fmt.Sprintf("%s", res.Token)
}

```

### Request Fields

Types for requests look like the following:

```go
type FooParams struct {
	ID     fields.Field[string] `json:"id,required"`
	Number fields.Field[int64]  `json:"number,required"`
	Name   fields.Field[string] `json:"name"`
	Other  fields.Field[Bar]    `json:"other"`
}

type Bar struct {
	Number fields.Field[int64]  `json:"number"`
	Name   fields.Field[string] `json:"name"`
}
```

For each field, you can either supply a value field with
`lithic.F(...)`, a `null` value with `lithic.NullField()`, or
some raw JSON value with `lithic.RawField(...)` that you specify as a
byte slice. We also provide convenient helpers `lithic.Int(...)` and
`lithic.Str(...)`. If you do not supply a value, then we do not
populate the field. An example request may look like

```go
params := &FooParams{
	// Normally populates this field as `"id": "food_id"`
	ID: lithic.F("foo_id"),

	// Integer helper casts integer values and literals to fields.Field[int64]
	Number: lithic.Int(12),

	// Explicitly sends this field as null, e.g., `"name": null`
	Name: lithic.NullField[string](),

	// Overrides this field as `"other": "ovveride_this_field"`
	Other: lithic.RawField[Bar]("override_this_field")
}
```

If you want to add or override a field in the JSON body, then you can use the
`options.WithJSONSet(key string, value interface{})` RequestOption, which you
can read more about [here](#requestoptions). Internally, this uses
'github.com/tidwall/sjson' library, so you can compose complex access as seen
[here](https://github.com/tidwall/sjson).

### Response Objects

Response objects in this SDK have value type members. Accessing properties on
response objects is as simple as:

```go
res, err := client.Service.Foo(context.TODO())
res.Name // is just some string value
```

If null, not present, or invalid, all fields will simply be their empty values.

If you want to be able to tell that the value is either `null`, not present, or
invalid, we provide metadata in the `JSON` property of every response object.

```go
// This is true if `name` is _either_ not present or explicitly null
res.JSON.Name.IsNull()

// This is true if `name` is not present
res.JSON.Name.IsMissing()

// This is true if `name` is present, but not coercable
res.JSON.Name.IsMissing()

// If needed, you can access the Raw JSON value of the field by accessing
res.JSON.Name.Raw()
```

You can also access the JSON value of the entire object with `res.JSON.Raw`.

There may be instances where we provide experimental or private API features
for some customers. In those cases, the related features will not be exposed to
the SDK as typed fields, and are instead deserialized to an internal map. We
provide methods to get and set these json fields in API objects.

```go
// Access the JSON value as
body := res.JSON.Extras["extra_field"].Raw()

// You can `Unmarshal` the JSON into a struct as needed
custom := struct{A string, B int64}{}
json.Unmarshal(body, &custom)
```

### RequestOptions

This library uses the functional options pattern. `RequestOptions` are closures
that mutate a `RequestConfig`. These options can be supplied to the client or
at individual requests, and they will cascade appropriately.

At each request, these closures will be run in the order that they are
supplied, after the defaults for that particular request.

For example:

```go
client := Lithic.NewLithic(
	// Adds header to every request made by client
	options.WithHeader("X-Some-Header", "custom_header_info"),

	// Overrides APIkey read from environment
	options.WithAPIKey("api_key"),
)

client.Cards.New(
	context.TODO(),
	...,
	// These options override the client options
	options.WithHeader("X-Some-Header", "some_other_custom_header_info"),
	options.WithAPIKey("other_api_key"),
)
```

### Pagination

In addition to exposing standard response values, this library provides an
auto-paginating iterator for each paginated response, so you do not have to
request successive pages manually:

To iterate over all elements in a paginated list, we call the iterated
endpoint, then iterate while there is an element to be read. When `Next()` is
called, it will load the next element into `Current()` or return `false`. If we
are at the end of the current page, `Next()` will attempt to fetch the next
page.

```go
iter := client.Cards.ListAutoPager(
	context.TODO(),
	...,
)
for iter.Next() {
	item := iter.Current()
	print(item.ID)
}
if err := iter.Err(); err != nil {
	panic(err.Error())
}
```

If you want to make a simple request and handle paging yourself, you can do
that! We provide a `GetNextPage()` function which uses the same arguments and
request options as the original request, except for the arguments responsible
for the paging.

```go
page, err := client.Cards.List(context.TODO(), ...)
for page != nil {
	for _, item := range page.Data {
		print(item.ID)
	}
	page, err = page.GetNextPage()
}
if err != nil {
	panic(err.Error())
}
```

### Errors

For the errors generated by the SDK, we provide extra convenience methods for debugging.

### Middleware

You may apply any middleware you wish by overriding the `http.Client` with
`options.WithClient(client)`. An example of a basic logging middleware is given
below:

```go
TODO
```

## Status

This package is in beta. Its internals and interfaces are not stable and
subject to change without a major semver bump; please reach out if you rely on
any undocumented behavior.

We are keen for your feedback; please email us at
[sdk-feedback@lithic.com](mailto:sdk-feedback@lithic.com) or open an issue with questions, bugs, or
suggestions.

## Requirements

This library requires Go 1.18+.