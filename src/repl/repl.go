package repl

import (
	"bufio"
	"fmt"
	"io"
	"lexer"
	"token"
)

const PROMPT = ">> "


func StartSession(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {return}

		lineText := scanner.Text()
		l := lexer.New(lineText)

		for tok := l.NextToken(); tok.Type != token.EOFILE ; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}


	}
}
