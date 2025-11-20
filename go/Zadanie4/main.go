package main

import (
	"fmt"
)

func main() {
	sequences := InitArray(5)
	patterns := InitArray(5)

	fmt.Print("Введите строку: ")
	var s string
	fmt.Scan(&s)
	sequences.Add(s)

	fmt.Print("Введите шаблон: ")
	var patternStr string
	fmt.Scan(&patternStr)

	for i := 0; i < sequences.Length(); i++ {
		current := sequences.Get(i)
		fmt.Printf("%s - ", current)
		if MatchPattern(current, patternStr) {
			fmt.Println("соответствует шаблону")
		} else {
			fmt.Println("не соответствует шаблону")
		}
	}

	sequences.Destroy()
	patterns.Destroy()
}
