package lexer

import (
	"testing"

	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/token"
)

func testLexer(t *testing.T, l *lexer.Lexer, tests []struct {
	expectedType    token.TokenType
	expectedLiteral string
}) {
	for idx, val := range tests {
		tok := l.NextToken()

		if tok.Type != val.expectedType {
			t.Fatalf(
				"Failed at index [%d] - tokenType wrong. expected=%q but got=%q",
				idx, val.expectedType, tok.Type,
			)
		}

		if tok.Literal != val.expectedLiteral {
			t.Fatalf(
				"Failed at index [%d] - literal wrong. expected=%q but got=%q",
				idx, val.expectedLiteral, tok.Literal,
			)
		}

	}
}

func TestLexingLetVariables(t *testing.T) {
	input := `
		let x = 5;

		let y = 10;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "y"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	testLexer(t, l, tests)
}

func TestLexingConstVariables(t *testing.T) {
	input := `
		const x = 5;

		const y = 10;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.CONST, "const"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.CONST, "const"},
		{token.IDENT, "y"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	testLexer(t, l, tests)
}

func TestLexingPrefixExpressions(t *testing.T) {
	input := `
		!5;
		-5;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.EXCLAMATION, "!"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.MINUS, "-"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	testLexer(t, l, tests)
}

func TestLexingInfixExpressions(t *testing.T) {
	input := `
		5 + 5;
		5 - 5;
		5 * 5;
		5 / 5;
		5 < 5;
		5 > 5;
		5 <= 5;
		5 >= 5;
		5 == 5;
		5 != 5;
		5 < 7 < 10;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "5"},
		{token.PLUS, "+"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.MINUS, "-"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.SLASH, "/"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.LESSTHAN, "<"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.GREATERTHAN, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.LESSTHANEQ, "<="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.GREATERTHANEQ, ">="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.EQUAL, "=="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.NOTEQUAL, "!="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.LESSTHAN, "<"},
		{token.INT, "7"},
		{token.LESSTHAN, "<"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	testLexer(t, l, tests)
}

func TestLexingString(t *testing.T) {
	input := `
		"DELANG";

		"DE!!";
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.STRING, "DELANG"},
		{token.SEMICOLON, ";"},

		{token.STRING, "DE!!"},
		{token.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	testLexer(t, l, tests)
}

func TestLexingBoolean(t *testing.T) {
	input := `
		true;

		false;

		true == false;

		true != false;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},

		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},

		{token.TRUE, "true"},
		{token.EQUAL, "=="},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},

		{token.TRUE, "true"},
		{token.NOTEQUAL, "!="},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	testLexer(t, l, tests)
}

func TestLexingFunction(t *testing.T) {
	input := `
		fun(x, y) {
			return x + y;
		}

		fun(x, y) {
			return x + y;
		}();
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.FUNCTION, "fun"},
		{token.LEFTPAR, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RIGHTPAR, ")"},
		{token.LEFTBRAC, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RIGHTBRAC, "}"},

		{token.FUNCTION, "fun"},
		{token.LEFTPAR, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RIGHTPAR, ")"},
		{token.LEFTBRAC, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RIGHTBRAC, "}"},
		{token.LEFTPAR, "("},
		{token.RIGHTPAR, ")"},
		{token.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	testLexer(t, l, tests)
}

func TestLexingCallExpression(t *testing.T) {
	input := `
		logs("DELANG");

		push([1, 2], 3);

		pop([1, 2]);

		skipFirst([1, 2]);

		skipLast([1, 2]);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "logs"},
		{token.LEFTPAR, "("},
		{token.STRING, "DELANG"},
		{token.RIGHTPAR, ")"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "push"},
		{token.LEFTPAR, "("},
		{token.LEFTSQPRAC, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RIGHTSQPRAC, "]"},
		{token.COMMA, ","},
		{token.INT, "3"},
		{token.RIGHTPAR, ")"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "pop"},
		{token.LEFTPAR, "("},
		{token.LEFTSQPRAC, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RIGHTSQPRAC, "]"},
		{token.RIGHTPAR, ")"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "skipFirst"},
		{token.LEFTPAR, "("},
		{token.LEFTSQPRAC, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RIGHTSQPRAC, "]"},
		{token.RIGHTPAR, ")"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "skipLast"},
		{token.LEFTPAR, "("},
		{token.LEFTSQPRAC, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RIGHTSQPRAC, "]"},
		{token.RIGHTPAR, ")"},
		{token.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	testLexer(t, l, tests)
}

