package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
)

func TestParseConstWithInteger(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      int64
	}{
		{"const x = 5;", "x", 5},
		{"const y = 10;", "y", 10},
		{"const num = 1234;", "num", 1234},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := program.Statements[0]

		if !testConstStatement(t, statement, val.expectedIdentifier) {
			return
		}

		value := statement.(*ast.ConstStatement).Value

		if !testLiteralExpression(t, value, val.expectedValue) {
			return
		}

		identifier := statement.(*ast.ConstStatement).Name

		if !testLiteralExpression(t, identifier, val.expectedIdentifier) {
			return
		}
	}
}

func TestParseConstWithBoolean(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      bool
	}{
		{"const x = true;", "x", true},
		{"const y = false;", "y", false},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := program.Statements[0]

		if !testConstStatement(t, statement, val.expectedIdentifier) {
			return
		}

		value := statement.(*ast.ConstStatement).Value

		if !testLiteralExpression(t, value, val.expectedValue) {
			return
		}

		identifier := statement.(*ast.ConstStatement).Name

		if !testLiteralExpression(t, identifier, val.expectedIdentifier) {
			return
		}
	}
}

func TestParseConstWithFloat(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      float64
	}{
		{"const x = 5.5;", "x", 5.5},
		{"const y = 10.5;", "y", 10.5},
		{"const PI = 3.14159;", "PI", 3.14159},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := program.Statements[0]

		if !testConstStatement(t, statement, val.expectedIdentifier) {
			return
		}

		value := statement.(*ast.ConstStatement).Value

		if !testLiteralExpression(t, value, val.expectedValue) {
			return
		}

		identifier := statement.(*ast.ConstStatement).Name

		if !testLiteralExpression(t, identifier, val.expectedIdentifier) {
			return
		}
	}
}

func TestParseConstWithString(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      string
	}{
		{`const x = "Hello World!";`, "x", "Hello World!"},
		{`const lang = "Delang!";`, "lang", "Delang!"},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := program.Statements[0]

		if !testConstStatement(t, statement, val.expectedIdentifier) {
			return
		}

		value := statement.(*ast.ConstStatement).Value

		if !testLiteralExpression(t, value, val.expectedValue) {
			return
		}

		identifier := statement.(*ast.ConstStatement).Name

		if !testLiteralExpression(t, identifier, val.expectedIdentifier) {
			return
		}
	}
}

func TestParseConstWithBooleanPrefixExpression(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      struct {
			operator string
			right    bool
		}
	}{
		{"const x = !true;", "x", struct {
			operator string
			right    bool
		}{
			"!", true,
		}},
		{"const y = !false;", "y", struct {
			operator string
			right    bool
		}{
			"!", false,
		}},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := program.Statements[0]

		if !testConstStatement(t, statement, val.expectedIdentifier) {
			return
		}

		constValue := statement.(*ast.ConstStatement).Value
		operator := val.expectedValue.operator
		right := val.expectedValue.right

		if !testPrefixExpression(t, constValue, operator, right) {
			return
		}

		identifier := statement.(*ast.ConstStatement).Name

		if !testLiteralExpression(t, identifier, val.expectedIdentifier) {
			return
		}
	}
}

func TestParseConstWithIntegerPrefixExpression(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      struct {
			operator string
			right    int64
		}
	}{
		{
			"const x = -5;",
			"x",
			struct {
				operator string
				right    int64
			}{
				"-", 5,
			},
		},
		{
			"const y = -10;",
			"y",
			struct {
				operator string
				right    int64
			}{
				"-", 10,
			},
		},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := program.Statements[0]

		if !testConstStatement(t, statement, val.expectedIdentifier) {
			return
		}

		constValue := statement.(*ast.ConstStatement).Value
		operator := val.expectedValue.operator
		right := val.expectedValue.right

		if !testPrefixExpression(t, constValue, operator, right) {
			return
		}

		identifier := statement.(*ast.ConstStatement).Name

		if !testLiteralExpression(t, identifier, val.expectedIdentifier) {
			return
		}
	}
}

