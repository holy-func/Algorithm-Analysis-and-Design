package main

import "fmt"

var fillNum int

func ChessBoard(minR, minC, specialR, specialC, size int, Board *[][]int) {
	if size == 1 {
		return
	}
	toFill := fillNum + 1
	fillNum++
	s := size >> 1
	// 左上
	if specialR < s+minR && specialC < minC+s {
		ChessBoard(minR, minC, specialR, specialC, s, Board)
	} else {
		(*Board)[minR+s-1][minC+s-1] = toFill
		ChessBoard(minR, minC, minR+s-1, minC+s-1, s, Board)
	}
	// 右上
	if specialR < s+minR && specialC >= minC+s {
		ChessBoard(minR, minC+s, specialR, specialC, s, Board)
	} else {
		(*Board)[minR+s-1][minC+s] = toFill
		ChessBoard(minR, minC+s, minR+s-1, minC+s, s, Board)
	}
	// 左下
	if specialR >= s+minR && specialC < minC+s {
		ChessBoard(minR+s, minC, specialR, specialC, s, Board)
	} else {
		(*Board)[minR+s][minC+s-1] = toFill
		ChessBoard(minR+s, minC, minR+s, minC+s-1, s, Board)
	}
	// 右下
	if specialR >= s+minR && specialC >= minC+s {
		ChessBoard(minR+s, minC+s, specialR, specialC, s, Board)
	} else {
		(*Board)[minR+s][minC+s] = toFill
		ChessBoard(minR+s, minC+s, minR+s, minC+s, s, Board)
	}

}
func main() {
	size := 8
	// size为2的次幂
	board := make([][]int, size)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
	}
	ChessBoard(0, 0, 1, 2, size, &board)
	for _, v := range board {
		fmt.Println(v)
	}
	// [3  3  4  4  8  8  9  9 ]
	// [3  2  0  4  8  7  7  9 ]
	// [5  2  2  6  10 10 7  11]
	// [5  5  6  6  1  10 11 11]
	// [13 13 14 1  1  18 19 19]
	// [13 12 14 14 18 18 17 19]
	// [15 12 12 16 20 17 17 21]
	// [15 15 16 16 20 20 21 21]
}
