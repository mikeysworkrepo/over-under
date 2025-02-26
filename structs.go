package main

type Track struct {
	name   string
	bet    int
	choice string
	payout int
}

type Game struct {
	PlayerID string // discord user ID
	Name     string // custom name ( mikey )
	Bet      int    // the amount you bet
	Choice   string // the bet you chose "over", "under", or "7"
	Dice1    int    // first dice roll - will have print to discord channel
	Dice2    int    // second dice roll - will hav print to discord channel
	Sum      int    // seperate print of sum of two dice rolls - will print to discord channel
	Payout   int    // payout (if there is one) - prints to discord channel
	Won      bool   // win or loss - for future leaderboard functionality
}
