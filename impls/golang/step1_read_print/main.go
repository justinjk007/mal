package main

import (
	"azul3d.org/engine/keyboard"
	"bufio"
	"fmt"
	"os"
	"strings"
	"container/list"
	"mal/reader"
	"mal/printer"
)

func main() {

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
		buff_reader := bufio.NewReader(os.Stdin)
		// ReadString will block until the delimiter is entered
		input, err := buff_reader.ReadString('\n')
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

// Change the READ function in step1_read_print.qx to call
// reader.read_str and the PRINT function to call printer.pr_str. EVAL
// continues to simply return its input but the type is now a mal data
// type.
func READ(input string) *list.List {
	return reader.Read_str(input)
}

func EVAL(input *list.List) *list.List {
	return input
}

func PRINT(input *list.List) string {
	return printer.Pr_str(input)
}

func rep(input string) string {
	return PRINT(EVAL(READ(input)))
}
