package Monads

type EitherType struct {
	left  error
	right struct{}
}

func Either() EitherType {
	return EitherType{}
}

func (e *EitherType) IsLeft() bool {
	panic("not implemented yet")
}

func (e *EitherType) IsRight() bool {
	panic("not implemented yet")
}

func (e *EitherType) Left(err error) EitherType {
	// TODO check right
	panic("not implemented yet")
}

func (e *EitherType) Right(data struct{}) EitherType {
	// TODO check left
	panic("not implemented yet")
}

func (e *EitherType) Get() interface{} {
	// TODO check left and Right
	panic("not implemented yet")
}
