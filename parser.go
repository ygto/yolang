package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
	"strings"
	"yolang/parser"
)

func parseExpr(ctx *parser.ExprContext) {
	i := 0
	for _, tree := range ctx.GetChildren() {
		switch expr := tree.(type) {
		case *parser.ImportFileContext:
			parseImportFile(expr)
		case *parser.VariableDefinitionContext:
			parseVariableDefinition(expr)
		case *parser.AssignContext:
			parseAssign(expr)
		case *parser.StructDeclContext:
			parseStructDecl(expr)
		default:
			fmt.Println("parseExpr unknown type:", i, structName(expr))
		}
		i += 1
	}
}

func parseStructDecl(ctx *parser.StructDeclContext) {
	i := 0
	s := NewStruct()
	name := getTerminalNodeImplValue(ctx.GetChild(1).(*antlr.TerminalNodeImpl))

	children := ctx.GetChildren()
	for _, tree := range children[4 : len(children)-1] {
		switch expr := tree.(type) {
		case *parser.VariableDefinitionContext:
			attrName, t := getStructVariableDefinition(expr)
			s.AddAttribute(attrName, t)

		default:
			fmt.Println("parseStructDecl unknown type:", i, structName(expr))
		}
		i += 1
	}
	structs.Add(name, s)
	types.Add(name, &Type{t: name, size: s.size()})

	fmt.Println(name)
}

func parseAssign(ctx *parser.AssignContext) {
	i := 0
	var varName string
	_ = varName
	var varPtr int
	var varType string
	var value Value
	for _, tree := range ctx.GetChildren() {
		switch expr := tree.(type) {
		case *parser.VariableContext:
			name := getVariableFromVariableContext(expr)
			variable := variables.Get(name)
			varPtr = variable.ptr
			varType = variable.t
		case *parser.StructVariableContext:

			varName, varPtr, varType = getStructVariableFromVariableContext(expr)
		case *parser.ValueContext:
			value = getValueFromValueContext(expr, varType)
		case *antlr.TerminalNodeImpl:
			val := getTerminalNodeImplValue(expr)
			if val == "=" {
				continue
			}

		default:
			fmt.Println("parseAssign unknown type:", i, structName(expr))
		}
		i += 1
	}
	fmt.Printf("varPtr: %d\n", varPtr)
	variables.SetByPtr(varPtr, value)
}

func getStructVariableFromVariableContext(ctx *parser.StructVariableContext) (string, int, string) {
	i := 0
	name := getVariableFromVariableContext(ctx.GetChild(0).(*parser.VariableContext))
	variable := variables.Get(name)
	s := structs.Get(variable.t)
	var offset int
	var t string
	var size int
	children := ctx.GetChildren()
	for _, tree := range children[1:] {
		switch expr := tree.(type) {
		case *antlr.TerminalNodeImpl:
			attrName := getTerminalNodeImplValue(expr)
			if attrName == "." {
				continue
			}
			attr := s.Type(attrName)
			t = attr.t
			size = attr.size
			if !types.IsPrimitive(attr.t) {
				s = structs.Get(attr.t)
				size = attr.size
			}
			offset += attr.size

		default:
			fmt.Println("getStructVariableFromVariableContext unknown type:", i, structName(expr))
		}
		i += 1
	}
	return name, variable.ptr + offset - size, t
}

func getValueFromValueContext(ctx *parser.ValueContext, expectedType string) Value {
	i := 0
	for _, tree := range ctx.GetChildren() {
		switch expr := tree.(type) {
		case *parser.StaticValueContext:
			return getStaticValueContext(expr, expectedType)
		case *parser.VariableContext:
			name := getVariableFromVariableContext(expr)
			variable := variables.Get(name)
			if variable.t != expectedType {
				fatalf("expected type %s, got %s", expectedType, variable.t)
			}
			return NewVariableValue(variable.ptr, variable.size)

		default:
			fmt.Println("getValueContext unknown type:", i, structName(expr))
		}
		i += 1
	}
	return Null
}

