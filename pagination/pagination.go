package pagination

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	pjson "github.com/lithic-com/lithic-go/core/json"
	"github.com/lithic-com/lithic-go/options"
)

type Paginated[T any] interface {
	NextPageConfig() *options.RequestConfig
	// shared
	GetResponse() PaginatedResponse[T]
	GetRawResponse() *http.Response
	Current() *T
	HasNext() bool
	Next() bool
	Err() error
	Index() int
}

type PaginatedResponse[T any] interface {
	GetItems() []T
	GetItem(index int) *T
	GetLength() int
}

type Page[T any] struct {
	Config       options.RequestConfig
	Options      []options.RequestOption
	runningIndex int
	index        int
	err          error
	current      *T
	raw          *http.Response
	res          *PageResponse[T]
}

func (r *Page[T]) NextPageConfig() *options.RequestConfig {
	if r.res == nil {
		return nil
	}
	currentPage := r.res.Page
	if currentPage >= r.res.TotalPages {
		return nil
	}
	cfg := r.Config.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	return cfg
}

func (r *Page[T]) Fire() (err error) {
	var res PageResponse[T]
	var raw *http.Response
	r.Config.ResponseInto = &raw
	r.Config.ResponseBodyInto = &res

	err = r.Config.Execute()
	r.res = &res
	r.raw = raw
	if err != nil {
		r.err = err
		return err
	}
	return nil
}

// Attempts to read the next page. If there is no next page, it returns nil.
// Otherwise, it returns a pointer to a _new_ page and leaves the current page
// as-is.
func (r *Page[T]) NextPage() (page *Page[T], err error) {
	cfg := r.NextPageConfig()
	if cfg == nil {
		return nil, nil
	}
	page = &Page[T]{
		Config:       *cfg,
		runningIndex: r.runningIndex,
	}
	err = page.Fire()
	if err != nil {
		return nil, err
	}
	return page, nil
}

// Current returns the element currently read. Calling Current before calling Next
// is a programming error and will cause your program to crash.
func (r *Page[T]) Current() *T {
	// The index is subtracted by one to differentiate between 0 (unset) and 1 (first element)
	// Calling Current before calling Next is a programming error
	return r.current
}

// Attempts to read the next element, if successful returns true, false otherwise.
// This function will cross page-boundaries and attempt to read more pages if it
// can, changing the internal page value.
func (r *Page[T]) Next() bool {
	if !r.HasNext() {
		return false
	}
	if r.index >= r.res.GetLength() {
		page, err := r.NextPage()
		if err != nil {
			r.err = err
			return false
		}
		if page == nil {
			return false
		}
		*r = *page
		return r.Next()
	}

	r.current = r.res.GetItem(r.index)
	r.runningIndex += 1
	r.index += 1
	return true
}

func (r *Page[T]) HasNext() bool {
	if r.err != nil || r.res.GetLength() == 0 {
		return false
	}
	if r.index >= r.res.GetLength() {
		return r.NextPageConfig() != nil
	}
	return true
}

func (r *Page[T]) Err() error {
	return r.err
}

// Returns the cummulative index of the 'current' element. Zero-indexed.
func (r *Page[T]) Index() int {
	return r.runningIndex - 1
}

func (r *Page[T]) GetResponse() PaginatedResponse[T] {
	return r.res
}

func (r *Page[T]) GetRawResponse() *http.Response {
	return r.raw
}

type PageResponse[T any] struct {
	Data []T `json:"data,required"`
	// Page number.
	Page int64 `json:"page,required"`
	// Total number of entries.
	TotalEntries int64 `json:"total_entries,required"`
	// Total number of pages.
	TotalPages int64 `json:"total_pages,required"`
	JSON       PageResponseJSON
}

type PageResponseJSON struct {
	Data         pjson.Metadata
	Page         pjson.Metadata
	TotalEntries pjson.Metadata
	TotalPages   pjson.Metadata
	Raw          []byte
	Extras       map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PageResponse[T] using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *PageResponse[T]) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

var _ PaginatedResponse[any] = (*PageResponse[any])(nil)