func TestParseConstWithFloatPrefixExpression(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      struct {
			operator string
			right    float64
		}
	}{
		{"const x = -5.5;", "x", struct {
			operator string
			right    float64
		}{
			"-", 5.5,
		}},
		{"const y = -10.5;", "y", struct {
			operator string
			right    float64
		}{
			"-", 10.5,
		}},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := program.Statements[0]

		if !testConstStatement(t, statement, val.expectedIdentifier) {
			return
		}

		constValue := statement.(*ast.ConstStatement).Value
		operator := val.expectedValue.operator
		right := val.expectedValue.right

		if !testPrefixExpression(t, constValue, operator, right) {
			return
		}

		identifier := statement.(*ast.ConstStatement).Name

		if !testLiteralExpression(t, identifier, val.expectedIdentifier) {
			return
		}
	}
}

func TestParseConstWithIntegerInfixExpression(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      struct {
			left     int64
			operator string
			right    int64
		}
	}{
		{
			"const x = 5 + 5;",
			"x",
			struct {
				left     int64
				operator string
				right    int64
			}{
				5, "+", 5,
			},
		},
		{
			"const y = 5 - 5;",
			"y",
			struct {
				left     int64
				operator string
				right    int64
			}{
				5, "-", 5,
			},
		},
		{
			"const z = 5 * 5;",
			"z",
			struct {
				left     int64
				operator string
				right    int64
			}{
				5, "*", 5,
			},
		},
		{
			"const a = 5 / 5;",
			"a",
			struct {
				left     int64
				operator string
				right    int64
			}{
				5, "/", 5,
			},
		},
		{
			"const b = 5 > 5;",
			"b",
			struct {
				left     int64
				operator string
				right    int64
			}{
				5, ">", 5,
			},
		},
		{
			"const c = 5 < 5;",
			"c",
			struct {
				left     int64
				operator string
				right    int64
			}{
				5, "<", 5,
			},
		},
		{
			"const d = 5 == 5;",
			"d",
			struct {
				left     int64
				operator string
				right    int64
			}{
				5, "==", 5,
			},
		},
		{
			"const e = 5 != 5;",
			"e",
			struct {
				left     int64
				operator string
				right    int64
			}{
				5, "!=", 5,
			},
		},
		{
			"const f = 5 >= 5;",
			"f",
			struct {
				left     int64
				operator string
				right    int64
			}{
				5, ">=", 5,
			},
		},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := program.Statements[0]

		if !testConstStatement(t, statement, val.expectedIdentifier) {
			return
		}

		constValue := statement.(*ast.ConstStatement).Value
		left := val.expectedValue.left
		operator := val.expectedValue.operator
		right := val.expectedValue.right

		if !testInfixExpression(t, constValue, left, operator, right) {
			return
		}

		identifier := statement.(*ast.ConstStatement).Name

		if !testLiteralExpression(t, identifier, val.expectedIdentifier) {
			return
		}
	}
}

func TestParseConstWithBooleanInfixExpression(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      struct {
			left     bool
			operator string
			right    bool
		}
	}{
		{
			"const x = true == true;",
			"x",
			struct {
				left     bool
				operator string
				right    bool
			}{
				true, "==", true,
			},
		},
		{
			"const y = true != false;",
			"y",
			struct {
				left     bool
				operator string
				right    bool
			}{
				true, "!=", false,
			},
		},
		{
			"const z = false == false;",
			"z",
			struct {
				left     bool
				operator string
				right    bool
			}{
				false, "==", false,
			},
		},
		{
			"const a = false != true;",
			"a",
			struct {
				left     bool
				operator string
				right    bool
			}{
				false, "!=", true,
			},
		},
		{
			"const b = true > false;",
			"b",
			struct {
				left     bool
				operator string
				right    bool
			}{
				true, ">", false,
			},
		},
		{
			"const c = true < false;",
			"c",
			struct {
				left     bool
				operator string
				right    bool
			}{
				true, "<", false,
			},
		},
		{
			"const d = true >= false;",
			"d",
			struct {
				left     bool
				operator string
				right    bool
			}{
				true, ">=", false,
			},
		},
		{
			"const e = true <= false;",
			"e",
			struct {
				left     bool
				operator string
				right    bool
			}{
				true, "<=", false,
			},
		},
		{
			"const f = false >= true;",
			"f",
			struct {
				left     bool
				operator string
				right    bool
			}{
				false, ">=", true,
			},
		},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := program.Statements[0]

		if !testConstStatement(t, statement, val.expectedIdentifier) {
			return
		}

		constValue := statement.(*ast.ConstStatement).Value
		left := val.expectedValue.left
		operator := val.expectedValue.operator
		right := val.expectedValue.right

		if !testInfixExpression(t, constValue, left, operator, right) {
			return
		}

		identifier := statement.(*ast.ConstStatement).Name

		if !testLiteralExpression(t, identifier, val.expectedIdentifier) {
			return
		}
	}
}

