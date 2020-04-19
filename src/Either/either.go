package Either

import (
	"github.com/sghaida/fp/src/errors"
	"github.com/sghaida/fp/src/utils"
)

// EitherType gives back and EitherType that has either left value or right value
type Type struct {
	left  error
	right interface{}
}

// Either factory fpr creating new EitherType
func Either() Type {
	return Type{}
}

// Left factory to set the left (error)
// if left is being already set or right is being set it returns empty EitherType and error
// should panic if left is nil
func Left(err error) Type {
	if err == nil {
		panic("left cant be nil")
	}
	return Type{left: err}
}

// Right factory to set the right (some data)
// if left is being already set or right is being set it returns empty EitherType and error
func Right(data interface{}) Type {
	return Type{right: data}
}

// IsLeft checks if left value is being defined
// left should be defined and not nil to consider this as left sided
func (e *Type) IsLeft() bool {
	if e.left != nil {
		return true
	}
	return false
}

// IsLeft checks if right value is being defined
func (e *Type) IsRight() bool {
	if e.left == nil {
		return true
	}
	return false
}

// Left helper function for the base type assuming that the Either (Type) is empty
// if left is being already set or right is being set it returns empty EitherType and error
// if left is set to nil it will panic
func (e *Type) Left(err error) (Type, error) {
	if err == nil {
		panic("left cant be nil")
	}
	if utils.IsNilOrEmpty(e.right) && e.left == nil {
		return Type{left: err}, nil
	}
	return Type{}, errors.IsRightError
}

// Right helper function for the base type assuming that the Either (Type) is empty
// if left is being already set or right is being set it returns empty EitherType and error
func (e *Type) Right(data interface{}) (Type, error) {
	if utils.IsNilOrEmpty(e.right) && e.left == nil {
		return Type{right: data}, nil
	}
	return Type{}, errors.IsLeftError
}

// Get extracts the value from right or and if fails return error
// if right is being defined return it
// if left and right is being defined return error
// if none is defined return IsEmptyError
func (e *Type) Get() (interface{}, error) {
	if e.IsLeft() {
		return nil, e.left
	}
	if e.right == nil && e.left == nil {
		return nil, errors.IsEmptyError
	}
	return e.right, nil
}
