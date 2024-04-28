package lexer

import (
	"testing"

	"github.com/jeffreyqdd/go-monkey/internal/token"
)

type Expectation struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestOperatorSamples(t *testing.T) {
	input := `=+-!*/<>==!=`
	tests := []Expectation{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.BANG, "!"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.EQ, "=="},
		{token.NOT_EQ, "!="},
	}

	assert_input_and_expected_equal(input, tests, t)
}

func TestNextMonkeyCodeSample(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn             (x,y) {
	x+             y                 ;
};
let result=add(five,ten);
`
	tests := []Expectation{

		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	assert_input_and_expected_equal(input, tests, t)
}

func TestPeek(t *testing.T) {
	// peak should not cause an array access out of bounds
	input := `=`
	tests := []Expectation{
		{token.ASSIGN, "="},
		{token.EOF, ""},
		{token.EOF, ""},
		{token.EOF, ""},
	}
	assert_input_and_expected_equal(input, tests, t)
}

func TestIllegalToken(t *testing.T) {
	input := `he#`
	tests := []Expectation{
		{token.IDENT, "he"},
		{token.ILLEGAL, ""},
	}

	assert_input_and_expected_equal(input, tests, t)
}

func assert_input_and_expected_equal(input string, expectations []Expectation, t *testing.T) {
	l := New(input)

	for idx, tt := range expectations {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				idx, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", idx, tt.expectedLiteral, tok.Literal)
		}
	}
}
