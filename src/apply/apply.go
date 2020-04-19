package apply

import (
	"reflect"
)

type Type struct {
	slice interface{}
}

// Lift Factory method to encapsulate the type within Type
func Lift(data interface{}) Type {
	// check if this is a slice
	s := reflect.ValueOf(data)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}
	return Type{slice: data}
}

// Map
// panic if map function input and output dont have the same type
// panic if slice type and f input type is not the same type
func (t *Type) Apply(f interface{}) Type {
	// just in case the content of the slice is string,
	if s, ok := t.slice.([]string); ok {
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
	return t.applyGeneric(fType, fv)
}

func (t *Type) Get() interface{} {
	return t.slice
}

// applyString deals with the map for strings
func (t *Type) applyString(s []string, f interface{}) Type {
	if fn, ok := f.(func(string) string); ok {
		newSlice := make([]string, len(s))
		for i, st := range s {
			newSlice[i] = fn(st)
		}
		return Type{slice: newSlice}
	}
	panic("function signature is not supported")
}

// applyGeneric encapsulates the logic for the map func for all types other than string
func (t *Type) applyGeneric(tp reflect.Type, fv reflect.Value) Type {
	// get th inbound param type
	inV := tp.In(0)
	inKind := inV.Kind()

	if reflect.TypeOf(t.slice).Elem().Kind() != inKind {
		panic("slice type != map function input type")
	}
	// create empty slice based on the original slice
	origSlice := reflect.ValueOf(t.slice)
	newSlice := make([]interface{}, origSlice.Len())
	// apply the method then append the item to the new slice
	var in [1]reflect.Value
	for i := 0; i < origSlice.Len(); i++ {
		in[0] = origSlice.Index(i)
		newSlice[i] = fv.Call(in[:])[0].Interface()
	}
	// return new mapped results
	return Type{slice: newSlice}
}

func checkFuncKind(v reflect.Value) {
	if v.Kind() != reflect.Func {
		panic("f should be a function")
	}
}
