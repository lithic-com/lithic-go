package option

import (
	"log"
	"net/http"
	"net/url"

	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/tidwall/sjson"
)

type RequestOption = func(*requestconfig.RequestConfig) error

func WithBaseURL(base string) RequestOption {
	u, err := url.Parse(base)
	if err != nil {
		log.Fatalf("failed to parse BaseURL: %s\n", err)
	}
	return func(r *requestconfig.RequestConfig) error {
		r.BaseURL = u
		return nil
	}
}

func WithHTTPClient(client *http.Client) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.HTTPClient = client
		return nil
	}
}

func WithMaxRetries(retries int) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.MaxRetries = retries
		return nil
	}
}

func WithHeader(key, value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.Request.Header[key] = []string{value}
		return nil
	}
}
func WithHeaderAdd(key, value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.Request.Header[key] = append(r.Request.Header[key], value)
		return nil
	}
}
func WithHeaderDel(key string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		delete(r.Request.Header, key)
		return nil
	}
}

func WithQuery(key, value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		query := r.Request.URL.Query()
		query.Set(key, value)
		r.Request.URL.RawQuery = query.Encode()
		return nil
	}
}
func WithQueryAdd(key, value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		query := r.Request.URL.Query()
		query.Add(key, value)
		r.Request.URL.RawQuery = query.Encode()
		return nil
	}
}
func WithQueryDel(key string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		query := r.Request.URL.Query()
		query.Del(key)
		r.Request.URL.RawQuery = query.Encode()
		return nil
	}
}

func WithJSONSet(key string, value interface{}) RequestOption {
	return func(r *requestconfig.RequestConfig) (err error) {
		r.Buffer, err = sjson.SetBytes(r.Buffer, key, value)
		return err
	}
}
func WithJSONDel(key string) RequestOption {
	return func(r *requestconfig.RequestConfig) (err error) {
		r.Buffer, err = sjson.DeleteBytes(r.Buffer, key)
		return err
	}
}

func WithResponseBodyInto(dst any) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.ResponseBodyInto = dst
		return nil
	}
}

func WithResponseInto(dst **http.Response) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.ResponseInto = dst
		return nil
	}
}

func WithAPIKey(key string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.APIKey = key
		return r.Apply(WithHeader("Authorization", r.APIKey))
	}
}

func WithEnvironmentProduction() RequestOption {
	return WithBaseURL("https://api.lithic.com/v1/")
}

func WithEnvironmentSandbox() RequestOption {
	return WithBaseURL("https://sandbox.lithic.com/v1/")
}

func WithWebhookSecret(value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.WebhookSecret = value
		return nil
	}
}
