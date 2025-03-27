package assert

import (
	"fmt"
	"reflect"
)

const prefix = "(assertion failed)"

func NotNil[T any](value any, message string, args ...any) {
	if value == nil {
		panic(
			fmt.Sprintf(
				"%s: %q is nil: %s",
				prefix,
				reflect.TypeOf(new(T)).String(),
				fmt.Sprintf(message, args...),
			),
		)
	}
}

func NotEmpty[T any](value []T, message string, args ...any) {
	if len(value) == 0 {
		panic(
			fmt.Sprintf(
				"%s: %q is empty: %s",
				prefix,
				"[]"+reflect.TypeOf(new(T)).String(),
				fmt.Sprintf(message, args...),
			),
		)
	}
}
