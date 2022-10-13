package reader

import (
	"container/list"
	"fmt"
	"mal/types"
	"regexp"
	"strconv"
)

type Reader interface {
	Next()
	Peek()
}

type Read struct {
	tokens   []string // map of tokens
	position int
}

type Mal struct {
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

func Read_str(x string) {
	tokens_from_ast := Tokenize(x)
	reader := Read{tokens_from_ast, 0}
	Read_form(reader)
}

// Add the function read_form to reader.qx. This function will peek at the first
// token in the Reader object and switch on the first character of that
// token. If the character is a left paren then read_list is called with the
// Reader object. Otherwise, read_atom is called with the Reader Object. The
// return value from read_form is a mal data type. If your target language is
// statically typed then you will need some way for read_form to return a
// variant or subclass type. For example, if your language is object oriented,
// then you can define a top level MalType (in types.qx) that all your mal data
// types inherit from. The MalList type (which also inherits from MalType) will
// contain a list/array of other MalTypes. If your language is dynamically typed
// then you can likely just return a plain list/array of other mal types.
func Read_form(reader Read) MalType {
	token := reader.Peek()
	if token == "(" {
		Read_list(reader)
	} else {
		Read_atom(reader)
	}
	return mal
}

func Read_list(reader Read) {
	list := list.New()
	for reader.Next() != ")" {
		token := Read_form(reader) // Maltype
		list.PushBack(token)
	}
}

// This function will look at the contents of the token and return the
// appropriate scalar (simple/single) data type value.
func Read_atom(reader Read) {
	token := reader.Peek()
	val, err := strconv.Atoi(token)
	if err != nil {
		// fmt.Printf("Value %s is not a number\n", x)
	} else {
		return val
	}
}

// Function will return a array(map) of strings that was tokenized
// Key starts at 0
func Tokenize(x string) []string {
	musComp := regexp.MustCompile("[\\s,]*(~@|[\\[\\]{}()'`~^@]|\"(?:\\.|[^\\\"])*\"?|;.*|[^\\s\\[\\]{}('\"`,;)]*)")
	tokens_from_ast := musComp.FindAllString(x, -1) // -1 means go through the entire string
	return tokens_from_ast
}

func Dummy() {
	fmt.Println("dummy")
}
