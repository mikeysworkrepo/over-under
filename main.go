package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to over under dice game!")
	fmt.Println("To place a bet, type out in this format into the channel:\n!bet name amount call (over, under or 7)")
	fmt.Println("Payouts for winning on over/under call 1:1 - Payouts for winning on 7 call 3:1")
	fmt.Println("Example: !bet mikey 500 under")

	// valid bet choices
	validCalls := map[string]bool{"over": true, "under": true, "7": true}

	reader := bufio.NewReader(os.Stdin)

	for {
		// asks for input
		fmt.Println("\nEnter your bet: (!bet name amount call) ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// splits the bet input into 4 words
		parts := strings.Fields(input)

		if len(parts) != 4 || parts[0] != "!bet" {
			fmt.Println("ah ah ah, invalid command. Command must be in format of: !bet name amount call (over, under or 7)")
			continue
		}

		// extract the input values
		name := parts[1]                  // name field
		amountStr := parts[2]             //bet amount, will convert to int later
		call := strings.ToLower(parts[3]) // convert to lowercase

		betAmount, err := strconv.Atoi(amountStr)
		if err != nil || betAmount <= 0 {
			fmt.Println("ah ah ah, invalid bet amount. Must be a number over 0")
			continue
		}

		// validate the call type (over, under or 7)

		if !validCalls[call] {
			fmt.Println("ah ah ah, you didnt use the magic word for your call. Use over under or 7")
			continue
		}

		// if all checks pass print the bet with the amount as an int
		fmt.Printf("Bet place!: %s bets %d on %s! Good luck!\n", name, betAmount, call)
		Roll(call, betAmount)
		break
	}

}
