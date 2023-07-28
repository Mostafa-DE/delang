package repl

import (
	"io"
	"os/exec"
)

func printParserErrors(output io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(output, "\t"+msg+"\n")
	}
}

func clearConsole(output io.Writer) {
	cmd := exec.Command("clear")
	cmd.Stdout = output
	cmd.Run()
}
