package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	guess_number := guesses()
	attempts := 0
	for {
		attempts++
		if attempts == 10 {
			fmt.Printf("You couldn't get it right in 10 attempts, the number id %v \n", guess_number)
			break
		}
		reader := bufio.NewReader(os.Stdin)
		userInput, _ := getInputs("Guess a number between 1 and 200: ", reader)
		i, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("input must be a number, try again")
			continue
		}
		if i < guess_number {
			fmt.Println("Your guess is too low, try again")
		} else if i > guess_number {
			fmt.Printf("Your guess is too high  \n")

		} else {
			fmt.Println("Correct!, You're must be a wizard")
			break
		}
	}
}
func getInputs(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	Input, err := r.ReadString('\n')
	return strings.TrimSpace(Input), err
}
