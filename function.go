package main

import (
	"fmt"
	"yolang/parser"
)

type Function struct {
	root       Scope
	name       string
	ptr        int
	returnType string
	paramType  string
	param      string
	values     map[string]*Value
	clousure   *parser.FuncExprContext
}

func NewFunction(parent Scope, name string, ptr int, returnType, paramType, param string, c *parser.FuncExprContext) *Function {
	f := &Function{
		root:       toRoot(parent),
		name:       name,
		ptr:        ptr,
		returnType: returnType,
		paramType:  paramType,
		param:      param,
		values:     make(map[string]*Value),
		clousure:   c,
	}
	return f
}

func toRoot(scope Scope) Scope {
	parent := scope.GetRoot()
	if parent == nil {
		return scope
	}
	return toRoot(parent)
}

func (l *Function) GetValue(k string) *Value {
	val, has := l.values[k]
	if !has {
		fmt.Printf("variable not found: %s\n", k)
		return nil
	}
	return val
}

func (l *Function) SetValue(k string, v interface{}) {
	t := l.GetValue(k)

	fmt.Printf("0x%d [%s] : %s = %v(%s)\n", t.Ptr, l.name, k, v, t.Type)
	l.values[k] = NewValue(t.Ptr, t.Type, v)
}
func (l *Function) CreateVariable(k string, t string) {
	l.values[k] = NewValue(l.ptr, t, nil)
	size := getTypeSize(t)
	fmt.Printf("0x%d : var[%s] %s = %s[%d] \n", l.GetPtr(), l.name, k, t, size)
	l.IncrPtr(size)
}

func (l *Function) GetRoot() Scope {
	return l.root
}

func (l *Function) GetPtr() int {
	return l.root.GetPtr()
}

func (l *Function) IncrPtr(v int) {
	l.root.IncrPtr(v)
}
