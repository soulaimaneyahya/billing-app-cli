package main

import (
	"fmt"
	"strings"
)

// string helpers
func ucwords(str string) string {
	if len(str) == 0 {
		return str
	}
	firstLetter := strings.ToUpper(string(str[0]))
	restOfStr := str[1:]
	return firstLetter + restOfStr
}

// store new bill
func newBill(name string) bill {
	return bill{
		name: name,
		menu: map[string]float32{},
		tip:  0,
	}
}

// update tip
func (b *bill) updateTip(tip float32) {
	(*b).tip = tip
	// b.tip = tip
}

func (b *bill) updateName(name string) {
	(*b).name = name
	// b.name = name
}

// add item to bill
func (b *bill) addItem(name string, price float32) {
	b.menu[name] = price
}

// Receiver functions
func (b *bill) format() string {
	fs := fmt.Sprintf("Bill Breakdown (%v):\n", ucwords(b.name))

	var total float32 = 0

	// list items
	for k, v := range b.menu {
		fs += fmt.Sprintf("%-25v ...$%0.2f\n", ucwords(k)+":", v)
		total += v
	}

	// add total
	fs += fmt.Sprintf("\n%v %-15v ...$%0.2f", ucwords(b.name), "Tip:", b.tip)
	total += b.tip
	fs += fmt.Sprintf("\n%-25v ...$%0.2f", "Total:", total)

	return fs
}
