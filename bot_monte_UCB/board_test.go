package main

import(
	"testing"
//	"fmt"
	"sort"
	"math/rand"
	"time"
)

func Test_SubBoard_GetWinner(t *testing.T){
	//---------t1---------
	var subboard SubBoard = SubBoard{
		0,0,0,
		0,0,0,
		0,0,0}
	
	if subboard.GetWinner() != 0{
		t.Error("FAIL: Blank")
	}

	//---------t2---------
	subboard = SubBoard{
		0,2,0,
		0,0,0,
		0,0,0}
	
	if subboard.GetWinner() != 0{
		t.Error("FAIL: Single")
	}

	//---------t3---------
	subboard = SubBoard{
		0,2,0,
		0,2,0,
		0,2,0}
	
	if subboard.GetWinner() != 2{
		t.Error("FAIL: Vertical")
	}

	//---------t4---------
	subboard = SubBoard{
		0,0,0,
		2,2,2,
		0,0,0}
	
	if subboard.GetWinner() != 2{
		t.Error("FAIL: Horizontal")
	}

	//---------t4---------
	subboard = SubBoard{
		2,0,0,
		0,2,0,
		0,0,2}
	
	if subboard.GetWinner() != 2{
		t.Error("FAIL: Diagonal")
	}
}

func Test_UltimateBoard_GetWinner(t *testing.T){
	//---------t1---------
	var b1 SubBoard = SubBoard{0,2,0,0,2,0,0,2,0} //Winner
	var b2 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b3 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b4 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b5 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b6 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b7 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b8 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b9 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}

	var ultimateBoard UltimateBoard = UltimateBoard{b1,b2,b3,b4,b5,b6,b7,b8,b9}

	//fmt.Println(ultimateBoard.GetWinner())
	if ultimateBoard.GetWinner() != 0{
		t.Error("FAIL NONE: (expected) 0 ==", ultimateBoard.GetWinner(), "(actual)")
	}

	//---------t2---------
	b1 = SubBoard{0,2,0,0,2,0,0,2,0} //Winner
	b2 = SubBoard{0,2,0,0,2,0,0,2,0} //Winner
	b3 = SubBoard{0,2,0,0,2,0,0,2,0} //Winner
	b4 = SubBoard{0,0,0,0,0,0,0,0,0}
	b5 = SubBoard{0,0,0,0,0,0,0,0,0}
	b6 = SubBoard{0,0,0,0,0,0,0,0,0}
	b7 = SubBoard{0,0,0,0,0,0,0,0,0}
	b8 = SubBoard{0,0,0,0,0,0,0,0,0}
	b9 = SubBoard{0,0,0,0,0,0,0,0,0}

	ultimateBoard = UltimateBoard{b1,b2,b3,b4,b5,b6,b7,b8,b9}

	//fmt.Println(ultimateBoard.GetWinner())
	if ultimateBoard.GetWinner() != 2{
		t.Error("FAIL NONE: (expected) 2 ==", ultimateBoard.GetWinner(), "(actual)")
	}

	//---------t3---------
	b1 = SubBoard{0,2,0,0,2,0,0,2,0} //Winner
	b2 = SubBoard{0,0,0,0,0,0,0,0,0}
	b3 = SubBoard{0,0,0,0,0,0,0,0,0}
	b4 = SubBoard{0,0,0,0,0,0,0,0,0}
	b5 = SubBoard{0,2,0,0,2,0,0,2,0} //Winner
	b6 = SubBoard{0,0,0,0,0,0,0,0,0}
	b7 = SubBoard{0,0,0,0,0,0,0,0,0}
	b8 = SubBoard{0,0,0,0,0,0,0,0,0}
	b9 = SubBoard{0,2,0,0,2,0,0,2,0} //Winner

	ultimateBoard = UltimateBoard{b1,b2,b3,b4,b5,b6,b7,b8,b9}

	//fmt.Println(ultimateBoard.GetWinner())
	if ultimateBoard.GetWinner() != 2{
		t.Error("FAIL NONE: (expected) 2 ==", ultimateBoard.GetWinner(), "(actual)")
	}

	//---------t4---------
	b1 = SubBoard{0,2,0,0,2,0,0,2,0} //Winner
	b2 = SubBoard{0,0,0,0,0,0,0,0,0}
	b3 = SubBoard{0,0,0,0,0,0,0,0,0}
	b4 = SubBoard{0,0,0,0,0,0,0,0,0}
	b5 = SubBoard{0,1,0,0,1,0,0,1,0} //Wrong Winner
	b6 = SubBoard{0,0,0,0,0,0,0,0,0}
	b7 = SubBoard{0,0,0,0,0,0,0,0,0}
	b8 = SubBoard{0,0,0,0,0,0,0,0,0}
	b9 = SubBoard{0,2,0,0,2,0,0,2,0} //Winner

	ultimateBoard = UltimateBoard{b1,b2,b3,b4,b5,b6,b7,b8,b9}

	//fmt.Println(ultimateBoard.GetWinner())
	if ultimateBoard.GetWinner() != 0{
		t.Error("FAIL NONE: (expected) 0 ==", ultimateBoard.GetWinner(), "(actual)")
	}
}

func Test_SubBoard_GetValidMoves(t *testing.T){

	//---------t1---------
	boardIndex := 1

	b1 := SubBoard{0,0,0,0,0,0,0,0,0} //Blank
	moves := b1.GetValidMoves(boardIndex)	
	// for _, ele := range moves{
	// 	fmt.Println(ele.board, ele.square)
	// }
	if moves[4].board != boardIndex &&
		moves[4].square != 4{
		t.Error("ERROR 1")
	}

	//---------t1---------
	boardIndex = 1

	b1 = SubBoard{1,1,1,1,1,1,1,1,1} //Blank
	moves = b1.GetValidMoves(boardIndex)	
	// for _, ele := range moves{
	// 	fmt.Println(ele.board, ele.square)
	// }
	if len(moves) != 0{
		t.Error("ERROR 2")
	}
}

