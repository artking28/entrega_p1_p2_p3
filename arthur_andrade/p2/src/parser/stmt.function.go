package parser

import (
	models "compilers/sharedModels"
	"compilers/utils"
)

type FuncStmt struct {
	Name string
	Body Scope
	MantisStmtBase
}

func (this FuncStmt) WriteMemASM() (string, error) {
	//TODO implement me
	panic("implement me | FuncStmt@WriteMemASM")
}

func NewFuncStmt(name string, body Scope, pos utils.Pos, parser *MantisParser) *FuncStmt {
	return &FuncStmt{
		Name: name,
		Body: body,
		MantisStmtBase: MantisStmtBase{
			Parser: parser,
			Title:  "FuncStmt",
			Pos:    pos,
		},
	}
}

func (parser *MantisParser) ParseFunction() (ret *FuncStmt, err error) {
	h0 := parser.Get(0)
	if h0 == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}
	parser.Consume(1)
	nameTk, _ := parser.HasNextConsume(OptionalSpaceMode, models.SPACE, models.ID)
	if nameTk == nil {
		return nil, utils.GetExpectedTokenErr(parser.Filename, "function name", h0.Pos)
	}
	if _, err = parser.HasNextConsume(OptionalSpaceMode, models.SPACE, models.L_PAREN); err != nil {
		return nil, utils.GetExpectedTokenErrOr(parser.Filename, "left parenthesis", err.Error(), h0.Pos)
	}
	if _, err = parser.HasNextConsume(OptionalSpaceMode, models.SPACE, models.R_PAREN); err != nil {
		return nil, utils.GetExpectedTokenErrOr(parser.Filename, "right parenthesis", err.Error(), h0.Pos)
	}
	if _, err = parser.HasNextConsume(OptionalSpaceMode, models.SPACE, models.L_BRACE); err != nil {
		return nil, utils.GetExpectedTokenErrOr(parser.Filename, "left brace", err.Error(), h0.Pos)
	}
	ast, err := parser.ParseScope(FuncScope)
	if err != nil {
		return nil, err
	}
	if _, err = parser.HasNextConsume(OptionalSpaceMode, models.SPACE, models.R_BRACE); err != nil {
		return nil, utils.GetExpectedTokenErrOr(parser.Filename, "right brace", err.Error(), h0.Pos)
	}
	return NewFuncStmt(string(nameTk.Value), ast, h0.Pos, parser), nil
}
