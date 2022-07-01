package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {

	err := keyboard.Open()

	// bail out of program if we cannot open keyboard package.
	if err != nil {
		log.Fatal(err)
	}

	// create a "anonymous" function. Defer makes not execute immediately... whatever follows is delayed
	defer func() {
		_ = keyboard.Close()
	}()

	//create a map
	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program.")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))

		fmt.Println(fmt.Sprintf("You chose %v", coffees[i]))

	}

	fmt.Println("Program exiting...")
}
