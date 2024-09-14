package parser

func parseObject(l *Lexer, results *map[string]interface{}) {

	key := l.NextToken()

}

func parseArray(l *Lexer) {

}

func Parser(l *Lexer, results *map[string]interface{}) {
	// parse string into map

	token := l.NextToken()

	switch token {
	case LBraces:
		parseObject(l, results)

	case LBrackets:
		parseArray(l)

	default:

	}

}
