package flip

import (
	"errors"
	"fmt"
)

var (
	// ErrInternal is general internal error.
	ErrInternal = errors.New("internal error")
)

func errRequiredField(str string) error {
	return fmt.Errorf("required field %s", str)
}

func errInvalidFormatField(str string) error {
	return fmt.Errorf("invalid format field %s", str)
}

func errGTField(str, value string) error {
	return fmt.Errorf("field %s must be greater than %s", str, value)
}

func errMaxField(str, value string) error {
	return fmt.Errorf("field %s max length is %s", str, value)
}
