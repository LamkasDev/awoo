package awerrors

import "errors"

var ErrorExpectedStatement = errors.New("expected statement")
var ErrorUnexpectedStatement = errors.New("unexpected statement")
var ErrorCantCompileStatement = errors.New("no idea how to compile statement")
var ErrorFailedToCompileStatement = errors.New("failed to compile statement")
var ErrorCantParseStatement = errors.New("no idea how to parse statement")
var ErrorFailedToConstructStatement = errors.New("failed to construct statement")
