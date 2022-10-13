package main

import (
	"azul3d.org/engine/keyboard"
	"bufio"
	"fmt"
	"os"
	"strings"
	"mal/reader"
)

func main() {

	reader.Dummy()

	watcher := keyboard.NewWatcher()
	// Query for the map containing information about all keys
	status := watcher.States()

	var history [100]string
	history_pointer := 0

	fmt.Print("GO LANG MAL REPL v0.1\n")
	for true {

		up_key := status[keyboard.ArrowUp]
		if up_key == keyboard.Down {
			fmt.Println(history[history_pointer-1])
		}

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
		// Add the input to history
		history[history_pointer] = input
		history_pointer++
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
