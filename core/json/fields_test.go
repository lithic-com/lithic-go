package json

import (
	"testing"
	"time"

	"github.com/lithic-com/lithic-go/core/fields"
)

type Struct struct {
	A string `json:"a"`
	B int64  `json:"b"`
}

type FieldStruct struct {
	A fields.Field[string]    `json:"a"`
	B fields.Field[int64]     `json:"b"`
	C fields.Field[Struct]    `json:"c"`
	D fields.Field[time.Time] `json:"d" format:"date"`
	E fields.Field[time.Time] `json:"e" format:"date-time"`
	F fields.Field[int64]     `json:"f"`
}

func TestFieldMarshal(t *testing.T) {
	tests := map[string]struct {
		value    interface{}
		expected string
	}{
		"null_string": {fields.Field[string]{Present: true, Null: true}, "null"},
		"null_int":    {fields.Field[int]{Present: true, Null: true}, "null"},
		"null_int64":  {fields.Field[int64]{Present: true, Null: true}, "null"},
		"null_struct": {fields.Field[Struct]{Present: true, Null: true}, "null"},

		"string": {fields.Field[string]{Present: true, Value: "string"}, `"string"`},
		"int":    {fields.Field[int]{Present: true, Value: 123}, "123"},
		"int64":  {fields.Field[int64]{Present: true, Value: int64(123456789123456789)}, "123456789123456789"},
		"struct": {fields.Field[Struct]{Present: true, Value: Struct{A: "yo", B: 123}}, `{"a":"yo","b":123}`},

		"string_raw": {fields.Field[int]{Present: true, Raw: "string"}, `"string"`},
		"int_raw":    {fields.Field[int]{Present: true, Raw: 123}, "123"},
		"int64_raw":  {fields.Field[int]{Present: true, Raw: int64(123456789123456789)}, "123456789123456789"},
		"struct_raw": {fields.Field[int]{Present: true, Raw: Struct{A: "yo", B: 123}}, `{"a":"yo","b":123}`},

		"field_struct": {
			FieldStruct{
				A: fields.Field[string]{Present: true, Value: "hello"},
				B: fields.Field[int64]{Present: true, Value: int64(12)},
				D: fields.Field[time.Time]{Present: true, Value: time.Date(2023, time.March, 18, 14, 47, 38, 0, time.UTC)},
				E: fields.Field[time.Time]{Present: true, Value: time.Date(2023, time.March, 18, 14, 47, 38, 0, time.UTC)},
			},
			`{"a":"hello","b":12,"d":"2023-03-18","e":"2023-03-18T14:47:38Z"}`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			b, err := Marshal(test.value)
			if err != nil {
				t.Fatalf("didn't expect error %v", err)
			}
			if string(b) != test.expected {
				t.Fatalf("expected %s, received %s", test.expected, string(b))
			}
		})
	}
}
