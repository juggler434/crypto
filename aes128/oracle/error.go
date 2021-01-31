package oracle

type ShortInputError struct{}

func (sie *ShortInputError) Error() string {
	return "input too short"
}

type MalformedInputError struct{}

func (mie *MalformedInputError) Error() string {
	return "Input not in correct format"
}
