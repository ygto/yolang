// Code generated from parser/Lang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Lang
import "github.com/antlr4-go/antlr/v4"

// LangListener is a complete listener for a parse tree produced by LangParser.
type LangListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterReservedFunctions is called when entering the reservedFunctions production.
	EnterReservedFunctions(c *ReservedFunctionsContext)

	// EnterFunctions is called when entering the functions production.
	EnterFunctions(c *FunctionsContext)

	// EnterTypes is called when entering the types production.
	EnterTypes(c *TypesContext)

	// EnterValExpr is called when entering the valExpr production.
	EnterValExpr(c *ValExprContext)

	// EnterVariableExpr is called when entering the variableExpr production.
	EnterVariableExpr(c *VariableExprContext)

	// EnterVarExpr is called when entering the varExpr production.
	EnterVarExpr(c *VarExprContext)

	// EnterReturnExpr is called when entering the returnExpr production.
	EnterReturnExpr(c *ReturnExprContext)

	// EnterAssignExpr is called when entering the assignExpr production.
	EnterAssignExpr(c *AssignExprContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterCallFuncExpr is called when entering the callFuncExpr production.
	EnterCallFuncExpr(c *CallFuncExprContext)

	// EnterFuncExpr is called when entering the funcExpr production.
	EnterFuncExpr(c *FuncExprContext)

	// EnterStaticValues is called when entering the staticValues production.
	EnterStaticValues(c *StaticValuesContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitReservedFunctions is called when exiting the reservedFunctions production.
	ExitReservedFunctions(c *ReservedFunctionsContext)

	// ExitFunctions is called when exiting the functions production.
	ExitFunctions(c *FunctionsContext)

	// ExitTypes is called when exiting the types production.
	ExitTypes(c *TypesContext)

	// ExitValExpr is called when exiting the valExpr production.
	ExitValExpr(c *ValExprContext)

	// ExitVariableExpr is called when exiting the variableExpr production.
	ExitVariableExpr(c *VariableExprContext)

	// ExitVarExpr is called when exiting the varExpr production.
	ExitVarExpr(c *VarExprContext)

	// ExitReturnExpr is called when exiting the returnExpr production.
	ExitReturnExpr(c *ReturnExprContext)

	// ExitAssignExpr is called when exiting the assignExpr production.
	ExitAssignExpr(c *AssignExprContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitCallFuncExpr is called when exiting the callFuncExpr production.
	ExitCallFuncExpr(c *CallFuncExprContext)

	// ExitFuncExpr is called when exiting the funcExpr production.
	ExitFuncExpr(c *FuncExprContext)

	// ExitStaticValues is called when exiting the staticValues production.
	ExitStaticValues(c *StaticValuesContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)
}
