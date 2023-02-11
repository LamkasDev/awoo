package parser_context

type AwooParserContextVariable struct {
	Name string
	Type uint16
}

func GetContextVariable(context *AwooParserContext, name string) (AwooParserContextVariable, bool) {
	variable, ok := context.Variables[name]
	return variable, ok
}

func SetContextVariable(context *AwooParserContext, variable AwooParserContextVariable) {
	context.Variables[variable.Name] = variable
}
