package statement

const ParserStatementTypeDefinition = 0x0000
const ParserStatementTypeAssignment = 0x0001

type AwooParserStatement struct {
	Error error
	Type  uint16
	Data  interface{}
}
