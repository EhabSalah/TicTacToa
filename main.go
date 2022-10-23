package main

import "fmt"

const x, o = "X", "O"

func main() {
	/* an array with 5 rows and 2 columns*/

	input := [3][3]string{
		{o, x, o}, // 0,0 . 0,1 . 0,2
		{x, o, x}, // 1,0 . 1,1 . 1,2
		{x, o, x}, // 2,0 . 2,1 . 2,2
	}

	// Win probabilities:
	// 1. Horizontal win
	//
	//		{x, x, o},
	//		{o, o, o},
	//		{x, o, x},
	//
	// 2. Vertical win
	//
	// 		{x, o, o},
	//		{x, x, o},
	//		{x, o, x},
	//
	// 3. Diagonal win
	//
	//		{x, o, x},
	//		{o, x, o},
	//		{x, o, x},

	displayTheGame(input)
	result := checkWhoIsTheWinner(input)

	fmt.Printf(result)
}

func displayTheGame(input [3][3]string) {
	fmt.Printf("_____________\n")
	for _, v := range input {
		fmt.Printf("| %s | %s | %s |\n", v[0], v[1], v[2])
	}
	fmt.Printf("-------------\n")
}

func checkWhoIsTheWinner(input [3][3]string) string {

	xWinner := isThePlayerWon(input, x)
	oWinner := isThePlayerWon(input, o)

	if xWinner {
		return "X is the winner 🏆"
	}

	if oWinner {
		return "O is the winner 🏆"
	}

	return "Draw game 🏁"
}
func isThePlayerWon(input [3][3]string, thePlayer string) bool {

	var winner = false
	var i, j int

	// 1. Horizontal win check
	for i = 0; i < 3; i++ {

		for j = 0; j < 3; j++ {
			winner = input[i][j] == thePlayer

			if winner == false {
				break
			}
		}

		if winner {
			fmt.Printf("Player %s won in Horizontal case.\n", thePlayer)
			return winner
		}
	}

	// 2. Vertical win check (the opposite of the Horizontal check)
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			winner = input[j][i] == thePlayer

			if winner == false {
				break
			}
		}

		if winner {
			fmt.Printf("Player %s won on Vertical case.\n", thePlayer)

			return winner
		}
	}

	// 3. Diagonal win check (from top start to bottom end)
	for i = 0; i < 3; i++ {
		winner = input[i][i] == thePlayer

		if !winner {
			break
		}
	}

	if winner {
		fmt.Printf("Player %s won in Diagonal case (from top start to bottom end).\n", thePlayer)
		return winner
	}

	if input[0][2] == thePlayer && input[1][1] == thePlayer && input[2][0] == thePlayer {
		fmt.Printf("Player %s won in Diagonal case (from top end to bottom start).\n", thePlayer)
		return winner
	}

	return winner
}
