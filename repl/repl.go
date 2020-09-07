package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/saulneri1998/saulanguage/lexer"
	"github.com/saulneri1998/saulanguage/token"
)

// PROMPT is the string to be printed in every REPL loop
const PROMPT = ">>"

// Start the REPL
func Start(in io.Reader, out io.Writer) {
	var scanner *bufio.Scanner = bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		var scanned bool = scanner.Scan()
		if !scanned {
			return
		}

		var line string = scanner.Text()
		var lex *lexer.Lexer = lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
