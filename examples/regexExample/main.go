package main

import (
	"fmt"
	"regexp"
)

func main() {
	x := "(+ 1 2)"
	musComp := regexp.MustCompile("[\\s,]*(~@|[\\[\\]{}()'`~^@]|\"(?:\\.|[^\\\"])*\"?|;.*|[^\\s\\[\\]{}('\"`,;)]*)")
	tokens_from_ast := musComp.FindAllStringSubmatch(x, -1)

	// fmt.Println(tokens_from_ast)
	// A foreach loop in Go
	for _, element := range tokens_from_ast {
		fmt.Println(element)
	}
}
