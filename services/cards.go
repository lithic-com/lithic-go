package services

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/lithic-com/lithic-go/options"
	"github.com/lithic-com/lithic-go/pagination"
	"github.com/lithic-com/lithic-go/requests"
	"github.com/lithic-com/lithic-go/responses"
)

type CardService struct {
	Options []options.RequestOption
}

func NewCardService(opts ...options.RequestOption) (r *CardService) {
	r = &CardService{}
	r.Options = opts
	return
}

// Create a new virtual or physical card. Parameters `pin`, `shipping_address`, and
// `product_id` only apply to physical cards.
func (r *CardService) New(ctx context.Context, body *requests.CardNewParams, opts ...options.RequestOption) (res *responses.Card, err error) {
	opts = append(r.Options[:], opts...)
	path := "cards"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Get card configuration such as spend limit and state.
func (r *CardService) Get(ctx context.Context, card_token string, opts ...options.RequestOption) (res *responses.Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s", card_token)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update the specified properties of the card. Unsupplied properties will remain
// unchanged. `pin` parameter only applies to physical cards.
//
// _Note: setting a card to a `CLOSED` state is a final action that cannot be
// undone._
func (r *CardService) Update(ctx context.Context, card_token string, body *requests.CardUpdateParams, opts ...options.RequestOption) (res *responses.Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s", card_token)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List cards.
func (r *CardService) List(ctx context.Context, query *requests.CardListParams, opts ...options.RequestOption) (res *responses.CardsPage, err error) {
	opts = append(r.Options, opts...)
	path := "cards"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.CardsPage{
		Page: &pagination.Page[responses.Card]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

// Handling full card PANs and CVV codes requires that you comply with the Payment
// Card Industry Data Security Standards (PCI DSS). Some clients choose to reduce
// their compliance obligations by leveraging our embedded card UI solution
// documented below.
//
// In this setup, PANs and CVV codes are presented to the end-user via a card UI
// that we provide, optionally styled in the customer's branding using a specified
// css stylesheet. A user's browser makes the request directly to api.lithic.com,
// so card PANs and CVVs never touch the API customer's servers while full card
// data is displayed to their end-users. The response contains an HTML document.
// This means that the url for the request can be inserted straight into the `src`
// attribute of an iframe.
//
// ```html
// <iframe
//
//	id="card-iframe"
//	src="https://sandbox.lithic.com/v1/embed/card?embed_request=eyJjc3MiO...;hmac=r8tx1..."
//	allow="clipboard-write"
//	class="content"
//
// ></iframe>
// ```
//
// You should compute the request payload on the server side. You can render it (or
// the whole iframe) on the server or make an ajax call from your front end code,
// but **do not ever embed your API key into front end code, as doing so introduces
// a serious security vulnerability**.
func (r *CardService) Embed(ctx context.Context, query *requests.CardEmbedParams, opts ...options.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Accept", "text/html")}, opts...)
	path := "embed/card"
	err = options.ExecuteNewRequest(ctx, "GET", path, query, &res, opts...)
	return
}

func (r *CardService) GetEmbedHTML(ctx context.Context, body *requests.EmbedRequestParams, opts ...options.RequestOption) (res []byte, err error) {
	opts = append(r.Options, opts...)
	buf, err := body.MarshalJSON()
	if err != nil {
		return nil, err
	}
	cfg, err := options.NewRequestConfig(ctx, "GET", "embed/card", nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	mac := hmac.New(sha256.New, []byte(cfg.APIKey))
	mac.Write(buf)
	sign := mac.Sum(nil)
	err = cfg.Apply(
		options.WithHeader("Accept", "text/html"),
		options.WithQuery("hmac", base64.StdEncoding.EncodeToString(sign)),
		options.WithQuery("embed_request", base64.StdEncoding.EncodeToString(buf)),
	)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	return

}

// Handling full card PANs and CVV codes requires that you comply with the Payment
// Card Industry Data Security Standards (PCI DSS). Some clients choose to reduce
// their compliance obligations by leveraging our embedded card UI solution
// documented below.
//
// In this setup, PANs and CVV codes are presented to the end-user via a card UI
// that we provide, optionally styled in the customer's branding using a specified
// css stylesheet. A user's browser makes the request directly to api.lithic.com,
// so card PANs and CVVs never touch the API customer's servers while full card
// data is displayed to their end-users. The response contains an HTML document.
// This means that the url for the request can be inserted straight into the `src`
// attribute of an iframe.
//
// ```html
// <iframe
//
//	id="card-iframe"
//	src="https://sandbox.lithic.com/v1/embed/card?embed_request=eyJjc3MiO...;hmac=r8tx1..."
//	allow="clipboard-write"
//	class="content"
//
// ></iframe>
// ```
//
// You should compute the request payload on the server side. You can render it (or
// the whole iframe) on the server or make an ajax call from your front end code,
// but **do not ever embed your API key into front end code, as doing so introduces
// a serious security vulnerability**.
func (r *CardService) GetEmbedURL(ctx context.Context, body *requests.EmbedRequestParams, opts ...options.RequestOption) (res *url.URL, err error) {
	buf, err := body.MarshalJSON()
	if err != nil {
		return nil, err
	}
	cfg, err := options.NewRequestConfig(ctx, "GET", "embed/card", nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	mac := hmac.New(sha256.New, []byte(cfg.APIKey))
	mac.Write(buf)
	sign := mac.Sum(nil)
	err = cfg.Apply(
		options.WithQuery("hmac", base64.StdEncoding.EncodeToString(sign)),
		options.WithQuery("embed_request", base64.StdEncoding.EncodeToString(buf)),
	)
	if err != nil {
		return nil, err
	}
	return cfg.Request.URL, nil

}

// Allow your cardholders to directly add payment cards to the device's digital
// wallet (e.g. Apple Pay) with one touch from your app.
//
// This requires some additional setup and configuration. Please
// [Contact Us](https://lithic.com/contact) or your Customer Success representative
// for more information.
func (r *CardService) Provision(ctx context.Context, card_token string, body *requests.CardProvisionParams, opts ...options.RequestOption) (res *responses.CardProvisionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s/provision", card_token)
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Initiate print and shipment of a duplicate physical card.
//
// Only applies to cards of type `PHYSICAL`.
func (r *CardService) Reissue(ctx context.Context, card_token string, body *requests.CardReissueParams, opts ...options.RequestOption) (res *responses.Card, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cards/%s/reissue", card_token)
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}
