package main

import (
	"bytes"
	"fmt"
)

type Variable struct {
	ptr  int
	size int
	t    string
}

var stack []int = make([]int, 0)
var stackPtr int = MEMORY_SIZE

func ptr() int {
	return stackPtr
}
func decrPtr(size int) {
	stackPtr -= size
}

func stackPush() {
	stack = append(stack, ptr())
}
func stackPop() {
	stack = stack[:len(stack)-1]
	stackPtr = stack[len(stack)-1]
}

func NewVariable() *Variable {
	return &Variable{}
}

type Variables map[string]*Variable

var variables = make(Variables)

func (t Variables) Add(name string, val *Variable) {
	val.ptr = ptr() - val.size
	t[name] = val

	fmt.Printf("Added variable: %s, size: %d address:%d\n", name, val.size, val.ptr)
	//fmt.Printf("mov variable: %s, size: %d\n", name, val.size)
	decrPtr(val.size)
	api.Variable(name, val)
}

func (t Variables) Get(name string) *Variable {
	val, has := t[name]
	if !has {
		fatalf("variable not found: %s\n", name)
	}
	return val
}

func (t Variables) Set(name string, val Value) {
	v := t.Get(name)
	b := val.Bytes()
	var size int
	buff := bytes.Buffer{}
	switch v.t {
	case "int":
		size = INT_SIZE
	case "varchar":
		if len(b) > v.size {
			fatalf("variable %s size too big: %d\n", name, len(b))
		}
		size = v.size
	}
	buff.Write(b)

	fmt.Printf("Setting variable: %s, size: %d and byte size %d\n", name, size, buff.Len())
	bs := buff.Bytes()
	for i := 0; i < len(bs); i++ {
		p := v.ptr + i
		//fmt.Printf("0x%d \t %s: % X\n", p, name, bs[i])
		memory[p] = bs[i]
	}
}
func (t Variables) SetByPtr(ptr int, val Value) {
	b := val.Bytes()
	buff := bytes.Buffer{}
	buff.Write(b)

	bs := buff.Bytes()
	for i := 0; i < len(b); i++ {
		p := ptr + i
		fmt.Printf("0x%d \t  % X\n", p, bs[i])
		memory[p] = bs[i]
	}
}
