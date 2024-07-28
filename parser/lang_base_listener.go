// Code generated from parser/Lang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Lang
import "github.com/antlr4-go/antlr/v4"

// BaseLangListener is a complete listener for a parse tree produced by LangParser.
type BaseLangListener struct{}

var _ LangListener = &BaseLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseLangListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseLangListener) ExitProg(ctx *ProgContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseLangListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseLangListener) ExitExpr(ctx *ExprContext) {}

// EnterReservedFunctions is called when production reservedFunctions is entered.
func (s *BaseLangListener) EnterReservedFunctions(ctx *ReservedFunctionsContext) {}

// ExitReservedFunctions is called when production reservedFunctions is exited.
func (s *BaseLangListener) ExitReservedFunctions(ctx *ReservedFunctionsContext) {}

// EnterFunctions is called when production functions is entered.
func (s *BaseLangListener) EnterFunctions(ctx *FunctionsContext) {}

// ExitFunctions is called when production functions is exited.
func (s *BaseLangListener) ExitFunctions(ctx *FunctionsContext) {}

// EnterTypes is called when production types is entered.
func (s *BaseLangListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *BaseLangListener) ExitTypes(ctx *TypesContext) {}

// EnterValExpr is called when production valExpr is entered.
func (s *BaseLangListener) EnterValExpr(ctx *ValExprContext) {}

// ExitValExpr is called when production valExpr is exited.
func (s *BaseLangListener) ExitValExpr(ctx *ValExprContext) {}

// EnterVariableExpr is called when production variableExpr is entered.
func (s *BaseLangListener) EnterVariableExpr(ctx *VariableExprContext) {}

// ExitVariableExpr is called when production variableExpr is exited.
func (s *BaseLangListener) ExitVariableExpr(ctx *VariableExprContext) {}

// EnterVarExpr is called when production varExpr is entered.
func (s *BaseLangListener) EnterVarExpr(ctx *VarExprContext) {}

// ExitVarExpr is called when production varExpr is exited.
func (s *BaseLangListener) ExitVarExpr(ctx *VarExprContext) {}

// EnterReturnExpr is called when production returnExpr is entered.
func (s *BaseLangListener) EnterReturnExpr(ctx *ReturnExprContext) {}

// ExitReturnExpr is called when production returnExpr is exited.
func (s *BaseLangListener) ExitReturnExpr(ctx *ReturnExprContext) {}

// EnterAssignExpr is called when production assignExpr is entered.
func (s *BaseLangListener) EnterAssignExpr(ctx *AssignExprContext) {}

// ExitAssignExpr is called when production assignExpr is exited.
func (s *BaseLangListener) ExitAssignExpr(ctx *AssignExprContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseLangListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseLangListener) ExitBlock(ctx *BlockContext) {}

// EnterCallFuncExpr is called when production callFuncExpr is entered.
func (s *BaseLangListener) EnterCallFuncExpr(ctx *CallFuncExprContext) {}

// ExitCallFuncExpr is called when production callFuncExpr is exited.
func (s *BaseLangListener) ExitCallFuncExpr(ctx *CallFuncExprContext) {}

// EnterFuncExpr is called when production funcExpr is entered.
func (s *BaseLangListener) EnterFuncExpr(ctx *FuncExprContext) {}

// ExitFuncExpr is called when production funcExpr is exited.
func (s *BaseLangListener) ExitFuncExpr(ctx *FuncExprContext) {}

// EnterStaticValues is called when production staticValues is entered.
func (s *BaseLangListener) EnterStaticValues(ctx *StaticValuesContext) {}

// ExitStaticValues is called when production staticValues is exited.
func (s *BaseLangListener) ExitStaticValues(ctx *StaticValuesContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseLangListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseLangListener) ExitExprList(ctx *ExprListContext) {}
