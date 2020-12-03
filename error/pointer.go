package error

//Create a named type for our new error type
type errorString string

//Implement the error interface
func (e errorString) Error() string{
	return string(e)
}
//New Creates interface value of type error
func New(text string) error {
	return errorString(text)
}
