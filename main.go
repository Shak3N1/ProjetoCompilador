//Author: Guilherme Nascimento
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Equanox/gotron"
	"github.com/shaken1/projetocompilador/lexer"
	"github.com/shaken1/projetocompilador/token"
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
	//UI
	window, err := gotron.New("ui_assets")
	if err != nil {
		panic(err)
	}
	window.On(&gotron.Event{Event: "close"}, func(bin []byte) {
		window.Close()
	})
	window.WindowOptions.Width = 1280
	window.WindowOptions.Height = 720
	window.WindowOptions.Title = "Rufus Editor"
	window.WindowOptions.Frame = false
	done, err := window.Start()
	if err != nil {
		panic(err)
	}
	<-done
}

func getFileAsString(filePath string) (string, error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(fileData), nil
}
