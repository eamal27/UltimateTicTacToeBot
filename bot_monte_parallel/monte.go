package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"sync"
	"time"
)

var TIME_PER_MOVE float64 = 500.0 //Time Per Move in Milliseconds

var mutex *sync.Mutex
var wins map[Move]float64
var losses map[Move]float64
var ties map[Move]float64
var weightedWins map[Move]float64
var weightedLosses map[Move]float64
var gamesPlayed int

func RunMonteCarlo(validBoards []int, board *UltimateBoard) *Move {
	startTime := time.Now()
	var movesToTry []*Move

	for _, validBoard := range validBoards {
		boardMoves := board[validBoard].GetValidMoves(validBoard)
		movesToTry = append(movesToTry, boardMoves...)
	}

	mutex = &sync.Mutex{}
	wins = make(map[Move]float64)
	losses = make(map[Move]float64)
	ties = make(map[Move]float64)
	weightedWins = make(map[Move]float64)
	weightedLosses = make(map[Move]float64)
	gamesPlayed = 0

	allValidMoves := make([]Move, len(movesToTry))
	for i := 0; i < len(movesToTry); i++ {
		allValidMoves[i] = *movesToTry[i]
	}

	var wg sync.WaitGroup
	for {
		for _, move := range movesToTry {
			gamesPlayed += 1
			simulatedBoard := board.Copy()

			//Run parllel
			wg.Add(1)
			go runParallel(simulatedBoard, move, &wg)
		}

		end := time.Now()
		if end.Sub(startTime).Seconds()*1000.0 > TIME_PER_MOVE {
			break
		}
	}

	bestScore := math.Inf(-1)
	var bestMove Move

	//bebug: False when actually running the bot on the engine
	debug := false
	if debug {
		fmt.Println("gamesPlayed =", gamesPlayed)
	}

	wg.Wait()
	for _, move := range allValidMoves {
		// score := UCB1(wins[move], wins[move]+losses[move]+ties[move], gamesPlayed) //UCB1
		score := (weightedWins[move] - (weightedLosses[move] * 4.0)) / (wins[move] + losses[move] + ties[move]) //Original
		if debug {
			fmt.Printf("%v: [score = %.4f] [wins = %v] [losses = %v] [ties = %v] [ww = %.4f] [wl = %.4f]\n",
				move, score, wins[move], losses[move], ties[move], weightedWins[move], weightedLosses[move])
		}
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}
	if debug {
		fmt.Println("bestMove =", bestMove)
	}
	return &bestMove
}

func runParallel(simulatedBoard *UltimateBoard, move *Move, wg *sync.WaitGroup) {
	var movesToGameOver float64
	previousPlayer := PLAYER
	simulatedMove := move.Copy()
	simulatedBoard.ApplyMove(simulatedMove, PLAYER)
	movesToGameOver += 1.0

	for simulatedBoard.GetWinner() == NO_WINNER {
		if len(simulatedBoard.GetValidMoves(simulatedMove)) == 0 {
			mutex.Lock()
			ties[*move] += 1.0
			mutex.Unlock()
			break
		}

		movesToGameOver += 1.0

		simulatedMove = MoveRandomizer(simulatedMove, simulatedBoard) //Original
		//simulatedMove = BiasMoveRandomizer(simulatedMove, simulatedBoard) //Works, but slow AF
		previousPlayer = 3 - previousPlayer
		simulatedBoard.ApplyMove(simulatedMove, previousPlayer)
	}
	simulatedWinner := simulatedBoard.GetWinner()
	mutex.Lock()
	if simulatedWinner == PLAYER {
		wins[*move] += 1.0
		weightedWins[*move] += (1.0 / movesToGameOver)
	} else if simulatedWinner == OPPONENT {
		losses[*move] += 1.0
		weightedLosses[*move] += (1.0 / movesToGameOver)
	}
	mutex.Unlock()
	wg.Done()
}

//Upper Confidence Bound Function
func UCB1(wi float64, ni float64, t int) float64 {
	// 	wi stands for the number of wins after the i-th move;
	// ni stands for the number of simulations after the i-th move;
	// c is the exploration parameter—theoretically equal to √2; in practice usually chosen empirically;
	// t stands for the total number of simulations for the node considered. It is equal to the sum of all the ni.
	var C float64 = math.Sqrt(2)
	exploit := wi / ni
	explore := C * math.Sqrt(math.Log(float64(t))/ni)
	score := exploit + explore
	return score
}
func MoveRandomizer(lastMove *Move, board *UltimateBoard) *Move {
	validMoves := board.GetValidMoves(lastMove)
	randomMoveIndex := rand.Intn(len(validMoves))
	return validMoves[randomMoveIndex]
}

//----------------TESTING ----------------
type Tuple struct {
	move   *Move
	chance float64
}

type Tuples []Tuple

func (slice Tuples) Len() int {
	return len(slice)
}

func (slice Tuples) Less(i, j int) bool {
	return slice[i].chance > slice[j].chance
}

func (slice Tuples) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

//Selects move based on the different probabilistic chances of the board
func BiasMoveRandomizer(lastMove *Move, board *UltimateBoard) *Move {
	rand.Seed(time.Now().UTC().UnixNano())

	chances := make([]float64, 9)
	chances[0] = 3 / 8.0
	chances[1] = 2 / 8.0
	chances[2] = 3 / 8.0
	chances[3] = 2 / 8.0
	chances[4] = 4 / 8.0
	chances[5] = 2 / 8.0
	chances[6] = 3 / 8.0
	chances[7] = 2 / 8.0
	chances[8] = 3 / 8.0

	//Get the valid moves from the current state of the board
	validMoves := board.GetValidMoves(lastMove)
	length := len(validMoves)

	//Create tuples and set them to the same Moves with their matching chances
	tuples := make(Tuples, length)
	for i, move := range validMoves {
		tuples[i] = Tuple{move, chances[move.board] * chances[move.square]}
	}

	//Sort them so that 0 index is the highest
	sort.Sort(tuples)

	//Generate a random index, with a higher distribution towards the 0th index
	rand1 := rand.Float64()
	rand2 := rand.Float64()
	randIndex := int(rand1 * rand2 * float64(length))

	//Finally return the randomly generated index
	return validMoves[randIndex]
}
