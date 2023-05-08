package awerrors

import "errors"

var ErrorPrimitiveOverflow = errors.New("primitive overflow")
var ErrorPrimitiveUnderflow = errors.New("primitive underflow")
