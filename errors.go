package phyphox

import "errors"

var ErrBufferParse = errors.New("cannot parse the buffer correctly")
var ErrBufferVarNotExist = errors.New("buffer does not contain this variable")
