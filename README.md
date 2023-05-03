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

You can also explicitly set the version of the package to use by updating your `go.mod` file:

```
module your_project_name

go 1.19

require (
        github.com/lithic-com/lithic-go v0.0.1
)
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
	"github.com/lithic-com/lithic-go/option"
)

func main() {
	client := lithic.NewClient(
		option.WithAPIKey("my api key"), // defaults to os.LookupEnv("LITHIC_API_KEY")
		option.WithEnvironmentSandbox(), // defaults to option.WithEnvironmentProduction()
	)
	card, err := client.Cards.New(context.TODO(), lithic.CardNewParams{
		Type: lithic.F(lithic.CardNewParamsTypeVirtual),
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", card)
}

```

### Request Fields

Types for requests look like the following:

```go
type FooParams struct {
	Type   field.Field[string] `json:"type,required"`
	Number field.Field[int64]  `json:"number,required"`
	Name   field.Field[string] `json:"name"`
	Other  field.Field[Bar]    `json:"other"`
}

type Bar struct {
	Number field.Field[int64]  `json:"number"`
	Name   field.Field[string] `json:"name"`
}
```

For each field, you can either supply a value field with
`lithic.F(...)`, a `null` value with `lithic.Null()`, or
some raw JSON value with `lithic.Raw(...)` that you specify as a
byte slice. We also provide convenient helpers `lithic.Int(...)` and
`lithic.String(...)`.

If you do not supply a value, then we do not populate the field.

An example request may look like this:

```go
params := FooParams{
	// Normally populates this field as `"type": "foo"`
	Type: lithic.F("foo"),

	// Integer helper casts integer values and literals to field.Field[int64]
	Number: lithic.Int(12),

	// Explicitly sends this field as null, e.g., `"name": null`
	Name: lithic.Null[string](),

	// Overrides this field as `"other": "ovveride_this_field"`
	Other: lithic.Raw[Bar]("override_this_field")
}
```

Fields enable us to differentiate between zero-values, null values, empty
values, and overriden values.

