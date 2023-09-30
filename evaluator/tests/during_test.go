package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/evaluator"
	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/object"
	"github.com/Mostafa-DE/delang/parser"
)

func TestDuringEval(t *testing.T) {
	tests := []struct {
		Desc     string
		input    string
		envVal   interface{}
		expected interface{}
	}{
		{
			"It should work fine with simple (during loop) with break statement",
			`
				let x = 0;
				during x < 4: {
					x = x + 1;

					logs(x);

					if x == 3: {
						break;
					}
				}
			`,
			"x",
			3,
		},
		{
			"It should work fine with IIFE",
			`
				let x = 0;
				during x < 4: {
					let x = fun() {
						return 1;
					}();

					if x == 1: {
						break;
					}
				}
			`,
			"x",
			1,
		},
		{
			"It should work fine with if/else statement",
			`
				let x = 0;
				during true: {
					if x + 10 == 20: {
						break;
					} else {
						x = x + 1;
						logs("ERROR");
					}
				}
			`,
			"x",
			10,
		},
		{
			"It shouldn't run with false condition",
			`
				let x = 0;
				during false: {
					logs("ERROR");
				}
			`,
			"x",
			0,
		},
		{
			"It should skip the rest of the loop body with skip statement",
			`
				let x = 1;
				during x < 3: {
					x = x + 1;
					skip;

					x = 100
				}
			`,
			"x",
			3,
		},
		{
			"It should work fine with skip & break statements",
			`
				let x = 0;
				during x < 10: {
					if x % 2 == 0: {
						x = x + 2;
						skip;
					}

					x = x + 1;

					logs(x);

					if x == 9: {
						break;
					}
				}
			`,
			"x",
			10,
		},
		{
			"It should work fine with nested during loops",
			`
				let x = 0;
				during x < 20: {
					x = x + 1;

					during x < 5: {
						x = x + 1;
						
						if x == 5: {
							break;
						}
					}

					if x == 15: {
						break;
					}
				}
			`,
			"x",
			15,
		},
	}

	for _, val := range tests {
		l := lexer.New(val.input)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) > 0 {
			throwError(p.Errors()[0])
			return
		}

		env := object.NewEnvironment()

		evaluator.Eval(program, env)

		envVal := val.envVal.(string)

		retVal, _ := env.Get(envVal)

		_val, ok := retVal.(*object.Integer)

		if !ok {
			t.Errorf("_val is not Integer, got %T", retVal)
			return
		}

		switch val.expected.(type) {
		case int:
			if _val.Value != int64(val.expected.(int)) {
				t.Errorf("Expected %v to be %d, got %d", envVal, val.expected, _val)
			}

		case int64:
			if _val.Value != val.expected.(int64) {
				t.Errorf("Expected %v to be %d, got %d", envVal, val.expected, _val)
			}

		}

	}

}
