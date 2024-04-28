package lexer

import (
	"github.com/jeffreyqdd/go-monkey/internal/token"
)

type Lexer struct {
	input   string // target string to tokenize
	pointer int    // index of the character we have read
	ch      byte   // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input, pointer: -1, ch: 0}
	l.readChar()
	return l
}

// NextToken
func (l *Lexer) NextToken() token.Token {

	// Monkey is white-space agnostic
	l.skipWhitespace()

	var tok token.Token
	switch l.ch {
	// operators
	case '=':
		switch l.peekChar() == '=' {
		case true: // ==
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: token.EQ}
		case false: // =
			tok = token.Token{Type: token.ASSIGN, Literal: token.ASSIGN}
		}
	case '!':
		switch l.peekChar() == '=' {
		case true: // !=
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: token.NOT_EQ}
		case false: // !
			tok = token.Token{Type: token.BANG, Literal: token.BANG}
		}

	case '+':
		tok = token.Token{Type: token.PLUS, Literal: token.PLUS}
	case '-':
		tok = token.Token{Type: token.MINUS, Literal: token.MINUS}
	case '*':
		tok = token.Token{Type: token.ASTERISK, Literal: token.ASTERISK}
	case '/':
		tok = token.Token{Type: token.SLASH, Literal: token.SLASH}
	case '<':
		tok = token.Token{Type: token.LT, Literal: token.LT}
	case '>':
		tok = token.Token{Type: token.GT, Literal: token.GT}

	// delimiters
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: token.COMMA}
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: token.SEMICOLON}
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: token.LPAREN}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: token.RPAREN}
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: token.LBRACE}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: token.RBRACE}
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}

	// Idents, Literals, and Keywords
	// If first digit is letter -> ident/keywords
	// If first digit is number -> literals
	default:

		KEYWORDS := map[string]token.TokenType{
			"fn":     token.FUNCTION,
			"let":    token.LET,
			"true":   token.TRUE,
			"false":  token.FALSE,
			"if":     token.IF,
			"else":   token.ELSE,
			"return": token.RETURN,
		}

		if l.isLetter() {
			literal := l.readLiteral() // need to check if it is a keyword

			literal_token, ok := KEYWORDS[literal]

			if ok {
				// keyword
				tok = token.Token{Type: literal_token, Literal: literal}
			} else {
				// ident
				tok = token.Token{Type: token.IDENT, Literal: literal}
			}
		} else if l.isDigit() {
			tok = token.Token{Type: token.INT, Literal: l.readLiteral()}
		} else {
			tok = token.Token{Type: token.ILLEGAL, Literal: ""}
		}
		return tok
	}

	l.readChar()
	return tok
}

// readChar reads the next character in l.input, stores the character in l.ch, and updates
// l.pointer to reflect the index of that character. If there is no next character, the value 0
// (ASCII NULL), but l.pointer is still advanced
func (l *Lexer) readChar() {
	nextPosition := l.pointer + 1

	if nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[nextPosition]
	}

	l.pointer = nextPosition
}

func (l *Lexer) peekChar() byte {
	nextPosition := l.pointer + 1

	if nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[nextPosition]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) isLetter() bool {
	return ('a' <= l.ch && l.ch <= 'z') || ('A' <= l.ch && l.ch <= 'Z') || l.ch == '_'
}

func (l *Lexer) isDigit() bool {
	return 0x30 <= l.ch && l.ch <= 0x39
}

func (l *Lexer) readLiteral() string {
	startPosition := l.pointer

	for l.isLetter() || l.isDigit() {
		l.readChar()
	}

	endPosition := l.pointer
	return l.input[startPosition:endPosition]
}
