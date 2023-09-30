package lexer

import (
	"testing"

	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/token"
)

func TestNextToken(t *testing.T) {
	input := `
		let five = 5;
		
		let ten = 10;
		
		let add = fun(x, y) {
			x + y;
		};
		
		let result = add(five, ten);
		
		!-/*5;

		5 < 10 > 5;

		if (5 < 10) {
			return true;
		} else {
			return false;
		}

		10 == 10;
		
		10 != 9;

		"DELANG";

		"DE!!";

		[1, 2];

		{"name": "Mostafa"};

		const x = 3;

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
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fun"},
		{token.LEFTPAR, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RIGHTPAR, ")"},
		{token.LEFTBRAC, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RIGHTBRAC, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LEFTPAR, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RIGHTPAR, ")"},
		{token.SEMICOLON, ";"},
		{token.EXCLAMATION, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LESSTHAN, "<"},
		{token.INT, "10"},
		{token.GREATERTHAN, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LEFTPAR, "("},
		{token.INT, "5"},
		{token.LESSTHAN, "<"},
		{token.INT, "10"},
		{token.RIGHTPAR, ")"},
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
		{token.INT, "10"},
		{token.EQUAL, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOTEQUAL, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.STRING, "DELANG"},
		{token.SEMICOLON, ";"},
		{token.STRING, "DE!!"},
		{token.SEMICOLON, ";"},
		{token.LEFTSQPRAC, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RIGHTSQPRAC, "]"},
		{token.SEMICOLON, ";"},
		{token.LEFTBRAC, "{"},
		{token.STRING, "name"},
		{token.COLON, ":"},
		{token.STRING, "Mostafa"},
		{token.RIGHTBRAC, "}"},
		{token.SEMICOLON, ";"},
		{token.CONST, "const"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
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
		{token.EOFILE, ""},
	}

	l := lexer.New(input)

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
