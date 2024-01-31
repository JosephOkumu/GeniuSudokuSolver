package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	//check if the number of arguments is not equal to nine
	if len(args) != 9 {
		fmt.Println("Error")
		return
	}

	//Creates the board of size 9 * 9
	var board [9][9]int

	//parsing command line arguments and populating the sudoku board
	for i := 0; i < 9; i++ {
		//check if the number of characters in the row is not equal to 9
		if len(args[i]) != 9 {
			fmt.Println("Error")

			return
		}

		for j := 0; j < 9; j++ {
			//check if the character is a dot (.) or a digit
			if args[i][j] == '.' {
				board[i][j] = 0
			} else {
				//convert the character to an integer, handle errors
				num, err := strconv.Atoi(string(args[i][j]))
				if err != nil {
					fmt.Println("Error")
					return
				}
				board[i][j] = num

			}

		}
	}
	//Attempt to solve the sudoku board
	if solveSudoku(&board) {
		printBoard(&board)
	} else {
		fmt.Println("No solution found for the puzzle!")
	}
}

func findEmptyCell(board *[9][9]int) *[2]int {
	//Find the first empty cell with a value of zero
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				return &[2]int{row, col}
			}
		}
	}
	return nil // return nil if no empty cell is found

}

func isValidNumRow(board *[9][9]int, num, row int) bool {
	//check if the number already exists in the current row
	for col := 0; col < 9; col++ {
		if board[row][col] == num {
			return false
		}
	}
	return true
}

func isValidNumCol(board *[9][9]int, num, col int) bool {
	//check if the number exists in the current column
	for row := 0; row < 9; row++ {
		if board[row][col] == num {
			return false
		}
	}
	return true
}

func isValidNumBox(board *[9][9]int, num, startRow, startCol int) bool {
	//check if the number already exists in the current 3X3 box
	for row := startRow; row < startRow+3; row++ {
		for col := startCol; col < startCol+3; col++ {
			if board[row][col] == num {
				return false
			}
		}

	}
	return true
}

func isValidNum(board *[9][9]int, num, row, col int) bool {
	//checks if the number exists in the current row, column and box
	return isValidNumRow(board, num, row) &&
		isValidNumCol(board, num, col) &&
		isValidNumBox(board, num, row-row%3, col-col%3)
}

func solveSudoku(board *[9][9]int) bool {
	//find the next empty cell (cell with a value of zero)
	emptyCell := findEmptyCell(board)
	if emptyCell == nil {
		return true //if there are no empty cells, the board is solved
	}

	row, col := emptyCell[0], emptyCell[1]

	//Try numbers 1 to 9 for the empty cell.

	for num := 1; num <= 9; num++ {
		//check if the number is valid in the current position

		if isValidNum(board, num, row, col) {
			board[row][col] = num

			//Recursively attempt to solve the sudoku

			if solveSudoku(board) {
				return true
			}
			board[row][col] = 0 //backtrack if no solution is found
		}

	}
	return false //if no number is valid return false

}

func printBoard(board *[9][9]int) {
	//print the solved sudoku board
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Printf("%d", board[row][col])
		}
		fmt.Println()
	}
}
