package parser_context

type AwooParserContextFunction struct {
	Name      string
	Arguments []AwooParserContextVariable
}

func GetContextFunction(context *AwooParserContext, name string) (AwooParserContextFunction, bool) {
	function, ok := context.Functions[name]
	return function, ok
}

func SetContextFunction(context *AwooParserContext, function AwooParserContextFunction) {
	context.Functions[function.Name] = function
}
