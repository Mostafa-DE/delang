package tests

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestAndLogicalWithBoolean(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true and true", true},
		{"true and false", false},
		{"false and true", false},
		{"false and false", false},
		{"1 and true", true},
		{"1 and false", false},
		{"0 and true", false},
		{"0 and false", false},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testBooleanObject(t, evaluated, val.expected)

	}

}

func TestAndLogicalWithInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"1 and 1", 1},
		{"1 and 0", 0},
		{"0 and 1", 0},
		{"0 and 0", 0},
		{"1 and 2", 2},
		{"2 and 1", 1},
		{"2 and 2", 2},
		{"2 and 0", 0},
		{"0 and 2", 0},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testIntegerObject(t, evaluated, val.expected)
	}
}

func TestAndLogicalWithFloat(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"1.1 and 1.1", 1.1},
		{"1.1 and 0.0", 0.0},
		{"0.0 and 1.1", 0.0},
		{"0.0 and 0.0", 0.0},
		{"1.1 and 2.2", 2.2},
		{"2.2 and 1.1", 1.1},
		{"2.2 and 2.2", 2.2},
		{"2.2 and 0.0", 0.0},
		{"0.0 and 2.2", 0.0},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testFloatObject(t, evaluated, val.expected)
	}
}

func TestAndLogicalWithDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected decimal.Decimal
	}{
		{"decimal(1) and decimal(1.1)", decimal.NewFromFloat(1.1)},
		{"decimal(1) and decimal(0.0)", decimal.NewFromFloat(0.0)},
		{"decimal(0.0) and decimal(1.1)", decimal.NewFromFloat(0.0)},
		{"decimal(0.0) and decimal(0.0)", decimal.NewFromFloat(0.0)},
		{"decimal(1.1) and decimal(2.2)", decimal.NewFromFloat(2.2)},
		{"decimal(2.2) and decimal(1.1)", decimal.NewFromFloat(1.1)},
		{"decimal(2.2) and decimal(2.2)", decimal.NewFromFloat(2.2)},
		{"decimal(2.2) and decimal(0.0)", decimal.NewFromFloat(0.0)},
		{"decimal(0.0) and decimal(2.2)", decimal.NewFromFloat(0.0)},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testDecimalObject(t, evaluated, val.expected)
	}
}
