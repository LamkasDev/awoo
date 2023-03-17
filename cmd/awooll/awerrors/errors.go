package awerrors

import "errors"

var ErrorFailedToSelectProgram = errors.New("failed to select program")

var ErrorIllegalCharacter = errors.New("illegal characters")
var ErrorExpectedToken = errors.New("expected one of")
var ErrorUnexpectedToken = errors.New("unexpected token")
var ErrorFailedToCreateToken = errors.New("failed to create token")
var ErrorNoMoreTokens = errors.New("no more tokens")

var ErrorPrimitiveOverflow = errors.New("primitive overflow")
var ErrorPrimitiveUnderflow = errors.New("primitive underflow")

var ErrorExpectedStatement = errors.New("expected statement")
var ErrorUnexpectedStatement = errors.New("unexpected statement")

var ErrorUnknownVariable = errors.New("unknown variable")
var ErrorAlreadyDefinedVariable = errors.New("already defined variable")
var ErrorFailedToGetVariableFromScope = errors.New("failed to get variable from scope")
var ErrorFailedToPushVariableIntoScope = errors.New("failed to push variable into scope")
var ErrorUnknownFunction = errors.New("unknown function")
var ErrorAlreadyDefinedFunction = errors.New("already defined function")
var ErrorFailedToGetFunctionFromScope = errors.New("failed to get function from scope")
var ErrorFailedToPushFunctionIntoScope = errors.New("failed to push function into scope")

var ErrorCantCompileOperator = errors.New("no idea how to compile operator")
var ErrorFailedToCompileOperator = errors.New("failed to compile operator")

var ErrorCantCompileNode = errors.New("no idea how to compile node")
var ErrorFailedToCompileNode = errors.New("failed to compile node")
var ErrorCantParseNode = errors.New("no idea how to parse node")
var ErrorFailedToConstructNode = errors.New("failed to construct node")

var ErrorFailedToConstructExpression = errors.New("failed to construct expression")

var ErrorCantCompileStatement = errors.New("no idea how to compile statement")
var ErrorFailedToCompileStatement = errors.New("failed to compile statement")
var ErrorCantParseStatement = errors.New("no idea how to parse statement")
var ErrorFailedToConstructStatement = errors.New("failed to construct statement")

var ErrorFailedToCompileProgramHeader = errors.New("failed to compile program header")
var ErrorFailedToEncodeInstruction = errors.New("failed to encode instruction")