func Test_Ultimate_GetValidMoves(t *testing.T){
	//---------t1---------
	var b1 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b2 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b3 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b4 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b5 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b6 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b7 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b8 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b9 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}

	var ultimateBoard UltimateBoard = UltimateBoard{b1,b2,b3,b4,b5,b6,b7,b8,b9}

	lastMove := &Move{0, 0}
	moves := ultimateBoard.GetValidMoves(lastMove)
	
	// for _, ele := range moves{
	// 	fmt.Println(ele.board, ele.square)
	// }

	if len(moves) != 9{
		t.Error("ERROR 1")
	}

	//---------t2---------
	b1 = SubBoard{2,0,0,2,0,0,2,0,0}
	b2 = SubBoard{0,0,0,0,0,0,0,0,0}
	b3 = SubBoard{0,0,0,0,0,0,0,0,0}
	b4 = SubBoard{0,0,0,0,0,0,0,0,0}
	b5 = SubBoard{0,0,0,0,0,0,0,0,0}
	b6 = SubBoard{0,0,0,0,0,0,0,0,0}
	b7 = SubBoard{0,0,0,0,0,0,0,0,0}
	b8 = SubBoard{0,0,0,0,0,0,0,0,0}
	b9 = SubBoard{0,0,0,0,0,0,0,0,0}

	ultimateBoard = UltimateBoard{b1,b2,b3,b4,b5,b6,b7,b8,b9}

	lastMove = &Move{0, 0}
	moves = ultimateBoard.GetValidMoves(lastMove)
	
	// for _, ele := range moves{
	// 	fmt.Println(ele.board, ele.square)
	// }

	if len(moves) != 8*9{
		t.Error("ERROR 2")
	}

	//---------t3---------
	b1 = SubBoard{2,0,0,2,0,0,2,0,0}
	b2 = SubBoard{2,0,0,2,0,0,2,0,0}
	b3 = SubBoard{2,0,0,2,0,0,2,0,0}
	b4 = SubBoard{2,0,0,2,0,0,2,0,0}
	b5 = SubBoard{2,0,0,2,0,0,2,0,0}
	b6 = SubBoard{2,0,0,2,0,0,2,0,0}
	b7 = SubBoard{2,0,0,2,0,0,2,0,0}
	b8 = SubBoard{2,0,0,2,0,0,2,0,0}
	b9 = SubBoard{2,0,0,2,0,0,2,0,0}

	ultimateBoard = UltimateBoard{b1,b2,b3,b4,b5,b6,b7,b8,b9}

	lastMove = &Move{0, 0}
	moves = ultimateBoard.GetValidMoves(lastMove)
	
	// for _, ele := range moves{
	// 	fmt.Println(ele.board, ele.square)
	// }

	if len(moves) != 0{
		t.Error("ERROR 3")
	}
}

func Test_ApplyMove(t *testing.T){
	move := &Move{0,0}
	playerId := 1

	var b1 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b2 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b3 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b4 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b5 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b6 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b7 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b8 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}
	var b9 SubBoard = SubBoard{0,0,0,0,0,0,0,0,0}


	var ultimateBoard UltimateBoard = UltimateBoard{b1,b2,b3,b4,b5,b6,b7,b8,b9}
	
	ultimateBoard.ApplyMove(move, playerId)
	//fmt.Println(ultimateBoard[0][0])
	if ultimateBoard[0][0] != 1 {
		t.Error("Error 1")
	}
	
}

//TEMP

type Tuple struct{
	move *Move
	chance float64
}

type Tuples []Tuple

func (slice Tuples) Len() int {
    return len(slice)
}

func (slice Tuples) Less(i, j int) bool {
    return slice[i].chance > slice[j].chance;
}

func (slice Tuples) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}

func Test_RandomMoveRandomizer(t *testing.T){
	rand.Seed(time.Now().UTC().UnixNano())

	chances := make([]float64, 9)
	chances[0] = 3/8.0
	chances[1] = 2/8.0
	chances[2] = 3/8.0
	chances[3] = 2/8.0
	chances[4] = 4/8.0
	chances[5] = 2/8.0
	chances[6] = 3/8.0
	chances[7] = 2/8.0
	chances[8] = 3/8.0

	m1 := &Move{1,1}  // 2/8 * 2/8
	m2 := &Move{0,1}  // 3/8 * 2/8
	m3 := &Move{0,0}  // 3/8 * 3/8
	m4 := &Move{4,4}  // 4/8 * 4/8

	var validMoves []*Move = []*Move{m4, m3, m2, m1}
	tuples := make(Tuples, len(validMoves))

	for i, move := range validMoves{
		tuples[i] = Tuple{move, chances[move.board] * chances[move.square]}
		//fmt.Println(tuples[i].move, tuples[i].chance)
	}


	sort.Sort(tuples)

	// fmt.Println("\nSorted")
    // for i, c := range tuples {
    //     fmt.Println(i, c.move, c.chance)
    // }

	//fmt.Println("len(validMoves) =", len(validMoves))
	
	length := 10//len(validMoves)
	count := make([]int, length)
	for i := 0; i < 10000; i++{
		rand1 := rand.Float64() 
		rand2 := rand.Float64()		
		randIndex := int(rand1 * rand2 * float64(length))
		count[randIndex]++
		//fmt.Println(randIndex)
	}
	//fmt.Println(count)

	//validMoves = append(validMoves, m1, m2, m3, m4)
}
