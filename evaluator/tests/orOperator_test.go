package tests

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestOrLogicalWithBoolean(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true or true", true},
		{"true or false", true},
		{"false or true", true},
		{"false or false", false},
		{"1 or true", true},
		{"1 or false", true},
		{"0 or true", true},
		{"0 or false", false},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		testBooleanObject(t, evaluated, val.expected)
	}
}

func TestOrLogicalWithInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"1 or 1", 1},
		{"1 or 0", 1},
		{"0 or 1", 1},
		{"0 or 0", 0},
		{"1 or 2", 1},
		{"2 or 1", 2},
		{"2 or 2", 2},
		{"2 or 0", 2},
		{"0 or 2", 2},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testIntegerObject(t, evaluated, val.expected)
	}
}

func TestOrLogicalWithFloat(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"1.1 or 1.1", 1.1},
		{"1.1 or 0.0", 1.1},
		{"0.0 or 1.1", 1.1},
		{"0.0 or 0.0", 0.0},
		{"1.1 or 2.2", 1.1},
		{"2.2 or 1.1", 2.2},
		{"2.2 or 2.2", 2.2},
		{"2.2 or 0.0", 2.2},
		{"0.0 or 2.2", 2.2},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testFloatObject(t, evaluated, val.expected)
	}
}

func TestOrLogicalWithDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected decimal.Decimal
	}{
		{"decimal(1) or decimal(1.1)", decimal.NewFromFloat(1)},
		{"decimal(1) or decimal(0.0)", decimal.NewFromFloat(1)},
		{"decimal(0.0) or decimal(1.1)", decimal.NewFromFloat(1.1)},
		{"decimal(0.0) or decimal(0.0)", decimal.NewFromFloat(0)},
		{"decimal(1.1) or decimal(2.2)", decimal.NewFromFloat(1.1)},
		{"decimal(2.2) or decimal(1.1)", decimal.NewFromFloat(2.2)},
		{"decimal(2.2) or decimal(2.2)", decimal.NewFromFloat(2.2)},
		{"decimal(2.2) or decimal(0.0)", decimal.NewFromFloat(2.2)},
		{"decimal(0.0) or decimal(2.2)", decimal.NewFromFloat(2.2)},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testDecimalObject(t, evaluated, val.expected)
	}
}
