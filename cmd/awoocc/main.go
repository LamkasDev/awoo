package main

import (
	"flag"
	"os"
	"os/user"
	"path"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_run"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_run"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement_compile"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement_parse"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/flags"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/LamkasDev/awoo-emu/cmd/common/paths"
)

func main() {
	u, _ := user.Current()
	defaultInput := path.Join(u.HomeDir, "Documents", "awoo", "data", "input.awoo")
	defaultOutput := path.Join(u.HomeDir, "Documents", "awoo", "data", "obj", "input.awoobj")

	var input string
	var output string
	var quiet bool
	flag.StringVar(&input, "i", defaultInput, "path to input .awoo file")
	flag.StringVar(&output, "o", defaultOutput, "path to output .awooobj file")
	flag.BoolVar(&quiet, "q", false, "set to disable log")
	flag.Parse()
	logger.AwooLoggerEnabled = !quiet
	input, output = paths.ResolvePaths(input, ".awoo", output, ".awoobj")
	flags.ResolveColor()

	file, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	lexSettings := lexer.AwooLexerSettings{
		Tokens: token.SetupTokenMap(),
		Mappings: lexer.AwooLexerMappings{
			PrintNode: map[uint16]lexer.AwooPrintNode{
				node.ParserNodeTypeIdentifier:  lexer.PrintNodeIdentifier,
				node.ParserNodeTypeType:        lexer.PrintNodeType,
				node.ParserNodeTypePointer:     lexer.PrintNodePointer,
				node.ParserNodeTypePrimitive:   lexer.PrintNodePrimitive,
				node.ParserNodeTypeExpression:  lexer.PrintNodeExpression,
				node.ParserNodeTypeNegative:    lexer.PrintNodeNegative,
				node.ParserNodeTypeReference:   lexer.PrintNodeReference,
				node.ParserNodeTypeDereference: lexer.PrintNodeDereference,
				node.ParserNodeTypeCall:        lexer.PrintNodeCall,
				node.ParserNodeTypeTypeArray:   lexer.PrintNodeTypeArray,
				node.ParserNodeTypeArrayIndex:  lexer.PrintNodeArrayIndex,
			},
		},
	}
	lex := lexer.SetupLexer(lexSettings)
	lexer.LoadLexer(&lex, []rune(string(file)))
	lexRes := lexer.RunLexer(&lex)

	parSettings := parser.AwooParserSettings{
		Lexer: lexSettings,
		Mappings: parser.AwooParserMappings{
			Statement: map[uint16]parser.AwooParseStatement{
				token.TokenTypeType:            statement_parse.ConstructStatementDefinitionVariable,
				token.TokenTypeIdentifier:      statement_parse.ConstructStatementIdentifier,
				token.TokenTypeTypeDefinition:  statement_parse.ConstructStatementDefinitionType,
				token.TokenTypeIf:              statement_parse.ConstructStatementIf,
				token.TokenTypeFunc:            statement_parse.ConstructStatementFunc,
				token.TokenTypeReturn:          statement_parse.ConstructStatementReturn,
				token.TokenOperatorDereference: statement_parse.ConstructStatementIdentifier,
				token.TokenTypeFor:             statement_parse.ConstructStatementFor,
			},
			NodeExpression: map[uint16]parser.AwooParseNodeExpression{
				token.TokenTypeBracketRight:       statement_parse.ConstructExpressionEndBracket,
				token.TokenOperatorAddition:       statement_parse.ConstructExpressionUnary,
				token.TokenOperatorSubstraction:   statement_parse.ConstructExpressionUnary,
				token.TokenOperatorMultiplication: statement_parse.ConstructExpressionUnary,
				token.TokenOperatorDivision:       statement_parse.ConstructExpressionUnary,
				token.TokenOperatorEq:             statement_parse.ConstructExpressionEquality,
				token.TokenOperatorNotEq:          statement_parse.ConstructExpressionNotEquality,
				token.TokenOperatorLT:             statement_parse.ConstructExpressionComparison,
				token.TokenOperatorGT:             statement_parse.ConstructExpressionComparison,
				token.TokenOperatorAnd:            statement_parse.ConstructExpressionAnd,
				token.TokenOperatorOr:             statement_parse.ConstructExpressionOr,
			},
			NodeValue: map[uint16]parser.AwooParseNodeValue{
				token.TokenTypePrimitive:        statement_parse.CreateNodePrimitiveSafe,
				token.TokenTypeIdentifier:       statement_parse.CreateNodeIdentifierSafe,
				token.TokenTypeBracketCurlyLeft: statement_parse.CreateNodeArray,
			},
			PrintStatement: map[uint16]parser.AwooPrintStatement{
				statement.ParserStatementTypeDefinitionVariable: parser.PrintStatementDefinitionVariable,
				statement.ParserStatementTypeAssignment:         parser.PrintStatementAssignment,
				statement.ParserStatementTypeDefinitionType:     parser.PrintStatementDefinitionType,
				statement.ParserStatementTypeIf:                 parser.PrintStatementIf,
				statement.ParserStatementTypeFunc:               parser.PrintStatementFunc,
			},
		},
	}
	par := parser.SetupParser(parSettings, lex.Context)
	parser.LoadParser(&par, lexRes)
	parRes := parser_run.RunParser(&par)

	compSettings := compiler.AwooCompilerSettings{
		Path:   output,
		Parser: parSettings,
		Mappings: compiler.AwooCompilerMappings{
			Statement: map[uint16]compiler.AwooCompileStatement{
				statement.ParserStatementTypeDefinitionVariable: statement_compile.CompileStatementDefinition,
				statement.ParserStatementTypeAssignment:         statement_compile.CompileStatementAssignment,
				statement.ParserStatementTypeDefinitionType: func(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, s statement.AwooParserStatement) error {
					return nil
				},
				statement.ParserStatementTypeIf:     statement_compile.CompileStatementIf,
				statement.ParserStatementTypeGroup:  statement_compile.CompileStatementGroup,
				statement.ParserStatementTypeFunc:   statement_compile.CompileStatementFunc,
				statement.ParserStatementTypeReturn: statement_compile.CompileStatementReturn,
				statement.ParserStatementTypeCall:   statement_compile.CompileStatementCall,
				statement.ParserStatementTypeFor:    statement_compile.CompileStatementFor,
			},
			NodeExpression: map[uint16]compiler.AwooCompileNodeExpression{
				token.TokenOperatorAddition:       statement_compile.HandleNodeExpressionLeftRight(instructions.AwooInstructionADD),
				token.TokenOperatorSubstraction:   statement_compile.HandleNodeExpressionLeftRight(instructions.AwooInstructionSUB),
				token.TokenOperatorMultiplication: statement_compile.HandleNodeExpressionLeftRight(instructions.AwooInstructionMUL),
				token.TokenOperatorDivision:       statement_compile.HandleNodeExpressionLeftRight(instructions.AwooInstructionDIV),
				token.TokenOperatorEqEq:           statement_compile.CompileNodeExpressionEqEq,
				token.TokenOperatorNotEq:          statement_compile.CompileNodeExpressionNotEq,
				token.TokenOperatorLT:             statement_compile.HandleNodeExpressionLeftRight(instructions.AwooInstructionSLT),
				token.TokenOperatorLTEQ:           statement_compile.CompileNodeExpressionLTEQ,
				token.TokenOperatorGT:             statement_compile.HandleNodeExpressionRightLeft(instructions.AwooInstructionSLT),
				token.TokenOperatorGTEQ:           statement_compile.CompileNodeExpressionGTEQ,
				token.TokenOperatorLS:             statement_compile.HandleNodeExpressionLeftRight(instructions.AwooInstructionSLL),
				token.TokenOperatorRS:             statement_compile.HandleNodeExpressionLeftRight(instructions.AwooInstructionSRL),
				token.TokenOperatorAnd:            statement_compile.HandleNodeExpressionLeftRight(instructions.AwooInstructionAND),
				token.TokenOperatorOr:             statement_compile.HandleNodeExpressionLeftRight(instructions.AwooInstructionOR),
			},
			NodeValue: map[uint16]compiler.AwooCompileNodeValue{
				node.ParserNodeTypeIdentifier:  statement_compile.CompileNodeIdentifier,
				node.ParserNodeTypePrimitive:   statement_compile.CompileNodePrimitive,
				node.ParserNodeTypeExpression:  statement_compile.CompileNodeExpression,
				node.ParserNodeTypeNegative:    statement_compile.CompileNodeNegative,
				node.ParserNodeTypeReference:   statement_compile.CompileNodeReference,
				node.ParserNodeTypeDereference: statement_compile.CompileNodeDereference,
				node.ParserNodeTypeCall:        statement_compile.CompileNodeCall,
				node.ParserNodeTypeArrayIndex:  statement_compile.CompileNodeArrayIndex,
				node.ParserNodeTypeArray:       statement_compile.CompileNodeArray,
			},
		},
	}
	comp := compiler.SetupCompiler(compSettings, par.Context)
	compiler.LoadCompiler(&comp, parRes)
	compiler_run.RunCompiler(&comp)
}
