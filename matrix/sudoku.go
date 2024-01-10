package matrix

import "fmt"

var sudoku = map[string]bool{
	"1": false,
	"2": false,
	"3": false,
	"4": false,
	"5": false,
	"6": false,
	"7": false,
	"8": false,
	"9": false,
}

func isValidSudoku(board [][]byte) bool {
	sudokuCheck := sudoku

	for _, row := range board {
		sudokuCheck = sudoku
		for _, v := range row {
			if sudokuCheck[string(v)] {
				return false
			}
			if string(v) != "." {
				sudokuCheck[string(v)] = true
			}
		}
	}

	for i := 0; i < len(board); i++ {
		sudokuCheck = sudoku
		for j := 0; j < len(board); j++ {
			value := string(board[j][i])
			if sudokuCheck[value] {
				return false
			}
			if value != "." {
				sudokuCheck[value] = true
			}
		}
	}

	for i := 0; i < len(board); i++ {
		sudokuCheck = sudoku
		for j := i * 3; j < (i+1)*3; j++ {

			value := string(board[j][i])
			if sudokuCheck[value] {
				return false
			}
			if value != "." {
				sudokuCheck[value] = true
			}
		}
	}

	return true
}

// chatGPT
func IsValidSudokuChatGPT(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j]

				// Check row
				if rows[i][num] {
					return false
				}
				rows[i][num] = true

				// Check column
				if cols[j][num] {
					return false
				}
				cols[j][num] = true

				// Check 3x3 box
				boxIndex := (i/3)*3 + j/3
				fmt.Printf("i: %d, j: %d, boxIndex: %d \n", i, j, boxIndex)
				if boxes[boxIndex][num] {
					return false
				}
				boxes[boxIndex][num] = true
			}
		}
	}

	return true
}

// other solution
func isValidSudokuOther(board [][]byte) bool {
	row := [9]map[int]struct{}{}
	col := [9]map[int]struct{}{}
	block := [9]map[int]struct{}{}
	initMap(&row)
	initMap(&col)
	initMap(&block)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}

			number := int(board[i][j] - '0')
			blockNum := i/3*3 + j/3
			if checkExist(row[i], number) || checkExist(col[j], number) || checkExist(block[blockNum], number) {
				return false
			}
			row[i][number] = struct{}{}
			col[j][number] = struct{}{}
			block[blockNum][number] = struct{}{}
		}
	}

	return true
}

func initMap(mapOfArray *[9]map[int]struct{}) {
	for i := 0; i < len(mapOfArray); i++ {
		mapOfArray[i] = make(map[int]struct{})
	}
}

func checkExist(m map[int]struct{}, key int) bool {
	if _, exist := m[key]; exist {
		return true
	}
	return false
}
