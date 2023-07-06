package main

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
