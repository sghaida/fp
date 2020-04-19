package errors

type EitherError string

func (e EitherError) Error() string { return string(e) }

const (
	IsLeftError  = EitherError("the either is left sided")
	IsRightError = EitherError("the either is right sided")
	IsEmptyError = EitherError("the either is empty")
)
