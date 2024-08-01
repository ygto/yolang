package main

import (
	"bytes"
	"encoding/binary"
)

type Value interface {
	Bytes() []byte
}

type IntValue struct {
	value int
}

func NewIntValue(value int) *IntValue {
	return &IntValue{value}
}

func (v *IntValue) Bytes() []byte {
	bs := make([]byte, INT_SIZE)
	binary.BigEndian.PutUint64(bs, uint64(v.value))
	return bs
}

type VarcharValue struct {
	value string
}

func NewVarcharValue(value string) *VarcharValue {
	return &VarcharValue{value}
}
func (v *VarcharValue) Bytes() []byte {
	return bytes.NewBufferString(v.value).Bytes()
}

type VariableValue struct {
	ptr  int
	size int
}

func NewVariableValue(ptr int, size int) *VariableValue {
	return &VariableValue{ptr, size}
}
func (v *VariableValue) Bytes() []byte {
	buff := bytes.Buffer{}
	for i := 0; i < v.size; i++ {
		buff.WriteByte(memory[v.ptr+i])
	}
	return buff.Bytes()
}

var Null = NewIntValue(0)
