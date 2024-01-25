package lexer

import "github.com/eduardolat/edulang/token"

// Lexer is a struct that holds the source code input and
// converts it into tokens.
type Lexer struct {
	// input is the source code input.
	input string
	// currentPosition is the current position in the input.
	currentPosition int
	// currentPositionByte is the current byte in the input.
	currentPositionByte byte
	// nextPosition is the next position to be readed from the input.
	nextPosition int
}

// New returns a new Lexer.
func New(input string) *Lexer {
	lx := Lexer{input: input}
	lx.readChar()
	return &lx
}

// readChar reads the next byte from the input and updates the
// current and next positions.
func (l *Lexer) readChar() {
	hasNextPosition := l.nextPosition < len(l.input)

	if hasNextPosition {
		l.currentPositionByte = l.input[l.nextPosition]
	}

	if !hasNextPosition {
		l.currentPositionByte = 0
	}

	l.currentPosition = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	t := token.Token{}

	currentPositionStr := string(l.currentPositionByte)

	if currentPositionStr == "=" {
		t.Type = token.ASSIGN
		t.Literal = currentPositionStr
	}

	if currentPositionStr == ";" {
		t.Type = token.SEMICOLON
		t.Literal = currentPositionStr
	}

	if currentPositionStr == "(" {
		t.Type = token.LPAREN
		t.Literal = currentPositionStr
	}

	if currentPositionStr == ")" {
		t.Type = token.RPAREN
		t.Literal = currentPositionStr
	}

	if currentPositionStr == "," {
		t.Type = token.COMMA
		t.Literal = currentPositionStr
	}

	if currentPositionStr == "+" {
		t.Type = token.PLUS
		t.Literal = currentPositionStr
	}

	if currentPositionStr == "{" {
		t.Type = token.LBRACE
		t.Literal = currentPositionStr
	}

	if currentPositionStr == "}" {
		t.Type = token.RBRACE
		t.Literal = currentPositionStr
	}

	if l.currentPositionByte == 0 {
		t.Type = token.EOF
		t.Literal = ""
	}

	l.readChar()
	return t
}
