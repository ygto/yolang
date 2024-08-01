package main

import (
	"github.com/davecgh/go-spew/spew"
	"os"
	"path/filepath"
)

func main() {
	_ = spew.Dump
	path := os.Args[1]
	f, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	currPath = filepath.Dir(path)
	err = imports.importFile(f)
	if err != nil {
		panic(err)
	}
	spew.Dump(memory)
}