If you want to add or override a field in the JSON body, then you can use the
`option.WithJSONSet(key string, value interface{})` RequestOption, which you
can read more about [here](#requestoptions). Internally, this uses
`github.com/tidwall/sjson` library, so you can compose complex access as seen
[here](https://github.com/tidwall/sjson).

### Response Objects

Response objects in this SDK have value type members. Accessing properties on
response objects is as simple as:

```go
res, err := client.Service.Foo(context.TODO())
res.Name // is just some `string` value
```

If the value received is null, not present, or invalid, the corresponding field
will simply be its empty value.

If you want to be able to tell that the value is either `null`, not present, or
invalid, we provide metadata in the `JSON` property of every response object.

```go
// This is true if `name` is _either_ not present or explicitly null
res.JSON.Name.IsNull()

// This is true if `name` is not present
res.JSON.Name.IsMissing()

// This is true if `name` is present, but not coercable
res.JSON.Name.IsInvalid()

// If needed, you can access the Raw JSON value of the field
// as a string by accessing
res.JSON.Name.Raw()
```

There may be instances where we provide experimental or private API features.
In those cases, the related features will not be exposed to
the SDK as typed fields, and are instead deserialized to an internal map. We
provide methods to get and set these json fields in API objects.

```go
// Access the JSON value as:
body := res.JSON.Extras["extra_field"].Raw()

// You can `Unmarshal` the JSON into a struct as needed
custom := struct{Foo string, Bar int64}{}
json.Unmarshal([]byte(body), &custom)
```

### RequestOptions

This library uses the functional options pattern. Functions defined in the
`option` package return a `RequestOption`, which is a closure that mutates a
`RequestConfig`. These options can be supplied to the client or at individual
requests, and they will cascade appropriately.

At each request, these closures will be run in the order that they are
supplied, after the defaults for that particular request.

For example:

```go
client := lithic.NewClient(
	// Adds a header to every request made by the client
	option.WithHeader("X-Some-Header", "custom_header_info"),
	// Adds a query param to every request made by the client
	option.WithQuery("test_token", "my_test_token"),
)

client.Cards.New(
	context.TODO(),
	...,
	// These options override the client options
	option.WithHeader("X-Some-Header", "some_other_custom_header_info"),
	option.WithQuery("test_token", "some_other_test_token"),
)

client.Cards.New(
	context.TODO(),
	...,
	// WithHeaderDel removes the header set in the client
	// from this request
	option.WithHeaderDel("X-Some-Header"),
	// WithQueryDel removes the query param set in the client
	// from this request
	option.WithQueryDel("test_token"),
)
```

### Pagination

This library provides some conveniences for working with paginated list endpoints.

You can use `.ListAutoPaging()` methods to iterate through items across all pages:

```go
iter := client.Cards.ListAutoPaging(context.TODO(), lithic.CardListParams{})
// Automatically fetches more pages as needed.
for iter.Next() {
	card := iter.Current()
	fmt.Printf("%+v\n", card)
}
if err := iter.Err(); err != nil {
	panic(err.Error())
}
```

Or you can use simple `.List()` methods to fetch a single page and receive a standard response object
with additional helper methods like `.GetNextPage()`, e.g.:

```go
page, err := client.Cards.List(context.TODO(), lithic.CardListParams{})
for page != nil {
	for _, card := range page.Data {
		fmt.Printf("%+v\n", card)
	}
	page, err = page.GetNextPage()
}
if err != nil {
	panic(err.Error())
}
```

### Errors

When the API returns a non-success status code, we return an error with type
`*lithic.Error`. This contains the `StatusCode`, `*http.Request`, and
`*http.Response` values of the request, as well as the JSON of the error body
(much like other response objects in the SDK).

To handle errors, we recommend that you use the `errors.As` pattern:

```go
_, err := client.Cards.New(context.TODO(), lithic.CardNewParams{
	Type: lithic.F(lithic.CardNewParamsTypeVirtual),
})
if err != nil {
	var apierr *lithic.Error
	if errors.As(err, &apierr) {
		println(string(apierr.DumpRequest(true)))  // Prints the serialized HTTP request
		println(string(apierr.DumpResponse(true))) // Prints the serialized HTTP response
		println(apierr.Message)                    // Invalid parameter(s): type
		println(apierr.DebuggingRequestID)         // 94d5e915-xxxx-4cee-a4f5-2xd6ebd279ac
	}
	panic(err.Error()) // GET "/cards": 400 Bad Request { ... }
}
```

When other errors occur, we return them unwrapped; for example, when HTTP
transport returns an error, we return the `*url.Error` which could wrap
`*net.OpError`.

### Timeouts

Requests do not time out by default; use context to configure a deadline for a request lifecycle.

Note that if a request is [retried](#retries), the context timeout does not start over. To set a per-retry timeout, use `option.WithRequestTimeout()`.

```go
// This sets the timeout for the request, including all the retries.
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()
client.Cards.List(
	ctx,
	lithic.CardListParams{
		PageSize: lithic.F(int64(10)),
	},
	// This sets the per-retry timeout
	option.WithRequestTimeout(20*time.Second),
)
```

## Retries

Certain errors will be automatically retried 2 times by default, with a short exponential backoff.
Connection errors (for example, due to a network connectivity problem), 409 Conflict, 429 Rate Limit,
and >=500 Internal errors will all be retried by default.

You can use the `WithMaxRetries` option to configure or disable this:

```go
// Configure the default for all requests:
client := lithic.NewClient(
	option.WithMaxRetries(0), // default is 2
)

// Override per-request:
client.Cards.List(
	context.TODO(),
	lithic.CardListParams{
		PageSize: lithic.F(int64(10)),
	},
	option.WithMaxRetries(5),
)
```

### Middleware

You may apply any middleware you wish by overriding the `http.Client` with
`option.WithClient(client)`. An example of a basic logging middleware is given
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