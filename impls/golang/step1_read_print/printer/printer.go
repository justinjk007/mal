package printer

import (
	Types "mal/types"
	"fmt"
)

func Pr_str(list Types.MalType) string {

	// symbol: return the string name of the symbol
	// number: return the number as a string
	// list: iterate through each element of the list calling pr_str
	// on it, then join the results with a space separator, and
	// surround the final result with parens

	for l := list.Front(); l != nil; l = l.Next() {
		switch l {
		case "":
			fmt.Println("This is the answer!")
		case 45:
			fmt.Println("Not the answer")
		default:
			fmt.Println("The guess is wrong!")
		}

	}

}
