//Author: Guilherme Nascimento
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"shaken1/projetocompilador/lexer"
	"shaken1/projetocompilador/token"
)

func main() {
	//Lexical Tests
	file := os.Args[1]
	fileStr, _ := getFileAsString(file)
	Lexer := lexer.NewLexer([]byte(fileStr))
	for {
		if tok := Lexer.Scan(); tok != nil {
			if tok.Type == token.TokMap.Type("INVALID") {
				fmt.Printf("Lexical error at line %d, column %d\n",
					tok.Pos.Line, tok.Pos.Column)
				return
			} else if tok.Type == token.EOF {
				break
			}
		} else {
			break
		}
	}
	fmt.Println("Lexical Sucess!")
}

func getFileAsString(filePath string) (string, error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(fileData), nil
}
