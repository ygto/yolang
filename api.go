package main

import (
	"github.com/llir/llvm/ir"
)
import irTypes "github.com/llir/llvm/ir/types"

var m = ir.NewModule()

type API struct {
	structs  map[string]irTypes.Type
	variable map[string]*ir.Global
	m        *ir.Module
}

func NewAPI() *API {
	return &API{structs: make(map[string]irTypes.Type), m: m, variable: make(map[string]*ir.Global)}
}

var api = NewAPI()

func (api *API) Variable(name string, v *Variable) {
	switch v.t {
	case INT_TYPE:
		api.variable[name] = m.NewGlobal(name, irTypes.NewArray(uint64(v.size), irTypes.I8))

	case VARCHAR_TYPE:
		api.variable[name] = m.NewGlobal(name, irTypes.NewArray(uint64(v.size), irTypes.I8))
	case STRUCT_TYPE:
		s, has := api.structs[name]
		if !has {
			fatalf("struct %s not found", name)
		}
		api.variable[name] = m.NewGlobal(name, s)
	}
}
func (api *API) Assign(name string, v Value) {
}

func (api *API) Struct(name string, s *Struct) {
	fields := make([]irTypes.Type, 0, 0)
	for attrName, attr := range s.attributes {
		switch attr.t {
		case INT_TYPE:
			fields = append(fields, irTypes.I64)
		case VARCHAR_TYPE:
			fields = append(fields, irTypes.NewArray(uint64(s.Type(attrName).size), irTypes.I8))
		case STRUCT_TYPE:
			attrStruct, has := api.structs[attrName]
			if !has {
				fatalf("struct %s not found", attrName)
			}
			fields = append(fields, attrStruct)
		}
	}
	m.NewTypeDef(name, irTypes.NewStruct(fields...))
}
