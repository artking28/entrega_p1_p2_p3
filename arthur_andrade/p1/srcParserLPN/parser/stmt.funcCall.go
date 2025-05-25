package parser

import (
	"exps-heredia/models"
	"exps-heredia/utils"
	"fmt"
)

type FuncCall struct {
	Name string `json:"name"`
	Arg  IExp   `json:"arg"`
	StmtBase
}

func (this FuncCall) WriteMemASM() (ret string, err error) {
	if this.Name == "print" {
		ret, err = this.Arg.WriteMemASM()
		if err != nil {
			return "", err
		}
		switch this.Arg.GetTitle() {
		case "IdExp":
			ret += fmt.Sprintf("GET %s\n", this.Arg.String())
			break
		case "ExpChain":
			ret += fmt.Sprintf("GET m99\n")
			break
		case "VExp":
			break
		default:
			ret += fmt.Sprintf("GET m99\n")
		}
		ret += fmt.Sprintf("RES m200\n")
	} else if this.Name == "exit" {
		ret = "HLT\n"
	}
	return ret, nil
}

func NewFuncCall(name string, arg IExp, pos utils.Pos, parser *Parser) *FuncCall {
	return &FuncCall{
		Name: name,
		Arg:  arg,
		StmtBase: StmtBase{
			Parser: parser,
			Title:  "FuncCall",
			Pos:    pos,
		}}
}

func (parser *Parser) ParseFuncCall() (ret *FuncCall, err error) {
	nameTk := parser.Get(0)
	if nameTk == nil {
		return nil, utils.GetUnexpectedTokenNoPosErr(parser.Filename, "EOF")
	}

	parser.Consume(1)
	if _, err = parser.HasNextConsume(NoSpaceMode, models.SPACE, models.L_PAREN); err != nil {
		return nil, utils.GetExpectedTokenErrOr(parser.Filename, "left parenthesis", err.Error(), nameTk.Pos)
	}

	arg, err := parser.ParseExpression(true)
	if err != nil {
		return nil, err
	}

	if _, err = parser.HasNextConsume(OptionalSpaceMode, models.SPACE, models.R_PAREN); err != nil {
		return nil, utils.GetExpectedTokenErrOr(parser.Filename, "right parenthesis", err.Error(), nameTk.Pos)
	}

	return NewFuncCall(string(nameTk.Value), arg, nameTk.Pos, parser), nil
}
