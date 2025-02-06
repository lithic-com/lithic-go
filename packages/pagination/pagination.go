// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"reflect"

	"github.com/lithic-com/lithic-go/internal/apijson"
	"github.com/lithic-com/lithic-go/internal/requestconfig"
	"github.com/lithic-com/lithic-go/option"
)

type CursorPage[T any] struct {
	Data    []T            `json:"data"`
	HasMore bool           `json:"has_more,required"`
	JSON    cursorPageJSON `json:"-"`
	cfg     *requestconfig.RequestConfig
	res     *http.Response
}

// cursorPageJSON contains the JSON metadata for the struct [CursorPage[T]]
type cursorPageJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CursorPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cursorPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorPage[T]) GetNextPage() (res *CursorPage[T], err error) {
	if !r.JSON.HasMore.IsMissing() && r.HasMore == false {
		return nil, nil
	}
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

func (r *CursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorPage[T]{}
	}
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
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
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

type SinglePage[T any] struct {
	Data    []T            `json:"data"`
	HasMore bool           `json:"has_more,required"`
	JSON    singlePageJSON `json:"-"`
	cfg     *requestconfig.RequestConfig
	res     *http.Response
}

// singlePageJSON contains the JSON metadata for the struct [SinglePage[T]]
type singlePageJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinglePage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singlePageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *SinglePage[T]) GetNextPage() (res *SinglePage[T], err error) {
	if !r.JSON.HasMore.IsMissing() && r.HasMore == false {
		return nil, nil
	}
	// This page represents a response that isn't actually paginated at the API level
	// so there will never be a next page.
	cfg := (*requestconfig.RequestConfig)(nil)
	if cfg == nil {
		return nil, nil
	}
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

func (r *SinglePage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &SinglePage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type SinglePageAutoPager[T any] struct {
	page *SinglePage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewSinglePageAutoPager[T any](page *SinglePage[T], err error) *SinglePageAutoPager[T] {
	return &SinglePageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *SinglePageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *SinglePageAutoPager[T]) Current() T {
	return r.cur
}

func (r *SinglePageAutoPager[T]) Err() error {
	return r.err
}

func (r *SinglePageAutoPager[T]) Index() int {
	return r.run
}
