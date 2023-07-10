package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(s string, r *bufio.Reader) (string, error) {
	fmt.Print(s)
	name, err := r.ReadString('\n')

	return strings.TrimSpace(name), err
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

	opt, _ := getInput("Choose option (a -add item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}
