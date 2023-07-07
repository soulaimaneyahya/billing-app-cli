package main

import "fmt"

func main() {

	menu := map[string]float32{
		"eggs": 2,
		"bread": 3,
		"coffee": 4.99,
	}

	newBill := storeBill("breakfast", menu, 3.5)

	newBill.addItemToBill("tea", 3.99)
	newBill.updateTip(4)

	fmt.Println(newBill.format())
}