func TestParseConstWithFloatInfixExpression(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      struct {
			left     float64
			operator string
			right    float64
		}
	}{
		{
			"const x = 5.5 + 5.5;",
			"x",
			struct {
				left     float64
				operator string
				right    float64
			}{
				5.5, "+", 5.5,
			},
		},
		{
			"const y = 5.5 - 5.5;",
			"y",
			struct {
				left     float64
				operator string
				right    float64
			}{
				5.5, "-", 5.5,
			},
		},
		{
			"const z = 5.5 * 5.5;",
			"z",
			struct {
				left     float64
				operator string
				right    float64
			}{
				5.5, "*", 5.5,
			},
		},
		{
			"const a = 5.5 / 5.5;",
			"a",
			struct {
				left     float64
				operator string
				right    float64
			}{
				5.5, "/", 5.5,
			},
		},
		{
			"const b = 5.5 > 5.5;",
			"b",
			struct {
				left     float64
				operator string
				right    float64
			}{
				5.5, ">", 5.5,
			},
		},
		{
			"const c = 5.5 < 5.5;",
			"c",
			struct {
				left     float64
				operator string
				right    float64
			}{
				5.5, "<", 5.5,
			},
		},
		{
			"const d = 5.5 == 5.5;",
			"d",
			struct {
				left     float64
				operator string
				right    float64
			}{
				5.5, "==", 5.5,
			},
		},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := program.Statements[0]

		if !testConstStatement(t, statement, val.expectedIdentifier) {
			return
		}

		constValue := statement.(*ast.ConstStatement).Value
		left := val.expectedValue.left
		operator := val.expectedValue.operator
		right := val.expectedValue.right

		if !testInfixExpression(t, constValue, left, operator, right) {
			return
		}

		identifier := statement.(*ast.ConstStatement).Name

		if !testLiteralExpression(t, identifier, val.expectedIdentifier) {
			return
		}
	}
}

func TestParseConstWithDecimal(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      string
	}{
		{"const x = decimal(5.5);", "x", "decimal(5.5)"},
		{"const y = decimal(10.5);", "y", "decimal(10.5)"},
		{"const PI = decimal(3.14159);", "PI", "decimal(3.14159)"},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := program.Statements[0]

		if !testConstStatement(t, statement, val.expectedIdentifier) {
			return
		}

		constValue := statement.(*ast.ConstStatement).Value

		if constValue.String() != val.expectedValue {
			t.Errorf(
				"statement.Value not %s. got=%s",
				val.expectedValue,
				constValue.String(),
			)
			return
		}

		identifier := statement.(*ast.ConstStatement).Name

		if !testLiteralExpression(t, identifier, val.expectedIdentifier) {
			return
		}
	}
}

func testConstStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "const" {
		t.Errorf("s.TokenLiteral not 'const'. got=%q", s.TokenLiteral())
		return false
	}

	constStatement, ok := s.(*ast.ConstStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if constStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value not '%s'. got=%s", constStatement.Name.Value, name)
		return false
	}

	if constStatement.Name.TokenLiteral() != name {
		t.Errorf("letStatement.Name not '%s'. got=%s", constStatement.Name, name)
		return false
	}

	return true
}
