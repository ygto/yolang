package main

type Struct struct {
	attributes map[string]*Type
	offsets    map[string]int
}
type Structs map[string]*Struct

var structs Structs = make(Structs)

func (s Structs) Add(name string, val *Struct) {
	s[name] = val
	api.Struct(name, val)
}
func (s Structs) Get(name string) *Struct {
	val, has := s[name]
	if !has {
		fatalf("struct not found: %s\n", name)
	}
	return val
}

func NewStruct() *Struct {
	return &Struct{attributes: make(map[string]*Type), offsets: make(map[string]int)}
}

func (s Struct) AddAttribute(name string, a *Type) {
	s.offsets[name] = s.size()
	s.attributes[name] = a
}

func (s Struct) size() int {
	totalSize := 0
	for _, i := range s.attributes {
		totalSize += i.size
	}
	return totalSize
}

func (s Struct) Type(attr string) *Type {
	val, has := s.attributes[attr]
	if !has {
		fatalf("struct attribute not found: %s\n", attr)
	}
	return val
}
