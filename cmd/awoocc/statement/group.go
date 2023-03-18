package statement

type AwooParserStatementDataGroup struct {
	Body []AwooParserStatement
}

func GetStatementGroupBody(s *AwooParserStatement) []AwooParserStatement {
	return s.Data.(AwooParserStatementDataGroup).Body
}

func SetStatementGroupBody(s *AwooParserStatement, b []AwooParserStatement) {
	d := s.Data.(AwooParserStatementDataGroup)
	d.Body = b
	s.Data = d
}

func CreateStatementGroup(b []AwooParserStatement) AwooParserStatement {
	return AwooParserStatement{
		Type: ParserStatementTypeGroup,
		Data: AwooParserStatementDataGroup{
			Body: b,
		},
	}
}
