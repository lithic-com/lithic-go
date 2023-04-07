package lithic

import (
	"github.com/lithic-com/lithic-go/core/field"
)

func F[T any](value T) field.Field[T]          { return field.Field[T]{Value: value, Present: true} }
func NullField[T any]() field.Field[T]         { return field.Field[T]{Null: true, Present: true} }
func RawField[T any](value any) field.Field[T] { return field.Field[T]{Raw: value, Present: true} }

func Float[T float32 | float64](value T) field.Field[float64] {
	return field.Field[float64]{Value: float64(value), Present: true}
}
func Int[T int | int8 | int16 | int32 | int64](value T) field.Field[int64] {
	return field.Field[int64]{Value: int64(value), Present: true}
}
func UInt[T uint | uint8 | uint16 | uint32 | uint64](value T) field.Field[uint64] {
	return field.Field[uint64]{Value: uint64(value), Present: true}
}
func Str(str string) field.Field[string] { return F(str) }
