package awerrors

import "errors"

var ErrorCantCompileOperator = errors.New("no idea how to compile operator")
var ErrorFailedToCompileOperator = errors.New("failed to compile operator")
