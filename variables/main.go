package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "don't type your number in - just press ENTER when ready."

func main() {

	// first seed RNG so we don't get exact same number every run
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer int
	reader := bufio.NewReader(os.Stdin)

	//Display welcome/instructions to console.
	fmt.Println("Guess The Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10 ", prompt)
	reader.ReadString('\n')

	// Take user through the game
	fmt.Println("Multiply your number by ", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract ", subtraction, prompt)
	reader.ReadString('\n')

	// Give user the answer
	answer = (firstNumber * secondNumber) - subtraction
	fmt.Println("The answer is", answer)
}
