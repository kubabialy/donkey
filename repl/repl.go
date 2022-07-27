package repl

import (
	"bufio"
	"fmt"
	"github.com/kubabialy/donkey/lexer"
	"github.com/kubabialy/donkey/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lxr := lexer.New(line)

		for tok := lxr.NextToken(); tok.Type != token.EOF; tok = lxr.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
