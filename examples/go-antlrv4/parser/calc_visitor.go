// Code generated from ./Calc.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CalcParser.
type CalcVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalcParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by CalcParser#expr.
	VisitExpr(ctx *ExprContext) interface{}
}
