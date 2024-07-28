package main

import (
	"fmt"
	"yolang/parser"

	"github.com/antlr4-go/antlr/v4"
)

type VariableNameType string

type langListener struct {
	*parser.BaseLangListener
	ptrStack   []int
	ptr        int
	scopeStack []Scope
	values     map[string]*Value
}

func NewLangListener() *langListener {
	l := &langListener{
		scopeStack: make([]Scope, 0, 0),
		values:     make(map[string]*Value),
	}
	l.PushScope(l)

	return l
}

func (l *langListener) ActiveScope() Scope {
	return l.scopeStack[len(l.scopeStack)-1]
}

func (l *langListener) PushScope(s Scope) {
	l.scopeStack = append(l.scopeStack, s)
	l.PushPtr()
}

func (l *langListener) PopScope() Scope {
	pop, s1 := l.scopeStack[len(l.scopeStack)-1], l.scopeStack[:len(l.scopeStack)-1]
	l.scopeStack = s1
	l.PopPtr()
	return pop
}

func (l *langListener) GetValue(k string) *Value {
	val, has := l.values[k]
	if !has {
		fmt.Printf("variable not found: %s\n", k)
		return nil
	}
	return val
}

func (l *langListener) SetValue(k string, v interface{}) {
	t := l.GetValue(k)

	fmt.Printf("0x%d : %s = %v(%s)\n", t.Ptr, k, v, t.Type)
	l.values[k] = NewValue(t.Ptr, t.Type, v)
}
func (l *langListener) CreateVariable(k string, t string) {
	l.values[k] = NewValue(l.GetPtr(), t, nil)
	size := getTypeSize(t)
	fmt.Printf("0x%d : var %s = %s[%d] \n", l.GetPtr(), k, t, size)
	l.IncrPtr(size)
}
func (l *langListener) GetRoot() Scope {
	return nil
}

func (l *langListener) PushPtr() {
	l.ptrStack = append(l.ptrStack, l.ptr)
}

func (l *langListener) PopPtr() {
	val, stack := l.ptrStack[len(l.ptrStack)-1], l.ptrStack[:len(l.ptrStack)-1]
	l.ptrStack = stack
	l.ptr = val
}

func (l *langListener) GetPtr() int {
	return l.ptr
}

func (l *langListener) IncrPtr(v int) {
	l.ptr += v
}

func (l *langListener) EnterProg(ctx *parser.ProgContext) {
	i := 0
	_ = i

	for _, tree := range ctx.GetChildren() {
		switch expr := tree.(type) {
		case *parser.ExprContext:
			parseExpr(l, expr)
		case *antlr.TerminalNodeImpl:
			if expr.GetText() == "<EOF>" {
				return
			}
			fmt.Println("unknown terminal node", expr.GetText())
		default:

		}
		i += 1
	}
}

func parseExpr(l *langListener, ctx *parser.ExprContext) interface{} {

	i := 0
	_ = i

	for _, tree := range ctx.GetChildren() {
		switch expr := tree.(type) {
		case *parser.VarExprContext:
			parseVarExpr(l, expr)
		case *parser.AssignExprContext:
			parseAssignExpr(l, expr)
		case *parser.FuncExprContext:
			parseFuncExpr(l, expr)
		case *parser.CallFuncExprContext:
			parseCallFuncExpr(l, expr)
		case *parser.ReturnExprContext:
			return parseReturnExpr(l, expr)
		case *parser.FunctionsContext:
			return parseFunctionExpr(l, expr)
		default:
			fmt.Println("unknown type x0", i, structName(tree))
		}
		i += 1
	}
	return nil
}

func parseFunctionExpr(l *langListener, expr *parser.FunctionsContext) interface{} {
	reservedFunc := parseReservedFunc(expr.GetChild(0))
	value := parseValExpr(l, expr.GetChild(2))

	switch name := value.(type) {
	case VariableNameType:
		value = l.ActiveScope().GetValue(string(name)).Val
	default:
		//fmt.Println("unknown type x1", structName(value), value)
	}
	switch reservedFunc {
	case "print":
		fmt.Printf("func print: %v\n", value)
		return 0
	}

	return value
}

func parseReservedFunc(tree antlr.Tree) string {
	switch t := tree.(type) {
	case *parser.ReservedFunctionsContext:
		return t.GetText()
	default:
		fmt.Println("parseReservedFunc default", structName(t))
	}
	fatalf("parseReservedFunc unnown type error:%v", tree.GetPayload())
	return ""
}

func parseReturnExpr(l *langListener, expr *parser.ReturnExprContext) interface{} {
	value := parseValExpr(l, expr.GetChild(1))

	switch name := value.(type) {
	case VariableNameType:
		value = l.ActiveScope().GetValue(string(name)).Val

	default:
		//fmt.Println("unknown type x12", structName(value))
	}
	return value

}

