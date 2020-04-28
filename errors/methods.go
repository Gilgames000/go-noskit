package errors

import "github.com/pkg/errors"

func New(msg string) error {
	return errors.New(msg)
}

func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func Errorf(format string, args ...string) error {
	return errors.Errorf(format, args)
}
