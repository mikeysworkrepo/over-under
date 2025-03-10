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

	// construct result message
	message := fmt.Sprintf("🎲 Rolling dice... You rolled: %d and %d\nTotal: %d\n", dice1[n], dice2[i], result)

	if call == "under" && result <= 6 {
		message += fmt.Sprintf("🎉 You won! Your payout is %d!", betAmount*2)
	} else if call == "under" && result >= 7 {
		message += fmt.Sprintf("❌ You lost! You lose %d.", betAmount)
	} else if call == "over" && result >= 8 {
		message += fmt.Sprintf("🎉 You won! Your payout is %d!", betAmount*2)
	} else if call == "over" && result <= 7 {
		message += fmt.Sprintf("❌ You lost! You lose %d.", betAmount)
	} else if call == "7" && result == 7 {
		message += fmt.Sprintf("🎉 JACKPOT! You won! Your payout is %d!", betAmount*4)
	} else {
		message += "❌ You lost!"
	}
	return message

}
