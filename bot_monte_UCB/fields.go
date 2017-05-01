package main

import (
	"strings"
)


var projection [9][9]string

func initFields(){
	projection[0][0] = "0 0"
	projection[0][1] = "1 0"
	projection[0][2] = "2 0"
	projection[0][3] = "0 1"
	projection[0][4] = "1 1"
	projection[0][5] = "2 1"
	projection[0][6] = "0 2"
	projection[0][7] = "1 2"
	projection[0][8] = "2 2"

	projection[1][0] = "3 0"
	projection[1][1] = "4 0"
	projection[1][2] = "5 0"
	projection[1][3] = "3 1"
	projection[1][4] = "4 1"
	projection[1][5] = "5 1"
	projection[1][6] = "3 2"
	projection[1][7] = "4 2"
	projection[1][8] = "5 2"

	projection[2][0] = "6 0"
	projection[2][1] = "7 0"
	projection[2][2] = "8 0"
	projection[2][3] = "6 1"
	projection[2][4] = "7 1"
	projection[2][5] = "8 1"
	projection[2][6] = "6 2"
	projection[2][7] = "7 2"
	projection[2][8] = "8 2"

	projection[3][0] = "0 3"
	projection[3][1] = "1 3"
	projection[3][2] = "2 3"
	projection[3][3] = "0 4"
	projection[3][4] = "1 4"
	projection[3][5] = "2 4"
	projection[3][6] = "0 5"
	projection[3][7] = "1 5"
	projection[3][8] = "2 5"

	projection[4][0] = "3 3"
	projection[4][1] = "4 3"
	projection[4][2] = "5 3"
	projection[4][3] = "3 4"
	projection[4][4] = "4 4"
	projection[4][5] = "5 4"
	projection[4][6] = "3 5"
	projection[4][7] = "4 5"
	projection[4][8] = "5 5"

	projection[5][0] = "6 3"
	projection[5][1] = "7 3"
	projection[5][2] = "8 3"
	projection[5][3] = "6 4"
	projection[5][4] = "7 4"
	projection[5][5] = "8 4"
	projection[5][6] = "6 5"
	projection[5][7] = "7 5"
	projection[5][8] = "8 5"

	projection[6][0] = "0 6"
	projection[6][1] = "1 6"
	projection[6][2] = "2 6"
	projection[6][3] = "0 7"
	projection[6][4] = "1 7"
	projection[6][5] = "2 7"
	projection[6][6] = "0 8"
	projection[6][7] = "1 8"
	projection[6][8] = "2 8"

	projection[7][0] = "3 6"
	projection[7][1] = "4 6"
	projection[7][2] = "5 6"
	projection[7][3] = "3 7"
	projection[7][4] = "4 7"
	projection[7][5] = "5 7"
	projection[7][6] = "3 8"
	projection[7][7] = "4 8"
	projection[7][8] = "5 8"

	projection[8][0] = "6 6"
	projection[8][1] = "7 6"
	projection[8][2] = "8 6"
	projection[8][3] = "6 7"
	projection[8][4] = "7 7"
	projection[8][5] = "8 7"
	projection[8][6] = "6 8"
	projection[8][7] = "7 8"
	projection[8][8] = "8 8"
}

func projectFields(fieldString string) [9][9]int {
	fs := strings.Split(fieldString, ",")
	fields := [9][9]int{}

	fields[0][0] = getInt(fs[0])
	fields[0][1] = getInt(fs[1])
	fields[0][2] = getInt(fs[2])
	fields[0][3] = getInt(fs[9])
	fields[0][4] = getInt(fs[10])
	fields[0][5] = getInt(fs[11])
	fields[0][6] = getInt(fs[18])
	fields[0][7] = getInt(fs[19])
	fields[0][8] = getInt(fs[20])

	fields[1][0] = getInt(fs[3])
	fields[1][1] = getInt(fs[4])
	fields[1][2] = getInt(fs[5])
	fields[1][3] = getInt(fs[12])
	fields[1][4] = getInt(fs[13])
	fields[1][5] = getInt(fs[14])
	fields[1][6] = getInt(fs[21])
	fields[1][7] = getInt(fs[22])
	fields[1][8] = getInt(fs[23])

	fields[2][0] = getInt(fs[6])
	fields[2][1] = getInt(fs[7])
	fields[2][2] = getInt(fs[8])
	fields[2][3] = getInt(fs[15])
	fields[2][4] = getInt(fs[16])
	fields[2][5] = getInt(fs[17])
	fields[2][6] = getInt(fs[24])
	fields[2][7] = getInt(fs[25])
	fields[2][8] = getInt(fs[26])

	fields[3][0] = getInt(fs[27])
	fields[3][1] = getInt(fs[28])
	fields[3][2] = getInt(fs[29])
	fields[3][3] = getInt(fs[36])
	fields[3][4] = getInt(fs[37])
	fields[3][5] = getInt(fs[38])
	fields[3][6] = getInt(fs[45])
	fields[3][7] = getInt(fs[46])
	fields[3][8] = getInt(fs[47])

	fields[4][0] = getInt(fs[30])
	fields[4][1] = getInt(fs[31])
	fields[4][2] = getInt(fs[32])
	fields[4][3] = getInt(fs[39])
	fields[4][4] = getInt(fs[40])
	fields[4][5] = getInt(fs[41])
	fields[4][6] = getInt(fs[48])
	fields[4][7] = getInt(fs[49])
	fields[4][8] = getInt(fs[50])

	fields[5][0] = getInt(fs[33])
	fields[5][1] = getInt(fs[34])
	fields[5][2] = getInt(fs[35])
	fields[5][3] = getInt(fs[42])
	fields[5][4] = getInt(fs[43])
	fields[5][5] = getInt(fs[44])
	fields[5][6] = getInt(fs[51])
	fields[5][7] = getInt(fs[52])
	fields[5][8] = getInt(fs[53])

	fields[6][0] = getInt(fs[54])
	fields[6][1] = getInt(fs[55])
	fields[6][2] = getInt(fs[56])
	fields[6][3] = getInt(fs[63])
	fields[6][4] = getInt(fs[64])
	fields[6][5] = getInt(fs[65])
	fields[6][6] = getInt(fs[72])
	fields[6][7] = getInt(fs[73])
	fields[6][8] = getInt(fs[74])

	fields[7][0] = getInt(fs[57])
	fields[7][1] = getInt(fs[58])
	fields[7][2] = getInt(fs[59])
	fields[7][3] = getInt(fs[66])
	fields[7][4] = getInt(fs[67])
	fields[7][5] = getInt(fs[68])
	fields[7][6] = getInt(fs[75])
	fields[7][7] = getInt(fs[76])
	fields[7][8] = getInt(fs[77])

	fields[8][0] = getInt(fs[60])
	fields[8][1] = getInt(fs[61])
	fields[8][2] = getInt(fs[62])
	fields[8][3] = getInt(fs[69])
	fields[8][4] = getInt(fs[70])
	fields[8][5] = getInt(fs[71])
	fields[8][6] = getInt(fs[78])
	fields[8][7] = getInt(fs[79])
	fields[8][8] = getInt(fs[80])

	return fields
}

func projectMove(macro_index int, micro_index int) string {
	r := "place_move"
	r = r + " " + projection[macro_index][micro_index]
	return r
}
