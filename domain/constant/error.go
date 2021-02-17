package constant

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Your requested Item is not found")
	ErrConflict            = errors.New("The the desired shortcode is already in use")
	ErrBadParamInput       = errors.New("Given Param is not valid")
	ErrBadEntity           = errors.New("The shortcode fails to meet the following regexp")
)
