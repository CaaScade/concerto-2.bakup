package manager

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/koki/concerto/pkg/parser"
)

func Parse(files []string) ([]antlr.ParserRuleContext, error) {
	parseTrees := []antlr.ParserRuleContext{}

	for i := range files {
		input, err := antlr.NewFileStream(files[i])
		if err != nil {
			return nil, err
		}

		lexer := parser.NewConcertoLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)

		p := parser.NewConcertoParser(stream)
		p.BuildParseTrees = true
		parseTrees = append(parseTrees, p.Prog())
	}
	return parseTrees, nil
}
