package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(propmt string, r *bufio.Reader) (string, error) {
	fmt.Print(propmt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)

	fmt.Printf("Created the bill - %v \n\n", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a -add item, s -save bill, t -add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price ($): ", reader)

		p, err := strconv.ParseFloat(price, 32)
		if err != nil {
			fmt.Println("the price must be a number")
			promptOptions(b)
		}

		b.addItem(name, float32(p))

		fmt.Println("item added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 32)
		if err != nil {
			fmt.Println("the tip must be a number")
			promptOptions(b)
		}

		b.updateTip(float32(t))

		fmt.Println("tip updated - ", tip)
		promptOptions(b)
	case "s":
		fmt.Println("you chose to save the bill", b)
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}
