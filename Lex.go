package main

import (
	"fmt"
	"trollzinho_da_night/lexer"
	"trollzinho_da_night/token"
)

type Token struct {
	lexema   string
	atributo string
	linha    uint
}

func lexAnalyze(data string) []Token {
	var tokens []Token
	Lexer := lexer.NewLexer([]byte(data))
	for {
		if tok := Lexer.Scan(); tok != nil {
			if tok.Type == token.TokMap.Type("INVALID") {
				_token := Token{"ERROR", string(tok.Lit), uint(tok.Pos.Line)}
				tokens = append(tokens, _token)
				break
			} else if tok.Type == token.EOF {
				break
			}
			_token := Token{token.TokMap.Id(tok.Type), string(tok.Lit), uint(tok.Pos.Line)}
			tokens = append(tokens, _token)
		} else {
			break
		}
	}
	for _, token := range tokens {
		fmt.Println(token)
	}
	return tokens
}
