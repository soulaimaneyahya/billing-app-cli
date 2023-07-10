package main

import "fmt"

func main() {

	newBill := newBill("breakfast")

	newBill.addItem("eggs", 2)
	newBill.addItem("bread", 3)
	newBill.addItem("coffee", 4.99)
	newBill.addItem("tea", 3.99)
	newBill.updateTip(4)

	fmt.Println(newBill.format())
}
