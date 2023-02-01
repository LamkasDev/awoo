package statement

const ParserStatementTypeDefinitionVariable = 0x0000
const ParserStatementTypeAssignment = 0x0001
const ParserStatementTypeDefinitionType = 0x0002

type AwooParserStatement struct {
	Error error
	Type  uint16
	Data  interface{}
}
