package main

import (
	"fmt"
	"regexp"
	// "github.com/gijsbers/go-pcre"
)

func main() {
	x := "(+ 1 2 (+ \"string\" 4))"
	musComp := regexp.MustCompile("[\\s,]*(~@|[\\[\\]{}()'`~^@]|\"(?:\\.|[^\\\"])*\"?|;.*|[^\\s\\[\\]{}('\"`,;)]*)")
	tokens_from_ast := musComp.FindAllString(x, -1) // -1 means go through the entire string

	for _, attribute := range tokens_from_ast {
		fmt.Println("The token is:", attribute)
	}

}
