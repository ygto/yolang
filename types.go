package main

const (
	INT_TYPE     = "int"
	VARCHAR_TYPE = "varchar"
	STRUCT_TYPE  = "struct"
)

type Type struct {
	size int
	t    string
}

type Types map[string]*Type

func (t Types) Add(name string, val *Type) {
	t[name] = val
}

func (t Types) IsPrimitive(name string) bool {
	switch name {
	case "int", "varchar":
		return true
	}

	return false
}

func (t Types) Size(name string) int {
	val := t[name]
	if val == nil {
		fatalf("type %q does not exist", name)
		return 0
	}
	return val.size
}
func (t Types) Type(name string) string {
	val := t[name]
	if val == nil {
		fatalf("type %q does not exist", name)
		return ""
	}
	return val.t
}

var types = make(Types)

var typeSizes = make(map[string]int)

func setTypeSize(t string, size int) {
	typeSizes[t] = size
}

func getTypeSize(t string) int {
	switch t {
	case "int":
		return INT_SIZE
	case "varchar":
		return VARCHAR_SIZE
	}
	return typeSizes[t]
}

func init() {
	types.Add("int", &Type{size: INT_SIZE, t: "int"})
}
