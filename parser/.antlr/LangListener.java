// Generated from e:/dev/mylang/parsing/Lang.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link LangParser}.
 */
public interface LangListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link LangParser#prog}.
	 * @param ctx the parse tree
	 */
	void enterProg(LangParser.ProgContext ctx);
	/**
	 * Exit a parse tree produced by {@link LangParser#prog}.
	 * @param ctx the parse tree
	 */
	void exitProg(LangParser.ProgContext ctx);
	/**
	 * Enter a parse tree produced by {@link LangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr(LangParser.ExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link LangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr(LangParser.ExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link LangParser#funcExpr}.
	 * @param ctx the parse tree
	 */
	void enterFuncExpr(LangParser.FuncExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link LangParser#funcExpr}.
	 * @param ctx the parse tree
	 */
	void exitFuncExpr(LangParser.FuncExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link LangParser#varExpr}.
	 * @param ctx the parse tree
	 */
	void enterVarExpr(LangParser.VarExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link LangParser#varExpr}.
	 * @param ctx the parse tree
	 */
	void exitVarExpr(LangParser.VarExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link LangParser#typeExpr}.
	 * @param ctx the parse tree
	 */
	void enterTypeExpr(LangParser.TypeExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link LangParser#typeExpr}.
	 * @param ctx the parse tree
	 */
	void exitTypeExpr(LangParser.TypeExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link LangParser#exprList}.
	 * @param ctx the parse tree
	 */
	void enterExprList(LangParser.ExprListContext ctx);
	/**
	 * Exit a parse tree produced by {@link LangParser#exprList}.
	 * @param ctx the parse tree
	 */
	void exitExprList(LangParser.ExprListContext ctx);
}