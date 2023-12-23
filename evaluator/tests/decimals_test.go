package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/object"
	"github.com/shopspring/decimal"
)

func TestDecimalOperationsExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected decimal.Decimal
	}{
		{"decimal(55)", decimal.NewFromFloat(55)},
		{"decimal(100)", decimal.NewFromFloat(100)},
		{"decimal(1.111111)", decimal.NewFromFloat(1.111111)},
		{"decimal(-50)", decimal.NewFromFloat(-50)},
		{"decimal(-100)", decimal.NewFromFloat(-100)},
		{"decimal(5) + decimal(5) + decimal(5) + decimal(5) - decimal(10)", decimal.NewFromFloat(10)},
		{"decimal(2) * decimal(2) * decimal(2) * decimal(2) * decimal(2)", decimal.NewFromFloat(32)},
		{"decimal(-50) + decimal(100) + decimal(-50)", decimal.NewFromFloat(0)},
		{"decimal(5) * decimal(2) + decimal(10)", decimal.NewFromFloat(20)},
		{"decimal(5) + decimal(2) * decimal(10)", decimal.NewFromFloat(25)},
		{"decimal(20) + decimal(2) * decimal(-10)", decimal.NewFromFloat(0)},
		{"decimal(50) / decimal(2) * decimal(2) + decimal(10)", decimal.NewFromFloat(60)},
		{"decimal(2) * (decimal(5) + decimal(10))", decimal.NewFromFloat(30)},
		{"decimal(3) * decimal(3) * decimal(3) + decimal(10)", decimal.NewFromFloat(37)},
		{"decimal(3) * (decimal(3) * decimal(3)) + decimal(10)", decimal.NewFromFloat(37)},
		{"(decimal(5) + decimal(10) * decimal(2) + decimal(15) / decimal(3)) * decimal(2) + decimal(-10)", decimal.NewFromFloat(50)},
		{"decimal(5) % decimal(2)", decimal.NewFromFloat(1)},
		{"decimal(5) % decimal(3)", decimal.NewFromFloat(2)},
		{"decimal(2) % decimal(2)", decimal.NewFromFloat(0)},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testDecimalObject(t, evaluated, val.expected)
	}
}

func TestDivRoundingDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`
				decimal(300) / decimal(1.2121)
			`,
			"247.50433133",
		},
		{
			`
				_getDecimalData["divPrec"] = 10;	
				
				decimal(300) / decimal(1.2121)
			`,
			"247.5043313258",
		},
		{
			`
				_getDecimalData["divPrec"] = 1;

				decimal(300) / decimal(1.2121)
			`,
			"247.5",
		},
		{
			`
				_getDecimalData["divPrec"] = 20;

				decimal(300) / decimal(1.2121)

			`,
			"247.50433132579820146853",
		},
		{
			`
				_getDecimalData["divPrec"] = -1;

				decimal(300) / decimal(1.2121)
			`,
			"Valid range for divPrec is [0 to 28]",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		if evaluated.Type() == object.ERROR_OBJ {
			evaluatedError, ok := evaluated.(*object.Error)

			if !ok {
				t.Errorf("Expected evaluated to be an error. Got=%T", evaluated)
			}

			if evaluatedError.Msg != val.expected {
				t.Errorf("Expected %s, got %s", val.expected, evaluatedError.Msg)
			}

			continue
		}

		evaluatedDecimal, ok := evaluated.(*object.Decimal)

		if !ok {
			t.Errorf("Expected evaluated to be a decimal. Got=%T", evaluated)
		}

		if evaluatedDecimal.Value.String() != val.expected {
			t.Errorf("Expected %s, got %s", val.expected, evaluatedDecimal.Value.String())
		}

	}
}

func TestMulRoundingDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`
				decimal(1.21113) * decimal(2.22113)
			`,
			"2.69007718",
		},
		{
			`
				_getDecimalData["prec"] = 5;	
				
				decimal(1.21113) * decimal(2.22113)
			`,
			"2.69008",
		},
		{
			`
				_getDecimalData["prec"] = 3;

				decimal(1.21113) * decimal(2.22113)
			`,
			"2.69",
		},
		{
			`
				_getDecimalData["prec"] = 1;

				decimal(1.21113) * decimal(2.22113)

			`,
			"2.7",
		},
		{
			`
				_getDecimalData["prec"] = 5;

				decimal(300) * decimal(1.2121);
			`,
			"363.63",
		},
		{
			`
				_getDecimalData["prec"] = -1;

				decimal(300) * decimal(1.2121)
			`,
			"Valid range for prec is [0 to 8]",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		if evaluated.Type() == object.ERROR_OBJ {
			evaluatedError, ok := evaluated.(*object.Error)

			if !ok {
				t.Errorf("Expected evaluated to be an error. Got=%T", evaluated)
			}

			if evaluatedError.Msg != val.expected {
				t.Errorf("Expected %s, got %s", val.expected, evaluatedError.Msg)
			}

			continue
		}

		evaluatedDecimal, ok := evaluated.(*object.Decimal)

		if !ok {
			t.Errorf("Expected evaluated to be a decimal. Got=%T", evaluated)
		}

		if evaluatedDecimal.Value.String() != val.expected {
			t.Errorf("Expected %s, got %s", val.expected, evaluatedDecimal.Value.String())
		}

	}
}

func TestAddRoundingDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`
				decimal(1.21113) + decimal(2.22113);
			`,
			"3.43226",
		},
		{
			`
				_getDecimalData["prec"] = 5;	
				
				decimal(1.21113) + decimal(2.22113);
			`,
			"3.43226",
		},
		{
			`
				_getDecimalData["prec"] = 3;

				decimal(1.21113) + decimal(2.22113);
			`,
			"3.432",
		},
		{
			`
				_getDecimalData["prec"] = 1;

				decimal(1.21113) + decimal(2.22113);

			`,
			"3.4",
		},
		{
			`
				_getDecimalData["prec"] = 5;

				decimal(300) + decimal(1.2121);
			`,
			"301.2121",
		},
		{
			`
				_getDecimalData["prec"] = -1;

				decimal(300) + decimal(1.2121)
			`,
			"Valid range for prec is [0 to 8]",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		if evaluated.Type() == object.ERROR_OBJ {
			evaluatedError, ok := evaluated.(*object.Error)

			if !ok {
				t.Errorf("Expected evaluated to be an error. Got=%T", evaluated)
			}

			if evaluatedError.Msg != val.expected {
				t.Errorf("Expected %s, got %s", val.expected, evaluatedError.Msg)
			}

			continue
		}

		evaluatedDecimal, ok := evaluated.(*object.Decimal)

		if !ok {
			t.Errorf("Expected evaluated to be a decimal. Got=%T", evaluated)
		}

		if evaluatedDecimal == nil {
			t.Errorf("Expected evaluated to be a decimal. Got=nil")
		}

		if evaluatedDecimal != nil && evaluatedDecimal.Value.String() != val.expected {
			t.Errorf("Expected %s, got %s", val.expected, evaluatedDecimal.Value.String())
		}

	}

}

func TestSubRoundingDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`
				decimal(1.2111) - decimal(2.2211);
			`,
			"-1.01",
		},
		{
			`
				_getDecimalData["prec"] = 3;	
				
				decimal(1.2911) - decimal(22.221);
			`,
			"-20.93",
		},
		{
			`
				_getDecimalData["prec"] = 3;

				decimal(32.7564) - decimal(34.5478);
			`,
			"-1.791",
		},
		{
			`
				_getDecimalData["prec"] = 1;

				decimal(1.21113) - decimal(2.22113);

			`,
			"-1",
		},
		{
			`
				_getDecimalData["prec"] = 5;

				decimal(300) - decimal(1.2121);
			`,
			"298.7879",
		},
		{
			`
				_getDecimalData["prec"] = -1;

				decimal(300) - decimal(1.2121)
			`,
			"Valid range for prec is [0 to 8]",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		if evaluated.Type() == object.ERROR_OBJ {
			evaluatedError, ok := evaluated.(*object.Error)

			if !ok {
				t.Errorf("Expected evaluated to be an error. Got=%T", evaluated)
			}

			if evaluatedError.Msg != val.expected {
				t.Errorf("Expected %s, got %s", val.expected, evaluatedError.Msg)
			}

			continue
		}

		evaluatedDecimal, ok := evaluated.(*object.Decimal)

		if !ok {
			t.Errorf("Expected evaluated to be a decimal. Got=%T", evaluated)
		}

		if evaluatedDecimal.Value.String() != val.expected {
			t.Errorf("Expected %s, got %s", val.expected, evaluatedDecimal.Value.String())
		}

	}
}
