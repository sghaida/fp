package apply

import (
	"reflect"
)

type Type struct {
	d interface{}
}

// Lift Factory method to encapsulate the type within Type
func Lift(data interface{}) Type {
	return Type{d: data}
}

// Map
// panic if d type and f input type is not the same type
func (t *Type) Apply(f interface{}) Type {
	// just in case the content of the d is string,
	if s, ok := t.d.([]string); ok {
		return t.applyString(s, f)
	}
	// make sure that this is a function
	fv := reflect.ValueOf(f)
	// check if d is a function to start with
	checkFuncKind(fv)
	// check input | output of the func
	fType := fv.Type()
	// check function take one parameter of type t and return one out put of the same type
	if fType.NumOut() != 1 || fType.NumIn() != 1 {
		panic("Function must take one input and produces one output")
	}
	switch reflect.ValueOf(t.d).Kind() {
	case reflect.Slice:
		return t.applySlice(fv)
	default:
		return t.applyGeneric(fv)
	}

}

// Get return the encapsulated type value
func (t *Type) Get() interface{} {
	return t.d
}

// applyString deals with the map for strings
func (t *Type) applyString(s []string, f interface{}) Type {
	if fn, ok := f.(func(string) string); ok {
		newSlice := make([]string, len(s))
		for i, st := range s {
			newSlice[i] = fn(st)
		}
		return Type{d: newSlice}
	}
	panic("function signature is not supported")
}

// applyGeneric applies a function for all types other than slice and string
func (t *Type) applyGeneric(fv reflect.Value) Type {
	in := []reflect.Value{reflect.ValueOf(t.d)}
	out := fv.Call(in[:])[0].Interface()
	return Type{d: out}
}

// applySlice encapsulates the logic for the map func for all slices
func (t *Type) applySlice(fv reflect.Value) Type {
	// create empty d based on the original d
	origSlice := reflect.ValueOf(t.d)
	newSlice := make([]interface{}, origSlice.Len())
	// apply the method then append the item to the new d
	var in [1]reflect.Value
	for i := 0; i < origSlice.Len(); i++ {
		in[0] = origSlice.Index(i)
		newSlice[i] = fv.Call(in[:])[0].Interface()
	}
	// return new mapped results
	return Type{d: newSlice}
}

// getInnerType gets the inner type of t
// this could panic in the cae of none slice and based on that a recover will return for none slice
func (t *Type) getInnerType() reflect.Kind {
	defer func() reflect.Kind {
		if err := recover(); err != nil {
		}
		return reflect.TypeOf(t.d).Kind()
	}()
	return reflect.TypeOf(t.d).Elem().Kind()
}

// checkFuncKind checks if v is a function type
func checkFuncKind(v reflect.Value) {
	if v.Kind() != reflect.Func {
		panic("f should be a function")
	}
}
