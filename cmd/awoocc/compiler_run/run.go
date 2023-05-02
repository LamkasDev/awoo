package compiler_run

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"os"
	"path/filepath"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/dependency"
	awooElf "github.com/LamkasDev/awoo-emu/cmd/awoocc/elf"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_run"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement_compile"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement_parse"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/LamkasDev/awoo-emu/cmd/common/paths"
	"github.com/jwalton/gchalk"
	"golang.org/x/exp/maps"
)

func RunCompilerFull(context map[string]elf.AwooElf, input paths.AwooPath, output paths.AwooPath, passedOutput string) {
	if _, ok := context[input.Absolute]; ok {
		return
	}
	file, err := os.ReadFile(input.Absolute)
	if err != nil {
		panic(err)
	}

	lexSettings := lexer.AwooLexerSettings{
		Path:   input.Absolute,
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
	lex := lexer.NewLexer(lexSettings)
	lexer.LoadLexer(&lex, []rune(string(file)))
	lexRes := lexer.RunLexer(&lex)

	dependencyContext := map[string]elf.AwooElf{}
	ProcessDependencies(dependencyContext, lexRes, input, passedOutput)
	maps.Copy(context, dependencyContext)

	celf := elf.NewAwooElf(filepath.Base(input.Absolute), elf.AwooElfTypeObject)
	for _, dependency := range dependencyContext {
		elf.MergeSimpleSymbolTable(celf.SymbolTable.External, dependency.SymbolTable.Internal, 0)
	}

	parSettings := parser.AwooParserSettings{
		Lexer: lexSettings,
		Mappings: parser.AwooParserMappings{
			Statement: map[uint16]parser.AwooParseStatement{
				token.TokenTypeType:              statement_parse.ConstructStatementDefinitionVariable,
				token.TokenTypeIdentifier:        statement_parse.ConstructStatementIdentifier,
				token.TokenKeywordTypeDefinition: statement_parse.ConstructStatementDefinitionType,
				token.TokenKeywordIf:             statement_parse.ConstructStatementIf,
				token.TokenKeywordFunc:           statement_parse.ConstructStatementFunc,
				token.TokenKeywordReturn:         statement_parse.ConstructStatementReturn,
				token.TokenOperatorDereference:   statement_parse.ConstructStatementIdentifier,
				token.TokenKeywordFor:            statement_parse.ConstructStatementFor,
				token.TokenKeywordImport:         statement_parse.ConstructStatementImport,
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
				statement.ParserStatementTypeIf:                 parser.PrintStatementIf,
				statement.ParserStatementTypeFunc:               parser.PrintStatementFunc,
			},
		},
	}
	par := parser.SetupParser(parSettings, lex.Context)
	parser.LoadParser(&par, lexRes)
	parser_run.LoadParserSymbols(&par, &celf)
	parRes := parser_run.RunParser(&par)

	compSettings := compiler.AwooCompilerSettings{
		Path:   output.Absolute,
		Parser: parSettings,
		Mappings: compiler.AwooCompilerMappings{
			Statement: map[uint16]compiler.AwooCompileStatement{
				statement.ParserStatementTypeDefinitionVariable: statement_compile.CompileStatementDefinition,
				statement.ParserStatementTypeAssignment:         statement_compile.CompileStatementAssignment,
				statement.ParserStatementTypeIf:                 statement_compile.CompileStatementIf,
				statement.ParserStatementTypeGroup:              statement_compile.CompileStatementGroup,
				statement.ParserStatementTypeFunc:               statement_compile.CompileStatementFunc,
				statement.ParserStatementTypeReturn:             statement_compile.CompileStatementReturn,
				statement.ParserStatementTypeCall:               statement_compile.CompileStatementCall,
				statement.ParserStatementTypeFor:                statement_compile.CompileStatementFor,
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
			InstructionTable: instructions.SetupInstructionTable(),
		},
	}
	comp := compiler.SetupCompiler(compSettings, par.Context)
	compiler.LoadCompiler(&comp, parRes)
	RunCompiler(&comp, &celf)

	context[input.Absolute] = celf
}

func ProcessDependencies(context map[string]elf.AwooElf, current lexer.AwooLexerResult, parentPath paths.AwooPath, passedOutput string) {
	deps, err := dependency.ResolveDependencies(&current)
	if err != nil {
		panic(err)
	}
	for _, dependency := range deps {
		anchor := filepath.Dir(parentPath.Absolute)
		input := paths.AwooPath{
			Absolute: filepath.Join(anchor, dependency),
			Anchor:   parentPath.Anchor,
		}
		output := paths.ResolveOutputPath(input, passedOutput, ".awoobj")
		RunCompilerFull(context, input, output, passedOutput)
	}
}

func RunCompiler(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf) {
	RunCompilerDry(ccompiler, celf)
	celf.SymbolTable.External = map[string]elf.AwooElfSymbolTableEntry{}

	if err := os.MkdirAll(filepath.Dir(ccompiler.Settings.Path), 0644); err != nil {
		panic(err)
	}
	file, err := os.Create(ccompiler.Settings.Path)
	if err != nil {
		panic(err)
	}
	var data bytes.Buffer
	if err := gob.NewEncoder(&data).Encode(celf); err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	writer.Write(data.Bytes())
	writer.Flush()
	file.Close()
}

func RunCompilerDry(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf) {
	logger.LogExtra(gchalk.Yellow("\n> Compiler\n"))

	for ok := true; ok; ok = compiler.AdvanceCompiler(ccompiler) {
		start := len(celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents)
		parser.PrintStatement(&ccompiler.Settings.Parser, &ccompiler.Context.Parser, &ccompiler.Current)
		if err := statement_compile.CompileStatement(ccompiler, celf, ccompiler.Current); err != nil {
			panic(err)
		}
		end := len(celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents)
		compiler.PrintNewCompile(ccompiler, &ccompiler.Current, celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents[start:end])
	}
	elf.AlignSections(celf)
	if err := awooElf.AlignSymbols(ccompiler, celf); err != nil {
		panic(err)
	}
}
