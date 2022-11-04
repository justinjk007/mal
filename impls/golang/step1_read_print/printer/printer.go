package printer

import (
	Types "mal/types"
	"container/list"
	"fmt"
)

func Pr_str(list *list.List) string {

	// symbol: return the string name of the symbol
	// number: return the number as a string
	// list: iterate through each element of the list calling pr_str
	// on it, then join the results with a space separator, and
	// surround the final result with parens

	for l := list.Front(); l != nil; l = l.Next() {
		var token Types.MalType
		token = l.Value.(Types.MalType)
		switch token.DataType {
		case "string":
			return token.StringVal
		case "int":
			return token.StringVal
		case "symbol":
			return token.StringVal
		default:
			fmt.Println("Unexpected type has been send to print")
			return "nil"
		}

	}

	return "nil"
}
