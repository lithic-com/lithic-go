package lithic

import (
	"github.com/lithic-com/lithic-go/internal/field"
)

func F[T any](value T) field.Field[T]          { return field.Field[T]{Value: value, Present: true} }
func NullField[T any]() field.Field[T]         { return field.Field[T]{Null: true, Present: true} }
func RawField[T any](value any) field.Field[T] { return field.Field[T]{Raw: value, Present: true} }
func Int(value int64) field.Field[int64]       { return F(value) }
func Str(str string) field.Field[string]       { return F(str) }
func Float(value float64) field.Field[float64] { return F(value) }
