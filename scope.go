package main

type Scope interface {
	GetRoot() Scope
	GetValue(k string) *Value
	SetValue(k string, v interface{})
	GetPtr() int
	IncrPtr(v int)
	CreateVariable(name string, varType string)
}