func TestLexingArray(t *testing.T) {
	input := `
		[1, 2];

		[1, 2, "Three"];

		[1, 2, "3", 4];

		[1, "test", 3, 4, 5];
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LEFTSQPRAC, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RIGHTSQPRAC, "]"},
		{token.SEMICOLON, ";"},

		{token.LEFTSQPRAC, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.COMMA, ","},
		{token.STRING, "Three"},
		{token.RIGHTSQPRAC, "]"},
		{token.SEMICOLON, ";"},

		{token.LEFTSQPRAC, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.COMMA, ","},
		{token.STRING, "3"},
		{token.COMMA, ","},
		{token.INT, "4"},
		{token.RIGHTSQPRAC, "]"},
		{token.SEMICOLON, ";"},

		{token.LEFTSQPRAC, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.STRING, "test"},
		{token.COMMA, ","},
		{token.INT, "3"},
		{token.COMMA, ","},
		{token.INT, "4"},
		{token.COMMA, ","},
		{token.INT, "5"},
		{token.RIGHTSQPRAC, "]"},
		{token.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	testLexer(t, l, tests)
}

func TestLexingArrayIndexExpression(t *testing.T) {
	input := `
		x[1];

		x["key"];
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "x"},
		{token.LEFTSQPRAC, "["},
		{token.INT, "1"},
		{token.RIGHTSQPRAC, "]"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "x"},
		{token.LEFTSQPRAC, "["},
		{token.STRING, "key"},
		{token.RIGHTSQPRAC, "]"},
		{token.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	testLexer(t, l, tests)
}

func TestLexingDict(t *testing.T) {
	input := `
		{"name": "Delang!"};

		{"name": "Delang!", "age": 1};
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LEFTBRAC, "{"},
		{token.STRING, "name"},
		{token.COLON, ":"},
		{token.STRING, "Delang!"},
		{token.RIGHTBRAC, "}"},
		{token.SEMICOLON, ";"},

		{token.LEFTBRAC, "{"},
		{token.STRING, "name"},
		{token.COLON, ":"},
		{token.STRING, "Delang!"},
		{token.COMMA, ","},
		{token.STRING, "age"},
		{token.COLON, ":"},
		{token.INT, "1"},
		{token.RIGHTBRAC, "}"},
		{token.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	testLexer(t, l, tests)
}

func TestLexingDictIndexExpression(t *testing.T) {
	input := `
		x["name"];

		x["name"] = "Delang!";
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "x"},
		{token.LEFTSQPRAC, "["},
		{token.STRING, "name"},
		{token.RIGHTSQPRAC, "]"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "x"},
		{token.LEFTSQPRAC, "["},
		{token.STRING, "name"},
		{token.RIGHTSQPRAC, "]"},
		{token.ASSIGN, "="},
		{token.STRING, "Delang!"},
		{token.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	testLexer(t, l, tests)
}

