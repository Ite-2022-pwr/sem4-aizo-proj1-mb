package submenus

import "fmt"

func DisplayList(list []any) {
	fmt.Print("Lista: [ ")
	for _, v := range list {
		fmt.Print(v, ", ")
	}
	fmt.Print("]")
	fmt.Println()
}