func getStaticValueContext(ctx *parser.StaticValueContext, expectedType string) Value {
	i := 0
	for _, tree := range ctx.GetChildren() {
		switch expr := tree.(type) {
		case *antlr.TerminalNodeImpl:
			valStr := getTerminalNodeImplValue(expr)
			switch expectedType {
			case INT_TYPE:
				val, err := strconv.Atoi(valStr)
				if err != nil {
					fatalf("getStaticValueContext unknown type: %s", expectedType)
				}
				return NewIntValue(val)
			case VARCHAR_TYPE:
				return NewVarcharValue(valStr)
			}
		default:
			fmt.Println("getStaticValueContext unknown type:", i, structName(expr))
		}
		i += 1
	}
	return Null
}

func parseImportFile(ctx *parser.ImportFileContext) {
	path := getTerminalNodeImplValue(ctx.GetChild(2).(*antlr.TerminalNodeImpl))
	err := imports.importFile(path)
	if err != nil {
		panic(err)
	}
}

func getStructVariableDefinition(ctx *parser.VariableDefinitionContext) (string, *Type) {
	i := 0

	var name string = ""
	var size int
	attr := &Type{}
	for _, tree := range ctx.GetChildren() {
		switch expr := tree.(type) {
		case *parser.TypesContext:
			attr.t = getTypeFromTypesContext(expr)
			attr.size = types.Size(attr.t)

		case *parser.VariableContext:
			name = getVariableFromVariableContext(expr)
		case *parser.VarcharDeclContext:
			name, size = getVariableFromVarcharDeclContext(expr)
			attr.t = "varchar"
			attr.size = size
		default:
			fmt.Println("parseVariableDefinition unknown type:", i, structName(expr))
		}
		i += 1
	}
	return name, attr
}
func parseVariableDefinition(ctx *parser.VariableDefinitionContext) {
	i := 0
	v := NewVariable()
	var name string = ""
	var size int

	for _, tree := range ctx.GetChildren() {
		switch expr := tree.(type) {
		case *parser.TypesContext:
			t := getTypeFromTypesContext(expr)
			v.t = t
			v.size = types.Size(t)
		case *parser.VariableContext:
			name = getVariableFromVariableContext(expr)
		case *parser.VarcharDeclContext:
			name, size = getVariableFromVarcharDeclContext(expr)
			v.t = "varchar"
			v.size = size
		default:
			fmt.Println("parseVariableDefinition unknown type:", i, structName(expr))
		}
		i += 1
	}
	variables.Add(name, v)
}

func getVariableFromVarcharDeclContext(ctx *parser.VarcharDeclContext) (string, int) {

	sizeStr := getTerminalNodeImplValue(ctx.GetChild(2).(*antlr.TerminalNodeImpl))
	name := getVariableFromVariableContext(ctx.GetChild(4).(*parser.VariableContext))

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		fatalf("getVariableFromVarcharDeclContext unknown type: %s", sizeStr)
	}
	return name, size
}

func getVariableFromVariableContext(ctx *parser.VariableContext) string {
	return getTerminalNodeImplValue(ctx.GetChild(1).(*antlr.TerminalNodeImpl))
}

func getTypeFromTypesContext(ctx *parser.TypesContext) string {
	i := 0
	for _, tree := range ctx.GetChildren() {
		switch expr := tree.(type) {
		case *antlr.TerminalNodeImpl:
			return getTerminalNodeImplValue(expr)
		case *parser.PrimitiveTypesContext:
			return getPrimitiveType(expr)
		default:
			fmt.Println("getTypeFromTypesContext unknown type:", i, structName(expr))
		}
		i += 1
	}
	return ""
}

func getPrimitiveType(ctx *parser.PrimitiveTypesContext) string {
	i := 0
	for _, tree := range ctx.GetChildren() {
		switch expr := tree.(type) {
		case *antlr.TerminalNodeImpl:
			return getTerminalNodeImplValue(expr)
		default:
			fmt.Println("getPrimitiveType unknown type:", i, structName(expr))
		}
		i += 1
	}
	return ""
}

func getTerminalNodeImplValue(expr *antlr.TerminalNodeImpl) string {

	val := expr.GetText()
	val = strings.TrimRight(val, "'")
	val = strings.TrimLeft(val, "'")
	return val
}
