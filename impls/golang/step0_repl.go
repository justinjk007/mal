package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "azul3d.org/engine/keyboard"
)

func main() {

	fmt.Print("GO LANG MAL REPL v0.1\n")
	for true {
		fmt.Print("users> ")
		reader := bufio.NewReader(os.Stdin)
		// ReadString will block until the delimiter is entered
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input", err)
			return
		}
		// remove the delimeter from the string
		input = strings.TrimSuffix(input, "\n")
		fmt.Println(rep(input))
	} // for loop
}

func READ(input string) string {
	return input
}

func EVAL(input string) string {
	return input
}

func PRINT(input string) string {
	return input
}

func rep(input string) string {
	return PRINT(EVAL(READ(input)))
}
