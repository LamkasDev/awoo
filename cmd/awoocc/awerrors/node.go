package awerrors

import "errors"

var ErrorCantCompileNode = errors.New("no idea how to compile node")
var ErrorFailedToCompileNode = errors.New("failed to compile node")
var ErrorCantParseNode = errors.New("no idea how to parse node")
var ErrorFailedToConstructNode = errors.New("failed to construct node")
