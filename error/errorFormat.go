package error

import "errors"


var err error


var (
	WrongFormat = errors.New("\"Error... Wrong Format, Enter Valid Format\"")
	InvalidPosition = errors.New("Error... Enter Valid Position, ")
	InvalidNoRooks = errors.New("\"Error... Invalid Number of rooks placed,\n rooks size is more than the allowed size,\n Enter Valid Position\"")
)
