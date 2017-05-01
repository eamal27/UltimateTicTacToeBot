package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var player int
var opponent int
var gameState Board
var timebank int
var time_per_move int

func getInt(strVal string) int {
	val, err := strconv.Atoi(strVal)
	if err != nil {
		panic(err)
	}
	return val
}

func storeSettings(settings []string) {
	switch field := settings[0]; field {
	case "timebank":
		//Maximum time in milliseconds
		timebank = getInt(settings[1])
	case "time_per_move":
		time_per_move = getInt(settings[1])
	case "your_botid":
		player = getInt(settings[1])
		if player == 1 {
			opponent = 2
		} else {
			opponent = 1
		}
	default:
		//dont care
	}
}

func updateGame(updates []string) {
	switch update := updates[0]; update {
	case "field":
		gameState.fields = projectFields(updates[1])
	case "macroboard":
		boardStrings := strings.Split(updates[1], ",")
		for i, boardString := range boardStrings {
			//fmt.Printf("boardString = %v\n", boardString)
			gameState.macroboard[i] = getInt(boardString)
		}
	default:
		//dont care
	}
}

func getAction(params []string) string {
	action := minimax(gameState, 0, player).action
	//fmt.Printf("Player Bot action = %v\n", action);
	return action
}

var maxDepth int = 5

func minimax(board Board, depth int, currentPlayer int) Board {
	if depth >= maxDepth ||
		math.IsInf(board.score, 1) ||
		math.IsInf(board.score, -1) {
		return board
	}
	var chosenBoard Board

	otherPlayer := 3 - currentPlayer

	if currentPlayer == player {
		bestScore := math.Inf(-1)
		possibleBoards := board.GetPossibleBoards(player)
		if len(possibleBoards) == 0 {
			return board
		}
		for _, possibleBoard := range possibleBoards {
			if test := possibleBoard.score; test > bestScore {
				bestScore = possibleBoard.score
				chosenBoard = possibleBoard
			}
		}
	} else {
		worstScore := math.Inf(1)
		possibleBoards := board.GetPossibleBoards(opponent)
		if len(possibleBoards) == 0 {
			return board
		}
		for _, possibleBoard := range possibleBoards {
			if test := possibleBoard.score; test < worstScore {
				worstScore = possibleBoard.score
				chosenBoard = possibleBoard
			}
		}
	}

	return minimax(chosenBoard, depth+1, otherPlayer)
}

func testBoardScore(board [9][9]int, macroString string) {
	var testBoard Board
	testBoard.fields = board
	macroStrings := strings.Split(macroString, ",")
	for i, boardString := range macroStrings {
		testBoard.macroboard[i] = getInt(boardString)
	}
	testBoard.GetScore()
	fmt.Println("Test board score: ", testBoard.score)
}

func test(){
	var oldMacro = [9]int{-1,0,0,0,0,0,0,0,2}
	player := 1
	playSquareIndex := 8
	
	fmt.Println(oldMacro)
	newMacro := getNewMacroboard(oldMacro, player, playSquareIndex)
	fmt.Println(newMacro)
}

func main() {

	// test()
	// return

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		switch words := strings.Fields(line); words[0] {
		case "settings":
			storeSettings(words[1:])
		case "update":
			updateGame(words[2:])
		case "action":
			action := getAction(words[1:])
			fmt.Println(action)
		case "test":
			board := projectFields(words[1])
			macroString := words[2]
			testBoardScore(board, macroString)
		case "#":
			//Just ignore, it's a comment
		default:
			fmt.Println("default: " + words[0])
			//throw an error
		}
	}
}
