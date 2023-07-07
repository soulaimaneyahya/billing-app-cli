package main

import (
	"fmt"
	"strings"
)

type bill struct {
	name string
	menu map[string]float32
	tip float32
}


// store new bill
func storeBill(name string, menu map[string]float32, tip float32) bill {
	return bill{
		name: name,
		menu: menu,
		tip: tip,
	}
}

// string helpers
func ucwords(str string) string {
	if len(str) == 0 {
		return str
	}
	firstLetter := strings.ToUpper(string(str[0]))
	restOfStr := str[1:]
	return firstLetter + restOfStr
}

// Receiver functions
func (b bill) format() string {
	fs := "Bill breakdown:\n"
	var total float32 = 0

	// list items
	for k, v := range b.menu {
		fs += fmt.Sprintf("%-25v ...$%0.2f\n", ucwords(k)+":", v)
		total += v
	}

	// add total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}
