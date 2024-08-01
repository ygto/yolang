package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"path/filepath"
	"strings"
	"yolang/parser"
)

type Import struct {
}

type Imports map[string]*Import

var imports Imports = make(Imports)
var currPath string

func (i Imports) importFile(path string) error {
	if !strings.HasSuffix(path, ".yo") {
		path = filepath.Join(currPath, path+".yo")
		currPath = filepath.Dir(path)
	}
	//pathInfo := filepath.Base(path)
	fmt.Printf("Importing %s\n", path)
	f := i[path]
	if f != nil {
		return nil
	}
	input, err := antlr.NewFileStream(path)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return err
	}

	lexer := parser.NewYoLangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewYoLangParser(stream)
	listener := NewYoLangListener()

	tree := p.Prog()

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	f = &Import{}
	i[path] = f

	return nil
}
