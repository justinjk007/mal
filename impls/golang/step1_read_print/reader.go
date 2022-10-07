package reader

import (
	"go/token"
	"regexp"
)

type Reader interface {
	Next()
	Peek()
}

type Read struct {
	position int
	tokens   []string // map of tokens
}

// return current token and increase the position
func (x Read) Next() string {
	x.position++
	return x.tokens[x.position]
}

// return current position
func (x Read) Peek() string {
	return x.tokens[x.position]
}

func () Read_str() {
	tokens_from_ast := Tokenize()
	reader = Read{0, tokens_from_ast}
	read_form(reader)
}

// Function will return a array(map) of strings that was tokenized
// Key starts at 0
func (x string) Tokenize() [][]string {
	musComp := regexp.MustCompile("[\\s,]*(~@|[\\[\\]{}()'`~^@]|\"(?:\\.|[^\\\"])*\"?|;.*|[^\\s\\[\\]{}('\"`,;)]*)")
	tokens_from_ast := musComp.FindAllStringSubmatch(x, -1)
	return tokens_from_ast
}
