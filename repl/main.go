package repl

import (
	"fmt"
	"log"

	"github.com/Mostafa-DE/delang/evaluator"
	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/object"
	"github.com/Mostafa-DE/delang/parser"
	"github.com/eiannone/keyboard"
)

func StartSession() {
	PROMPT := ">>> "

	env := object.NewEnvironment()

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	fmt.Printf("Hi! Welcome to DE v0.0.9\n")
	fmt.Printf("Type '.help' to see a list of commands.\n")

	history := []string{}
	historyIndex := 0
	currentInput := ""
	cursorPosition := 0

	fmt.Print("\n")
	fmt.Print(PROMPT)

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyCtrlC {
			fmt.Println("\nBye!")
			break

		} else if key == keyboard.KeyEnter {
			fmt.Println()
			if currentInput != "" {
				history = append(history, currentInput)
				historyIndex = len(history)

				startExec(currentInput, env)

				currentInput = ""
				cursorPosition = 0
				fmt.Print(PROMPT)
			}

		}

		if key == keyboard.KeyArrowUp {
			if historyIndex > 0 {
				historyIndex--
				currentInput = history[historyIndex]
				cursorPosition = len(currentInput)
				clearCurrentLine()
				fmt.Print(PROMPT)
				fmt.Print(currentInput)
			}

		}

		if key == keyboard.KeyArrowDown {
			if historyIndex < len(history)-1 {
				historyIndex++
				currentInput = history[historyIndex]
				cursorPosition = len(currentInput)
				clearCurrentLine()
				fmt.Print(PROMPT)
				fmt.Print(currentInput)

			} else if historyIndex == len(history)-1 {
				historyIndex++
				currentInput = ""
				cursorPosition = 0
				clearCurrentLine()
				fmt.Print(PROMPT)
			}
		}

		if key == keyboard.KeyBackspace || key == keyboard.KeyBackspace2 {
			if cursorPosition > 0 {
				currentInput = currentInput[:cursorPosition-1] + currentInput[cursorPosition:]
				cursorPosition--
				clearCurrentLine()
				fmt.Print(PROMPT)
				fmt.Print(currentInput)
				moveCursorLeft(len(currentInput) - cursorPosition - 1)
			}
		}

		if key == keyboard.KeyArrowLeft {
			if cursorPosition > 0 {
				cursorPosition--
				moveCursorLeft(1)
			}
		}

		if key == keyboard.KeyArrowRight {
			if cursorPosition < len(currentInput) {
				cursorPosition++
				moveCursorRight(1)
			}
		}

		if key == keyboard.KeyCtrlL {
			fmt.Print("\033[H\033[2J")
			fmt.Print(PROMPT)
		}

		if char != 0 {
			currentInput = currentInput[:cursorPosition] + string(char) + currentInput[cursorPosition:]
			cursorPosition++
			clearCurrentLine()
			fmt.Print(PROMPT)
			fmt.Print(currentInput)
			moveCursorLeft(len(currentInput) - cursorPosition - 1)
		}

		if key == keyboard.KeySpace {
			currentInput = currentInput[:cursorPosition] + " " + currentInput[cursorPosition:]
			cursorPosition++
			clearCurrentLine()
			fmt.Print(PROMPT)
			fmt.Print(currentInput)

			// TODO: Fix the cursor position after adding a space
			moveCursorLeft(len(currentInput) - cursorPosition - 1)
		}
	}
}

func startExec(command string, env *object.Environment) {
	l := lexer.New(command)
	p := parser.New(l)

	if command == ".clear" {
		fmt.Print("\033[H\033[2J")
		return
	}

	if command == ".help" {
		// This should be a separate function in the future
		fmt.Println("Commands:")
		fmt.Println(" ctrl + c: Exit the REPL")
		fmt.Println(" .clear: Clear the terminal screen")
		return
	}

	program := p.ParseProgram()

	if parserErrors(p) {
		return
	}

	eval := evaluator.Eval(program, env)

	if eval != nil {
		if eval.Type() == object.STRING_OBJ {
			fmt.Printf("'%s'\n", eval.Inspect())
		} else {
			fmt.Println(eval.Inspect())
		}
	} else {
		fmt.Println("null")
	}

}
