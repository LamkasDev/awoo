package statement

const ParserStatementTypeDefinitionVariable = 0x000
const ParserStatementTypeAssignment = 0x001
const ParserStatementTypeDefinitionType = 0x002
const ParserStatementTypeIf = 0x003
const ParserStatementTypeGroup = 0x004
const ParserStatementTypeFunc = 0x005
const ParserStatementTypeReturn = 0x006
const ParserStatementTypeCall = 0x007
const ParserStatementTypeFor = 0x008

type AwooParserStatement struct {
	Type uint16
	Data interface{}
}