func TestLexingConditionExpression(t *testing.T) {
	input := `
		if 5 < 10: {
			return true;
		} else {
			return false;
		}

		if 1 < 2 and 2 < 3: {
			logs("DE!");
		}

		if 1 < 2 or 2 < 3: {
			logs("DE!");
		}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.INT, "5"},
		{token.LESSTHAN, "<"},
		{token.INT, "10"},
		{token.COLON, ":"},
		{token.LEFTBRAC, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RIGHTBRAC, "}"},
		{token.ELSE, "else"},
		{token.LEFTBRAC, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RIGHTBRAC, "}"},

		{token.IF, "if"},
		{token.INT, "1"},
		{token.LESSTHAN, "<"},
		{token.INT, "2"},
		{token.AND, "and"},
		{token.INT, "2"},
		{token.LESSTHAN, "<"},
		{token.INT, "3"},
		{token.COLON, ":"},
		{token.LEFTBRAC, "{"},
		{token.IDENT, "logs"},
		{token.LEFTPAR, "("},
		{token.STRING, "DE!"},
		{token.RIGHTPAR, ")"},
		{token.SEMICOLON, ";"},
		{token.RIGHTBRAC, "}"},

		{token.IF, "if"},
		{token.INT, "1"},
		{token.LESSTHAN, "<"},
		{token.INT, "2"},
		{token.OR, "or"},
		{token.INT, "2"},
		{token.LESSTHAN, "<"},
		{token.INT, "3"},
		{token.COLON, ":"},
		{token.LEFTBRAC, "{"},
		{token.IDENT, "logs"},
		{token.LEFTPAR, "("},
		{token.STRING, "DE!"},
		{token.RIGHTPAR, ")"},
		{token.SEMICOLON, ";"},
		{token.RIGHTBRAC, "}"},
	}

	l := lexer.New(input)

	testLexer(t, l, tests)
}

func TestLexingDuringLoop(t *testing.T) {
	input := `
		during x < 10: {
			logs(x);
			skip;
			break;
		}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.DURING, "during"},
		{token.IDENT, "x"},
		{token.LESSTHAN, "<"},
		{token.INT, "10"},
		{token.COLON, ":"},
		{token.LEFTBRAC, "{"},
		{token.IDENT, "logs"},
		{token.LEFTPAR, "("},
		{token.IDENT, "x"},
		{token.RIGHTPAR, ")"},
		{token.SEMICOLON, ";"},
		{token.SKIP, "skip"},
		{token.SEMICOLON, ";"},
		{token.BREAK, "break"},
		{token.SEMICOLON, ";"},
		{token.RIGHTBRAC, "}"},
	}

	l := lexer.New(input)

	testLexer(t, l, tests)
}

func TestLexingForLoop(t *testing.T) {
	input := `
		for num in [1, 2]: {
			logs(num);
		}

		for idx, num in [1, 2]: {
			logs(idx);
			logs(num);
		}

		for _, num in [1, 2]: {
			logs(num);
		}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.FOR, "for"},
		{token.IDENT, "num"},
		{token.IN, "in"},
		{token.LEFTSQPRAC, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RIGHTSQPRAC, "]"},
		{token.COLON, ":"},
		{token.LEFTBRAC, "{"},
		{token.IDENT, "logs"},
		{token.LEFTPAR, "("},
		{token.IDENT, "num"},
		{token.RIGHTPAR, ")"},
		{token.SEMICOLON, ";"},
		{token.RIGHTBRAC, "}"},

		{token.FOR, "for"},
		{token.IDENT, "idx"},
		{token.COMMA, ","},
		{token.IDENT, "num"},
		{token.IN, "in"},
		{token.LEFTSQPRAC, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RIGHTSQPRAC, "]"},
		{token.COLON, ":"},
		{token.LEFTBRAC, "{"},
		{token.IDENT, "logs"},
		{token.LEFTPAR, "("},
		{token.IDENT, "idx"},
		{token.RIGHTPAR, ")"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "logs"},
		{token.LEFTPAR, "("},
		{token.IDENT, "num"},
		{token.RIGHTPAR, ")"},
		{token.SEMICOLON, ";"},
		{token.RIGHTBRAC, "}"},

		{token.FOR, "for"},
		{token.IDENT, "_"},
		{token.COMMA, ","},
		{token.IDENT, "num"},
		{token.IN, "in"},
		{token.LEFTSQPRAC, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RIGHTSQPRAC, "]"},
		{token.COLON, ":"},
		{token.LEFTBRAC, "{"},
		{token.IDENT, "logs"},
		{token.LEFTPAR, "("},
		{token.IDENT, "num"},
		{token.RIGHTPAR, ")"},
		{token.SEMICOLON, ";"},
		{token.RIGHTBRAC, "}"},
	}

	l := lexer.New(input)

	testLexer(t, l, tests)
}

func TestLexingDecimal(t *testing.T) {
	input := `
		decimal(1.1);

		decimal(1.1) + decimal(1.1);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "decimal"},
		{token.LEFTPAR, "("},
		{token.FLOAT, "1.1"},
		{token.RIGHTPAR, ")"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "decimal"},
		{token.LEFTPAR, "("},
		{token.FLOAT, "1.1"},
		{token.RIGHTPAR, ")"},
		{token.PLUS, "+"},
		{token.IDENT, "decimal"},
		{token.LEFTPAR, "("},
		{token.FLOAT, "1.1"},
		{token.RIGHTPAR, ")"},
		{token.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	testLexer(t, l, tests)
}
