package reader

import (
	"container/list"
	"fmt"
	Types "mal/types"
	"regexp"
	"unicode"
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
func Read_form(reader Read) Types.MalType {
	token := reader.Peek()
	Typelist := list.New()
	mal := Types.MalType{"null","anytype"}
	if token == "(" {
		Typelist = Read_list(reader)
	} else {
		mal = Read_atom(reader)
	}
	return mal
}

// Add the function read_list to reader.qx. This function will repeatedly call
// read_form with the Reader object until it encounters a ')' token (if it reach
// EOF before reading a ')' then that is an error). It accumulates the results
// into a List type. If your language does not have a sequential data type that
// can hold mal type values you may need to implement one (in types.qx). Note
// that read_list repeatedly calls read_form rather than read_atom. This
// mutually recursive definition between read_list and read_form is what allows
// lists to contain lists.
func Read_list(reader Read) List {
	list := list.New()
	for reader.Next() != ")" {
		token := Read_form(reader) // Maltype
		list.PushBack(token)
	}
	return list
}

// This function will look at the contents of the token and return the
// appropriate scalar (simple/single) data type value.
func Read_atom(reader Read) Types.MalType {
	mal := Types.MalType{"null","anytype"}
	token := reader.Peek()
	trimmed_token := strings.TrimSpace(token)
	// For now only thinking about numbers ans symbols so everything will be
	// 1 char length, which can be converted to rune which
	rune_token := strconv.FormatUint(uint64([]rune(trimmed_token)[0]), 16)
	if unicode.IsNumber(rune_token) {
		mal.Val = trimmed_token
		mal.DataType = "int"
		return mal 
	} else if unicode.IsSymbol(rune_token) {
		mal.Val = trimmed_token
		mal.DataType = "symbol"
		return mal 
	} else if unicode.IsSpace(rune_token) {
		fmt.Println("Found a space as a token, this shouldn't happen!")
		exit(1)
	} else {
		fmt.Println("Found an unexpected type!")
		exit(1)
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
