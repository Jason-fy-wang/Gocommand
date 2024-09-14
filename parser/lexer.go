package parser

import (
	"slices"
	"strings"
)

type Lexer struct {
	Input        string
	curPostion   int
	nextPosition int
	CurChar      string
}

func (l *Lexer) literal() string {
	var token string = ""

	if res := l.curChar(); res != Quotes {
		return token
	}

	for res := l.getChar(); res != Quotes; {
		token += res
		res = l.getChar()
	}
	l.getChar()

	return token
}

func (l *Lexer) lexNumber() string {

	nums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	var num string = ""
	for slices.Contains(nums, l.curChar()) {
		num += l.curChar()
		l.getChar()
	}

	return num
}

// true, false
func (l *Lexer) lexBool() string {
	var res string = ""
	if len(l.Input[l.curPostion:]) >= TRUE_LEN && strings.ToLower(string(l.Input[l.curPostion:l.curPostion+TRUE_LEN])) == "true" {
		res = string(l.Input[l.curPostion : l.curPostion+TRUE_LEN])
		l.getChar()
		l.getChar()
		l.getChar()
		l.getChar()
	} else if len(l.Input[l.curPostion:]) >= FALSE_LEN && strings.ToLower(string(l.Input[l.curPostion:l.curPostion+FALSE_LEN])) == "false" {
		res = string(l.Input[l.curPostion : l.curPostion+FALSE_LEN])
		l.getChar()
		l.getChar()
		l.getChar()
		l.getChar()
		l.getChar()
	}

	return res
}

// null, NULL,
func (l *Lexer) lexNull() string {
	var res = ""
	if len(l.Input[l.curPostion:]) >= NULL_LEN && strings.ToLower(string(l.Input[l.curPostion:l.curPostion+NULL_LEN])) == "null" {
		res = string(l.Input[l.curPostion : l.curPostion+NULL_LEN])
		l.getChar()
		l.getChar()
		l.getChar()
		l.getChar()
	}
	return res
}

func (l *Lexer) NextToken() string {
	for l.curChar() != EOF {
		switch l.curChar() {
		case Quotes:
			return l.literal()
		case LBraces, RBraces, LBrackets, RBrackets, Colun:
			curChar := l.curChar()
			l.getChar()
			return curChar

		case " ", "\t", "\n", "\r":
			l.getChar()

		default:
			// parse number
			if token := l.lexNumber(); token != "" {
				return token
			}
			// parse bool
			if token := l.lexBool(); token != "" {
				return token
			}
			// parse null
			if token := l.lexNull(); token != "" {
				return token
			}
		}
	}

	return EOF
}

func (l *Lexer) getChar() string {
	if l.nextPosition < len(l.Input) {
		l.CurChar = string(l.Input[l.nextPosition])
		l.curPostion = l.nextPosition
		l.nextPosition = l.curPostion + 1
		return l.CurChar
	} else {
		l.CurChar = EOF
	}
	return l.CurChar
}

func (l *Lexer) curChar() string {
	return l.CurChar
}

func (l *Lexer) peekChar() string {
	if l.nextPosition < len(l.Input) {
		return string(l.Input[l.nextPosition])
	}
	return ""
}
func NewLex(input string) *Lexer {
	lex := &Lexer{
		Input: input,
	}
	lex.getChar()
	return lex
}
