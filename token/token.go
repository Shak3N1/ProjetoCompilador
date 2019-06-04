// Code generated by gocc; DO NOT EDIT.

package token

import (
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func (p Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", p.Offset, p.Line, p.Column)
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (m TokenMap) Id(tok Type) string {
	if int(tok) < len(m.typeMap) {
		return m.typeMap[tok]
	}
	return "unknown"
}

func (m TokenMap) Type(tok string) Type {
	if typ, exist := m.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (m TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", m.Id(tok.Type), tok.Type, tok.Lit)
}

func (m TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", m.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"empty",
		"var",
		"minus",
		"endline",
		"identifier",
		"comma",
		"float",
		"int",
		"char",
		"void",
		"proc",
		"lbracket",
		"rbracket",
		"lbrace",
		"rbrace",
		"assign",
		"prnt",
		"return",
		"while",
		"arithmeticSymbol",
		"char_literal",
		"else",
		"fp_literal",
		"if",
		"int_literal",
		"not",
		"or",
		"relacionalSymbol",
		"sn_literal",
	},

	idMap: map[string]Type{
		"INVALID":          0,
		"$":                1,
		"empty":            2,
		"var":              3,
		"minus":            4,
		"endline":          5,
		"identifier":       6,
		"comma":            7,
		"float":            8,
		"int":              9,
		"char":             10,
		"void":             11,
		"proc":             12,
		"lbracket":         13,
		"rbracket":         14,
		"lbrace":           15,
		"rbrace":           16,
		"assign":           17,
		"prnt":             18,
		"return":           19,
		"while":            20,
		"arithmeticSymbol": 21,
		"char_literal":     22,
		"else":             23,
		"fp_literal":       24,
		"if":               25,
		"int_literal":      26,
		"not":              27,
		"or":               28,
		"relacionalSymbol": 29,
		"sn_literal":       30,
	},
}
