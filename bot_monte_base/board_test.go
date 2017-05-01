package main

import(
	"testing"
	"fmt"
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
	fmt.Println(ultimateBoard[0][0])
	if ultimateBoard[0][0] != 1 {
		t.Error("Error 1")
	}
	
}


