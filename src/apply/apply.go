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
	// make sure that this is a function
	v := reflect.ValueOf(f)
	// check if d is a function to start with
	checkFuncKind(v)
	// check input | output of the func
	fType := v.Type()
	// check function take one parameter of type t and return one out put of the same type
	if fType.NumOut() != 1 || fType.NumIn() != 1 {
		panic("Function must take one input and produces one output")
	}
	// get th inbound param type
	inV := fType.In(0)
	inKind := inV.Kind()
	// get outbound param type
	outV := fType.Out(0)
	outKind := outV.Kind()
	// check types for input and output
	if inKind != outKind {
		panic("function input type != output type")
	}
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
		newSlice[i] = v.Call(in[:])[0].Interface()
	}
	// return new mapped results
	return Type{slice: newSlice}
}

func (t *Type) Get() interface{} {
	return t.slice
}

func checkFuncKind(v reflect.Value) {
	if v.Kind() != reflect.Func {
		panic("f should be a function")
	}
}
