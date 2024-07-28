package main

import (
	"fmt"
	"os"
	"yolang/parser"

	"github.com/antlr4-go/antlr/v4"
)

func main() {

	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}

	lexer := parser.NewLangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewLangParser(stream)

	listener := NewLangListener()
	tree := p.Prog()

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}
