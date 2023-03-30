package main

import (
	"fmt"
	"strings"
)

func main() {
	// KQRBNP
	var board [8][8]rune

	// set pieces
	const pieces = "rnbkqbnr"
	piecesCopy := strings.ToUpper(pieces)
	for i := 0; i < 8; i++ {
		board[0][i] = rune(piecesCopy[i])
		board[1][i] = 'P'
		board[6][i] = 'p'
		board[7][i] = rune(pieces[i])
	}

	// paint board

	fmt.Println("|-------|-------|-------|-------|-------|-------|-------|-------|")
	for _, row := range board {
		for _, col := range row {
			fmt.Printf("|   %c\t", col)
		}
		fmt.Print("|")
		fmt.Println()
	}
	fmt.Println("|-------|-------|-------|-------|-------|-------|-------|-------|")

}
