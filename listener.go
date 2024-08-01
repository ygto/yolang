package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"yolang/parser"
)

type YoLangListener struct {
	*parser.BaseYoLangListener
}

func NewYoLangListener() *YoLangListener {
	return new(YoLangListener)
}

func (l *YoLangListener) EnterProg(ctx *parser.ProgContext) {
	i := 0
	_ = i

	for _, tree := range ctx.GetChildren() {
		switch expr := tree.(type) {
		case *parser.ExprContext:
			parseExpr(expr)
		case *antlr.TerminalNodeImpl:
			if expr.GetText() == "<EOF>" {
				continue
			}
			fmt.Println("unknown terminal node", expr.GetText())
		default:

		}
		i += 1
	}
}
