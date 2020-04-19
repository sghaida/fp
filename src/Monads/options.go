package Monads

type value interface{}
type Optional struct {
	value
}

type SomeValue Optional
type NoneValue Optional

type None interface {
	None() Optional
}
type Some interface {
	Some(value interface{}) Optional
}

type HasValue interface {
	HasValue() bool
}

type GetValue interface {
	GetValue() interface{}
}

// Option factory method that accepts any type or nil
// of nil is passed is going to return NoneType else its going to return SomeType
// Option(nil), will return NoneType
// Option("test none empty") will return some string type
func Option(value interface{}) Optional {
	switch value.(type) {
	case nil:
		return Optional{value: NoneValue{value: "empty"}}
	default:
		return Optional{value: SomeValue{value}}
	}
}

// Some is a factory method that create SomeValue type that encapsulates the value
func (o Optional) Some(value interface{}) Optional {
	return Optional{value: SomeValue{value: value}}
}

// None is a factory that makes NoneValue type
func (o *Optional) None() Optional {
	return Optional{value: NoneValue{"empty"}}
}

// HasValue check if option has some value which returns bool
func (o *Optional) HasValue() bool {
	switch o.value.(type) {
	case SomeValue:
		return true
	default:
		return false
	}
}

// getValue get the encapsulated value of the option
func (o *Optional) GetValue() interface{} {
	switch inner := o.value.(type) {
	case SomeValue:
		return inner.GetValue()
	case NoneValue:
		return inner.GetValue()
	default:
		return NoneValue{"empty"}
	}
}

// getValue get the encapsulated value of the option
func (o *SomeValue) GetValue() interface{} {
	return o.value
}

// getValue get the encapsulated value of the option
func (o *NoneValue) GetValue() NoneValue {
	return NoneValue{"empty"}
}
