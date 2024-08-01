package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

const BYTE_SIZE = 1
const INT_SIZE = BYTE_SIZE * 8
const VARCHAR_SIZE = BYTE_SIZE * 128

func dump(v interface{}) {
	fmt.Println(strings.Repeat("-", 10))
	fmt.Printf("Type : %s : %#v \n", structName(v), v)
	fmt.Println(strings.Repeat("-", 10))
}

func structName(v interface{}) string {
	return fmt.Sprintf("%v", reflect.TypeOf(v))
}

func dumpTree(t []antlr.Tree) {
	for _, t := range t {
		dump(t.GetPayload())
	}
}

func fatalf(format string, a ...any) {
	panic(fmt.Errorf(format, a...))
}
