/**
* Created by @soulaimaneyh on 2024-09-14
*/

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

func promptOptions(b bill) error {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a -add item, s -save bill, t -add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price ($): ", reader)

		p64, err := strconv.ParseFloat(price, 64)
		if err != nil {
			return fmt.Errorf("error: the price must be a number: %v", err)
		}

		p32 := float32(p64)
		b.addItem(name, p32)

		fmt.Println("item added - ", name, p32)
		return promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t64, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			return fmt.Errorf("error: the tip must be a number: %v", err)
		}

		t32 := float32(t64)
		b.updateTip(t32)

		fmt.Println("tip updated - ", t32)
		return promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you choose to save the bill: ", b.name)
	default:
		return fmt.Errorf("error: That was not a valid option...")
	}

	return nil
}

func main() {
	mybill := createBill()

	if err := promptOptions(mybill); err != nil {
		fmt.Println("Error:", err)
	}
}
