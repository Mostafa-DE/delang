package repl

import (
	"fmt"
	"log"
	"os"

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

	fmt.Printf("Hi! Welcome to DE\n")
	fmt.Printf("Type '.help' to see a list of commands.\n")

	history := []string{}

	history = loadHistoryFromFile("history.txt")

	historyIndex := 0
	currentInput := ""
	cursorPosition := 0
	quitCount := 0

	fmt.Print("\n")
	fmt.Print(PROMPT)

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch key {
		case keyboard.KeyCtrlC:
			quitCount++
			if quitCount == 2 {
				fmt.Println("\nBye!")
				return
			} else {
				if err := saveHistoryToFile(history, "history.txt"); err != nil {
					log.Fatalf("Failed to save history: %v", err)
				}

				fmt.Println()
				fmt.Printf("\033[33m%s\033[0m\n", "Press ctrl + c again to exit.")
			}

		case keyboard.KeyEnter:
			handleEnterKey(&history, &historyIndex, &currentInput, &cursorPosition, env)

		case keyboard.KeyArrowUp, keyboard.KeyArrowDown:
			handleArrowUpDown(key, &history, &historyIndex, &currentInput, &cursorPosition)

		case keyboard.KeyArrowLeft, keyboard.KeyArrowRight:
			handleArrowLeftRight(key, &currentInput, &cursorPosition)

		case keyboard.KeyBackspace, keyboard.KeyBackspace2:
			handleBackspace(&currentInput, &cursorPosition)

		case keyboard.KeySpace:
			insertCharacterAtCursor(&currentInput, &cursorPosition, " ")

		case keyboard.KeyCtrlL:
			clearScreen()

		default:
			if char != 0 {
				insertCharacterAtCursor(&currentInput, &cursorPosition, string(char))
			}
		}

		refreshLine(PROMPT, currentInput, cursorPosition)
	}
}

func startExec(command string, env *object.Environment, history *[]string, historyIndex *int) {
	l := lexer.New(command)
	p := parser.New(l)

	if command == ".clear" {
		fmt.Print("\033[H\033[2J")
		return
	}

	if command == ".exit" {
		fmt.Println("Bye!")
		os.Exit(0)
	}

	if command == ".clearHistory" {
		*history = []string{}
		*historyIndex = 0
		fmt.Println("History cleared.")
		return
	}

	if command == ".help" {
		// This should be a separate function in the future
		fmt.Println("Commands:")
		fmt.Println(" ctrl + c or .exit: Exit the REPL")
		fmt.Println(" .clear: Clear the terminal screen")
		fmt.Println(" .clearHistory: Clear the command history")
		fmt.Println(" .help: Show this help message")
		fmt.Println()

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

func handleEnterKey(history *[]string, historyIndex *int, currentInput *string, cursorPosition *int, env *object.Environment) {
	fmt.Println()
	if *currentInput != "" {
		*history = append(*history, *currentInput)
		*historyIndex = len(*history)

		startExec(*currentInput, env, history, historyIndex)

		*currentInput = ""
		*cursorPosition = 0
	}
}

func handleArrowUpDown(key keyboard.Key, history *[]string, historyIndex *int, currentInput *string, cursorPosition *int) {
	if key == keyboard.KeyArrowUp && *historyIndex > 0 { // we check if historyIndex > 0 to prevent index out of range error
		*historyIndex--
	} else if key == keyboard.KeyArrowDown && *historyIndex < len(*history)-1 {
		*historyIndex++
	}

	if *historyIndex == len(*history) {
		*currentInput = ""
		*cursorPosition = 0

	} else if key == keyboard.KeyArrowUp || (key == keyboard.KeyArrowDown && *historyIndex < len(*history)) {
		*currentInput = (*history)[*historyIndex]
		*cursorPosition = len(*currentInput)

	} else if key == keyboard.KeyArrowDown && *historyIndex == len(*history) {
		// This is to handle the case when the user presses the down arrow key after reaching the end of the history
		*currentInput = ""
		*cursorPosition = 0
	}
}

func handleArrowLeftRight(key keyboard.Key, currentInput *string, cursorPosition *int) {
	if key == keyboard.KeyArrowLeft && *cursorPosition > 0 {
		*cursorPosition--
	} else if key == keyboard.KeyArrowRight && *cursorPosition < len(*currentInput) {
		*cursorPosition++
	}
}

func handleBackspace(currentInput *string, cursorPosition *int) {
	if *cursorPosition > 0 {
		/*
			** We use `cursorPosition-1` because we want to delete the character before the cursor
			and then move the cursor back one position.

			==> e.g. if the cursor is at position 5, we want to delete the character at position 4
				and then move the cursor back to position 4.

			** Not only this but also we want to keep anything after the cursor position.

			==> e.g. if the currentInput is "hello world" and the cursor is at position 5,
				we want to delete the character at position 4 and then move the cursor back to position 4
				and keep the rest of the string after the cursor position. so the result will be "hell world".
		*/

		*currentInput = (*currentInput)[:*cursorPosition-1] + (*currentInput)[*cursorPosition:]
		*cursorPosition--
	}
}

func insertCharacterAtCursor(currentInput *string, cursorPosition *int, char string) {
	// This function inserts a character at the cursor position and then moves the cursor forward one position.
	*currentInput = (*currentInput)[:*cursorPosition] + char + (*currentInput)[*cursorPosition:]
	*cursorPosition++
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func refreshLine(prompt string, currentInput string, cursorPosition int) {
	clearCurrentLine()
	fmt.Print(prompt)
	fmt.Print(currentInput)

	// We only need to move the cursor if the cursor is not at the end of the line
	if len(currentInput) > cursorPosition {
		moveCursorLeft(len(currentInput) - cursorPosition)
	}
}