func parseFuncExpr(l *langListener, expr *parser.FuncExprContext) {
	funcName := parseFuncNameExpr(expr.GetChild(1))
	paramType := parseTypesExpr(expr.GetChild(3))
	paramName := parseVariableExpr(expr.GetChild(4))
	returnType := parseTypesExpr(expr.GetChild(6))

	f := NewFunction(l, funcName, l.GetPtr(), returnType, paramType, string(paramName), expr)
	l.ActiveScope().CreateVariable(funcName, "function")
	l.ActiveScope().SetValue(funcName, f)
}

func parseAssignExpr(l *langListener, expr *parser.AssignExprContext) {

	variableName := parseVariableExpr(expr.GetChild(0))
	value := parseValExpr(l, expr.GetChild(2))
	switch val := value.(type) {
	case VariableNameType:
		value = l.ActiveScope().GetValue(string(val)).Val
	}

	l.ActiveScope().SetValue(string(variableName), value)
}

func parseValExpr(l *langListener, ctx antlr.Tree) interface{} {
	i := 0
	_ = i

	for _, tree := range ctx.GetChildren() {
		switch expr := tree.(type) {
		case *parser.StaticValuesContext:
			return parseStaticValuesExpr(expr)
		case *parser.CallFuncExprContext:
			return parseCallFuncExpr(l, expr)
		case *parser.VariableExprContext:
			name := parseVariableExpr(expr)
			val := l.ActiveScope().GetValue(string(name))
			return val.Val
		default:

			fmt.Println("parseValExpr unknown type", i, structName(tree))
		}
		i += 1
	}
	return nil
}

func parseCallFuncExpr(l *langListener, expr *parser.CallFuncExprContext) interface{} {
	funcName := parseFuncNameExpr(expr.GetChild(1))
	function := l.ActiveScope().GetValue(funcName)
	if function.Type != "function" {
		fatalf("function type mismatch:%s", funcName)
	}
	callParam := string(parseVariableExpr(expr.GetChild(3)))
	callValue := l.ActiveScope().GetValue(callParam)
	functionBody := function.Val.(*Function)

	l.PushScope(functionBody)

	fmt.Printf("0x%d : call %s(%s %s) %s \n", functionBody.ptr, funcName, callValue.Type, callValue.Val, functionBody.returnType)

	l.ActiveScope().CreateVariable(functionBody.param, functionBody.paramType)
	l.ActiveScope().SetValue(functionBody.param, callValue.Val)
	returnVal := callFuncExpr(l, functionBody.clousure.GetChild(7).(*parser.BlockContext))

	l.PopScope()
	fmt.Printf("0x%d : call %s(%s %s) = %v \n", functionBody.ptr, funcName, callValue.Type, callValue.Val, returnVal)

	return returnVal
}

func callFuncExpr(l *langListener, ctx *parser.BlockContext) interface{} {
	var i = 0
	for _, tree := range ctx.GetChildren() {
		switch t := tree.(type) {
		case *parser.ExprContext:
			val := parseExpr(l, t)
			if val != nil {
				return val
			}
		case *antlr.TerminalNodeImpl:

		default:

			fmt.Println("callFuncExpr unknown type", i, structName(tree))
		}
		i++
	}
	return nil
}

func parseFuncNameExpr(tree antlr.Tree) string {
	switch t := tree.(type) {
	case *antlr.TerminalNodeImpl:
		return t.GetText()
	default:
		fmt.Println("parseFuncNameExpr unknown type error", structName(t))
	}
	fatalf("parseFuncNameExpr unknown type error:%v", tree.GetPayload())
	return ""
}

func parseStaticValuesExpr(ctx *parser.StaticValuesContext) interface{} {
	return ctx.GetText()
}

func parseVarExpr(l *langListener, expr *parser.VarExprContext) {
	if expr.GetChildCount() != 3 {
		fatalf("parameter mismatch:%s", expr.GetText())
	}
	variableName := string(parseVariableExpr(expr.GetChild(1)))
	variableType := parseTypesExpr(expr.GetChild(2))
	l.ActiveScope().CreateVariable(variableName, variableType)
}

func parseVariableExpr(tree antlr.Tree) VariableNameType {
	switch t := tree.(type) {
	case *parser.VariableExprContext:
		return VariableNameType(t.GetChild(1).GetPayload().(*antlr.CommonToken).GetText())
	default:
		fmt.Println("parseVariableExpr unknown type error", structName(t))
	}
	fatalf("parseVariableExpr unknown type error:%v", tree.GetPayload())
	return ""
}

func parseTypesExpr(tree antlr.Tree) string {
	switch t := tree.(type) {
	case *parser.TypesContext:
		return t.GetText()

	default:
		fmt.Println("parseTypesExpr uknnown type error", structName(t))
	}
	fatalf("parseTypesExpr uknnown type error:%v", tree.GetPayload())

	return ""
}
