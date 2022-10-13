package reader

import (
	"regexp"
	"fmt"
)

type Reader interface {
	Next()
	Peek()
}

type Read struct {
	tokens   []string // map of tokens
	position int
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
	reader = Read{tokens_from_ast,0}
	read_form(reader)
}

// Function will return a array(map) of strings that was tokenized
// Key starts at 0
func (x string) Tokenize() []string {
	musComp := regexp.MustCompile("[\\s,]*(~@|[\\[\\]{}()'`~^@]|\"(?:\\.|[^\\\"])*\"?|;.*|[^\\s\\[\\]{}('\"`,;)]*)")
	tokens_from_ast := musComp.FindAllString(x, -1) // -1 means go through the entire string
	return tokens_from_ast
}

func () Dummy() {
	fmt.Println("dummy")
}
