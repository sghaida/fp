package options

type value interface{}
type Optional struct {
	value
}

type SomeValue Optional
type NoneValue Optional

type HasValue interface {
	HasValue() bool
}

type GetValue interface {
	Get() interface{}
}

// Option factory method that accepts any type or nil
// of nil is passed is going to return NoneType else its going to return SomeType
// Option(nil), will return NoneType
// Option("test none empty") will return some string type
func Option(value interface{}) Optional {
	switch value.(type) {
	case nil:
		return None()
	default:
		return Some(value)
	}
}

// Some Factory to create Some type
func Some(value interface{}) Optional {
	return Optional{value: SomeValue{value}}
}

// None Factory to create None type
func None() Optional {
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
func (o *Optional) Get() interface{} {
	switch inner := o.value.(type) {
	case SomeValue:
		return inner.Get()
	case NoneValue:
		return inner.GetValue()
	default:
		return None()
	}
}

// getValue get the encapsulated value of the option
func (o *SomeValue) Get() interface{} {
	return o.value
}

// getValue get the encapsulated value of the option
func (o *NoneValue) GetValue() NoneValue {
	return NoneValue{"empty"}
}
