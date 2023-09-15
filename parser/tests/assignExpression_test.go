package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/parser"
)

func TestAssignExpression(t *testing.T) {
	input := "x = 5;"

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 2, len(program.Statements))
	}

	if program.Statements[0].String() != "x = 5;" {
		t.Fatalf("program.Statements[0] is not 'x = 5;', got=%q", program.Statements[0].String())
	}

	if program.Statements[0].TokenLiteral() != "x" {
		t.Fatalf("program.Statements[0].TokenLiteral() is not 'x', got=%q", program.Statements[0].TokenLiteral())
	}
}
