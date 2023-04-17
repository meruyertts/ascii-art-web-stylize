package printascii

import "errors"

var (
	ErrFont     error = errors.New("only thinkertoy, standard and shadow files are available")
	ErrNonAscii error = errors.New("non-ASCII character was entered")
	ErrTxtFile  error = errors.New("the txt file has been changed")
	ErrString   error = errors.New("string was not entered")
	ErrRead     error = errors.New("file is not readable")
)
