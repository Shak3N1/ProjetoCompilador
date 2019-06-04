package main

import (
	"github.com/shaken1/ProjetoCompilador/lexer"
	"github.com/shaken1/ProjetoCompilador/parser"
	"github.com/shaken1/ProjetoCompilador/token"
)

type Token struct {
	token  string
	lexeme string
	line   uint
}

func synAnalyze(data string) error {
	Lexer := lexer.NewLexer([]byte(data))
	p := parser.NewParser()
	_, err := p.Parse(Lexer)
	return err
}
func lexAnalyze(data string) []Token {
	var tokens []Token
	Lexer := lexer.NewLexer([]byte(data))
	for {
		if tok := Lexer.Scan(); tok != nil {
			if tok.Type == token.TokMap.Type("INVALID") {
				_token := Token{"ERROR", string(tok.Lit), uint(tok.Line)}
				tokens = append(tokens, _token)
				break
			} else if tok.Type == token.EOF {
				break
			}
			_token := Token{token.TokMap.Id(tok.Type), string(tok.Lit), uint(tok.Line)}
			tokens = append(tokens, _token)
		} else {
			break
		}
	}
	return tokens
}
