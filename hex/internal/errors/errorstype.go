package errors

import "github.com/friendsofgo/errors"

type notFound struct {
	error
}

// NewNotFound returns an error which wraps err that satisfies IsNotFound().
func NewNotFound(format string, args ...interface{}) error {
	return &notFound{errors.Errorf(format, args...)}
}

// WrapNotFound returns an error which wraps err that satisfies IsNotFound()
func WrapNotFound(err error, format string, args ...interface{}) error {
	return &notFound{errors.Wrapf(err, format, args...)}
}

// IsNotFound reports whether err was created with NewNotFound or WrapNotFound
func IsNotFound(err error) bool {
	var target *notFound
	return errors.As(err, &target)
}

type notSavable struct {
	error
}

// NewNotSavable returns an error which wraps err that satisfies IsNotSavable().
func NewNotSavable(format string, args ...interface{}) error {
	return &notSavable{errors.Errorf(format, args...)}
}

// WrapNotSavable returns an error which wraps err that satisfies IsNotSavable()
func WrapNotSavable(err error, format string, args ...interface{}) error {
	return &notSavable{errors.Wrapf(err, format, args...)}
}

// IsNotSavable reports whether err was created with NewNotSavable or WrapNotSavable
func IsNotSavable(err error) bool {
	var target *notSavable
	return errors.As(err, &target)
}

type wrongInput struct {
	error
}

// NewWrongInput returns an error which wraps err that satisfies IsWrongInput().
func NewWrongInput(format string, args ...interface{}) error {
	return &wrongInput{errors.Errorf(format, args...)}
}

// WrapWrongInput returns an error which wraps err that satisfies IsWrongInput()
func WrapWrongInput(err error, format string, args ...interface{}) error {
	return &wrongInput{errors.Wrapf(err, format, args...)}
}

// IsWrongInput reports whether err was created with NewWrongInput or WrapWrongInput
func IsWrongInput(err error) bool {
	var target *wrongInput
	return errors.As(err, &target)
}
