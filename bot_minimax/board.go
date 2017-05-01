package main

import (
//	"fmt"
	"math"
)

type Board struct {
	fields     [9][9]int
	macroboard [9]int
	score      float64
	action     string
}

var lines = [8][3]int{
	{2, 4, 6},
	{6, 7, 8},
	{3, 4, 5},
	{0, 1, 2},
	{0, 4, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8}}

func getNewMacroboard(oldMacro [9]int, player int, playSquareIndex int) [9]int{
	//fmt.Println("getNewMacroboard() called");
	//generates a new macroboard based on the old one and the square that is being
	//player by the current player

	var newMacro [9]int
	//Set the board to be the same as the oldMacro board in regard to
	//The player won squares
	for i, player := range oldMacro{
		if player != 1 && player != 2{
			newMacro[i] = 0
		}else{
			newMacro[i] = player
		}
	}

	//Set where the new player can place on the macro board (-1)
	if newMacro[playSquareIndex] == 0{
		//The square is not taken by either player, can be played on
		newMacro[playSquareIndex] = -1
	}else{
		//Allow the player to play on any square that is not taken
		for i, player := range newMacro{
			if player == 0{
				newMacro[i] = -1
			}
		}
	}

	// //Test
	// for _, ele := range newMacro{
	// 	fmt.Printf("%v ", ele)
	// }
	// fmt.Println()

	return newMacro
}

func (b *Board) GetPossibleBoards(player int) []Board {
	boards := make([]Board, 0, 9)

	for i, canPlay := range b.macroboard {
		if canPlay == -1 {
			for j, playSquare := range b.fields[i] {
				if playSquare == 0 {
					c := Board{b.fields, b.macroboard, b.score, projectMove(i, j)}
					c.GetScore()
					// newMacro := getNewMacroboard(b.macroboard, player, j);					
					// c.macroboard = newMacro;
					c.fields[i][j] = player
					boards = append(boards, c)
				}
			}
		}
	}
	return boards
}

func (b *Board) GetScore() {

	var board_scores [9]int

	for i, canPlay := range b.macroboard {
		if canPlay == -1 {
			if temp := heuristic(b.fields[i]); math.IsInf(temp, 1) {
				board_scores[i] = 1
			} else if math.IsInf(temp, -1) {
				board_scores[i] = 2
			} else {
				board_scores[i] = 0
			}
		}
	}

	b.score = heuristic(board_scores)
}

func heuristic(small_board [9]int) float64 {

	EMPTY := 0.0
	FULL_BLOCK := 0.0
	SEMI_BLOCK := 0.0
	SINGLE := 10.0
	DOUBLE := 100.0

	var score [3][3][3]float64

	score[0][0][0] = EMPTY
	score[0][0][player] = SINGLE
	score[0][0][opponent] = SINGLE * -1
	score[0][player][0] = SINGLE
	score[0][player][player] = DOUBLE
	score[0][player][opponent] = SEMI_BLOCK
	score[0][opponent][0] = SINGLE * -1
	score[0][opponent][player] = SEMI_BLOCK
	score[0][opponent][opponent] = DOUBLE * -1

	score[player][0][0] = SINGLE
	score[player][0][player] = DOUBLE
	score[player][0][opponent] = SEMI_BLOCK
	score[player][player][0] = DOUBLE
	score[player][player][player] = math.Inf(1)
	score[player][player][opponent] = FULL_BLOCK
	score[player][opponent][0] = SEMI_BLOCK
	score[player][opponent][player] = FULL_BLOCK
	score[player][opponent][opponent] = FULL_BLOCK

	score[opponent][0][0] = SINGLE * -1
	score[opponent][0][player] = SEMI_BLOCK
	score[opponent][0][opponent] = DOUBLE * -1
	score[opponent][player][0] = SEMI_BLOCK
	score[opponent][player][player] = FULL_BLOCK
	score[opponent][player][opponent] = FULL_BLOCK
	score[opponent][opponent][0] = DOUBLE * -1
	score[opponent][opponent][player] = FULL_BLOCK
	score[opponent][opponent][opponent] = math.Inf(-1)

	var value float64 = 0
	for _, line := range lines {
		value += score[small_board[line[0]]][small_board[line[1]]][small_board[line[2]]]
	}

	return value
}
