package awerrors

import "errors"

var ErrorUnknownVariable = errors.New("unknown variable")
var ErrorAlreadyDefinedVariable = errors.New("already defined variable")
var ErrorFailedToGetVariableFromScope = errors.New("failed to get variable from scope")
var ErrorFailedToPushVariableIntoScope = errors.New("failed to push variable into scope")
var ErrorUnknownFunction = errors.New("unknown function")
var ErrorAlreadyDefinedFunction = errors.New("already defined function")
var ErrorFailedToGetFunctionFromScope = errors.New("failed to get function from scope")
var ErrorFailedToPushFunctionIntoScope = errors.New("failed to push function into scope")
