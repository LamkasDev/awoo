package awerrors

import "errors"

var ErrorFailedToSelectProgram = errors.New("failed to select program")

var ErrorFailedToCompileProgramHeader = errors.New("failed to compile program header")
var ErrorFailedToEncodeInstruction = errors.New("failed to encode instruction")
