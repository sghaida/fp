package errors

type EitherError string

func (e EitherError) Error() string { return string(e) }

const (
	IsLeftError  = EitherError("the Either is left sided")
	IsRightError = EitherError("the Either is right sided")
	IsEmptyError = EitherError("the Either is empty")
)
