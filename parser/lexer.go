package parser

type Lexer struct {
	Input        string
	curPostion   int
	nextPosition int
}

func (l *Lexer) literal() string {
	var token string
	for res := l.getChar(); res != Quotes; {
		token += res
		res = l.getChar()
	}

	return token
}

func lexNumber(input string) string {

	return ""
}

// true, false
func lexBool(input string) string {

	return ""
}

// null, NULL,
func lexNull(input string) string {

	return ""
}

func (l *Lexer) NextToken() string {

	c := l.getChar()
	for c != EOF {
		switch c {
		case Quotes:
			return l.literal()
		case LBraces, RBraces, LBrackets, RBrackets, Colun:
			return c

		case " ", "\t", "\n", "\r":
			c = l.getChar()

		default:
			// parse number

			// parse bool

			// parse null
		}
	}

	return EOF
}

func (l *Lexer) getChar() string {
	if l.curPostion < len(l.Input) {
		c := l.Input[l.curPostion]
		l.curPostion++
		l.nextPosition = l.curPostion + 1
		return string(c)
	}
	return ""
}

func NewLex(input string) *Lexer {
	lex := &Lexer{
		Input: input,
	}

	return lex
}
