package awerrors

import "errors"

var ErrorIllegalCharacter = errors.New("illegal characters")

var ErrorExpectedToken = errors.New("expected one of")
var ErrorUnexpectedToken = errors.New("unexpected token")
var ErrorFailedToCreateToken = errors.New("failed to create token")
var ErrorNoMoreTokens = errors.New("no more tokens")
