package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Mostafa-DE/delang/parser"
)

func parserErrors(p *parser.Parser) bool {
	if len(p.Errors()) != 0 {
		fmt.Println("Error parsing program:")
		fmt.Println(p.Errors()[0])

		return true
	}

	return false
}

func clearCurrentLine() {
	fmt.Print("\r")
	fmt.Print(strings.Repeat(" ", 80))
	fmt.Print("\r")
}

func moveCursorLeft(n int) {
	fmt.Printf("\033[%dD", n)
}

func saveHistoryToFile(history []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, cmd := range history {
		if _, err := file.WriteString(cmd + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func loadHistoryFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}

	defer file.Close()

	var history []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		history = append(history, scanner.Text())
	}

	return history
}