func (r *PageResponse[T]) GetItems() []T {
	return r.Data
}

func (r *PageResponse[T]) GetItem(index int) *T {
	return &r.GetItems()[index]
}

func (r *PageResponse[T]) GetLength() int {
	return len(r.GetItems())
}

type CursorPage[T any] struct {
	Config       options.RequestConfig
	Options      []options.RequestOption
	runningIndex int
	index        int
	err          error
	current      *T
	raw          *http.Response
	res          *CursorPageResponse[T]
}

func (r *CursorPage[T]) NextPageConfig() *options.RequestConfig {
	if r.res == nil {
		return nil
	}
	items := r.GetResponse().GetItems()
	if items == nil || len(items) == 0 {
		return nil
	}
	cfg := r.Config.Clone(r.Config.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("Token")
	cfg.Apply(options.WithQuery("starting_after", field.Interface().(string)))
	return cfg
}

func (r *CursorPage[T]) Fire() (err error) {
	var res CursorPageResponse[T]
	var raw *http.Response
	r.Config.ResponseInto = &raw
	r.Config.ResponseBodyInto = &res

	err = r.Config.Execute()
	r.res = &res
	r.raw = raw
	if err != nil {
		r.err = err
		return err
	}
	return nil
}

// Attempts to read the next page. If there is no next page, it returns nil.
// Otherwise, it returns a pointer to a _new_ page and leaves the current page
// as-is.
func (r *CursorPage[T]) NextPage() (page *CursorPage[T], err error) {
	cfg := r.NextPageConfig()
	if cfg == nil {
		return nil, nil
	}
	page = &CursorPage[T]{
		Config:       *cfg,
		runningIndex: r.runningIndex,
	}
	err = page.Fire()
	if err != nil {
		return nil, err
	}
	return page, nil
}

// Current returns the element currently read. Calling Current before calling Next
// is a programming error and will cause your program to crash.
func (r *CursorPage[T]) Current() *T {
	// The index is subtracted by one to differentiate between 0 (unset) and 1 (first element)
	// Calling Current before calling Next is a programming error
	return r.current
}

// Attempts to read the next element, if successful returns true, false otherwise.
// This function will cross page-boundaries and attempt to read more pages if it
// can, changing the internal page value.
func (r *CursorPage[T]) Next() bool {
	if !r.HasNext() {
		return false
	}
	if r.index >= r.res.GetLength() {
		page, err := r.NextPage()
		if err != nil {
			r.err = err
			return false
		}
		if page == nil {
			return false
		}
		*r = *page
		return r.Next()
	}

	r.current = r.res.GetItem(r.index)
	r.runningIndex += 1
	r.index += 1
	return true
}

func (r *CursorPage[T]) HasNext() bool {
	if r.err != nil || r.res.GetLength() == 0 {
		return false
	}
	if r.index >= r.res.GetLength() {
		return r.NextPageConfig() != nil
	}
	return true
}

func (r *CursorPage[T]) Err() error {
	return r.err
}

// Returns the cummulative index of the 'current' element. Zero-indexed.
func (r *CursorPage[T]) Index() int {
	return r.runningIndex - 1
}

func (r *CursorPage[T]) GetResponse() PaginatedResponse[T] {
	return r.res
}

func (r *CursorPage[T]) GetRawResponse() *http.Response {
	return r.raw
}

type CursorPageResponse[T any] struct {
	Data    []T  `json:"data,required"`
	HasMore bool `json:"has_more,required"`
	JSON    CursorPageResponseJSON
}

type CursorPageResponseJSON struct {
	Data    pjson.Metadata
	HasMore pjson.Metadata
	Raw     []byte
	Extras  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CursorPageResponse[T] using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CursorPageResponse[T]) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

var _ PaginatedResponse[any] = (*CursorPageResponse[any])(nil)

func (r *CursorPageResponse[T]) GetItems() []T {
	return r.Data
}

func (r *CursorPageResponse[T]) GetItem(index int) *T {
	return &r.GetItems()[index]
}

func (r *CursorPageResponse[T]) GetLength() int {
	return len(r.GetItems())
}
