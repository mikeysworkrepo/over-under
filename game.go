package main

import (
	"fmt"
	"math/rand"
)

func Roll(call string, betAmount int) string {
	var dice1 = []int{1, 2, 3, 4, 5, 6}
	var dice2 = []int{1, 2, 3, 4, 5, 6}
	n := rand.Intn(6)
	i := rand.Intn(6)
	result := dice1[n] + dice2[i]

	fmt.Println("Rolling dice...You have rolled a ", dice1[n], "\nRolling dice...You have rolled a ", dice2[i])
	fmt.Println("Your dice total is", result)
	if call == "under" && result <= 6 {
		fmt.Printf("You have won! Your payout is %d!", betAmount)
	} else if call == "under" && result >= 7 {
		fmt.Printf("You have have lost! Your loss is %d", betAmount)
	} else if call == "over" && result >= 8 {
		fmt.Printf("You have won! Your payout is %d!", betAmount)
	} else if call == "over" && result <= 7 {
		fmt.Printf("You have have lost! Your loss is %d", betAmount)
	} else if call == "7" && result == 7 {
		fmt.Printf("You have won! Your payout is %d!", betAmount)
	}
	return "You lose"

}
