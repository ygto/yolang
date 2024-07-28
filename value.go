package main

type Value struct {
	Ptr  int
	Type string
	Val  interface{}
}

func NewValue(ptr int, t string, v interface{}) *Value {
	return &Value{
		Ptr:  ptr,
		Type: t,
		Val:  v,
	}
}
