package errors

import "github.com/friendsofgo/errors"

// New encapsulates the errors library
func New(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}
