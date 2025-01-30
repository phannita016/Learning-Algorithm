package main

import "fmt"

const N = 9

func solveSudoku(board [][]int) bool {
	row, col := findEmptyCell(board)
	if row == -1 && col == -1 { // หากไม่มีช่องว่างเหลือ (ไม่มีค่า 0) เท่ากับ แก้ไขเรียบร้อยแล้ว
		return true
	}

	// เมื่อมีช่องว่าง ใส่ตัวเลขจาก 1-9 ทีละตัว
	for num := 1; num <= 9; num++ {
		if isSafe(board, row, col, num) { // ทุกตัวก่อนกรอกจะต้องตรวจสอบก่อนว่าใส่ได้ไหม
			board[row][col] = num   // ถ้าสามารถกรอกได้ เติมตัวเลขลงในช่องว่าง
			if solveSudoku(board) { // เรียกใช้อีกครั้ง
				return true
			}
			board[row][col] = 0 // หากไม่สามารถกรอกได้ ให้ย้อนกลับbracktrack รีเซ็ตช่องว่างเท่ากับ 0
		}
	}

	// ถ้าลองตัวเลขทั้งหมดแล้วยังแก้ปัญหาไม่ได้ คืนค่า false
	return false
}

// ตรวจสอบช่องว่าง
func findEmptyCell(board [][]int) (int, int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 0 { // พบช่องว่างหรือเท่ากับ 0
				return i, j // คืนค่าตำแหน่งนั้น
			}
		}
	}
	return -1, -1 // ไม่พบช่องว่างหรือไม่เท่ากับ 0
}

// การตรวจตัวเลข สามารถกรอกลงได้หรือไม่
func isSafe(board [][]int, row, col, num int) bool {
	for i := 0; i < N; i++ {
		if board[row][i] == num || board[i][col] == num { // ตรวจสอบว่า row,col มีตัวเลขซ้ำหรือไม่
			return false
		}
	}

	// ตรวจสอบข้อมูล 3*3 ช่อง มี num ซ้ำไหม  หาจุดเริ่มต้นของข้อมูลจาก row-row%3 และ col-col%3 แล้ววนผ่านข้อมูลนั้น
	startRow, startCol := row-row%3, col-col%3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}

	// ตรวจสอบแล้วไม่มี num ซ้ำเลย
	return true
}

func printBoard(board [][]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
}

func main() {
	board := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	if solveSudoku(board) {
		printBoard(board)
	} else {
		fmt.Println("No solution exists")
	}
}
