package main

import (
	"fmt"
	"log"
	"time"
)

var possibilityMatrix [9][9]map[int]bool

func main() {
	start := time.Now()
	sudoku := read()
	initPossibilityMatrix()
	print(sudoku)
	solve(sudoku)
	fmt.Println()
	print(sudoku)
	elapsed := time.Since(start)
	fmt.Printf("Execution time %s", elapsed)
}

func solve(sudoku *[9][9]int) {
	solved := checkIfSolved(sudoku)
	for !solved {
		for row := 0; row < len(sudoku); row++ {
			for col := 0; col < len(sudoku[row]); col++ {
				if sudoku[row][col] != 0 {
					continue
				}
				checkRow(sudoku, row, col)
				checkCol(sudoku, row, col)
				checkSquare(sudoku, row, col)
			}
		}
		stillSolvable := isStillSolvable(sudoku)
		if !stillSolvable {
			log.Fatalln("Sudoku is not solvable")
		}
		updateSudoku(sudoku)
		solved = checkIfSolved(sudoku)
	}
}

func updateSudoku(sudoku *[9][9]int) {
	for i := 0; i < len(sudoku); i++ {
		for j := 0; j < len(sudoku[i]); j++ {
			if sudoku[i][j] == 0 && len(possibilityMatrix[i][j]) == 1 {
				counter := 1
				for key := range possibilityMatrix[i][j] {
					if counter == 1 {
						sudoku[i][j] = key
						delete(possibilityMatrix[i][j], key)
					}
					counter++
				}
			}
		}
	}
}

func checkIfSolved(sudoku *[9][9]int) bool {
	for i := 0; i < len(sudoku); i++ {
		for j := 0; j < len(sudoku[i]); j++ {
			if sudoku[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func isStillSolvable(sudoku *[9][9]int) bool {
	for row := 0; row < len(sudoku); row++ {
		for col := 0; col < len(sudoku[row]); col++ {
			if sudoku[row][col] == 0 && len(possibilityMatrix[row][col]) == 1 {
				return true
			}
		}
	}
	return false
}

func checkRow(sudoku *[9][9]int, i, j int) {
	for number := 1; number <= 9; number++ {
		for col := 0; col < len(sudoku[i]); col++ {
			if sudoku[i][col] == number {
				delete(possibilityMatrix[i][j], number)
			}
		}
	}
}

func checkCol(sudoku *[9][9]int, i, j int) {
	for number := 1; number <= 9; number++ {
		for row := 0; row < len(sudoku); row++ {
			if sudoku[row][j] == number {
				delete(possibilityMatrix[i][j], number)
			}
		}
	}
}

func checkSquare(sudoku *[9][9]int, i, j int) {
	if i < 3 && j < 3 {
		for number := 1; number <= 9; number++ {
			for row := 0; row < 3; row++ {
				for col := 0; col < 3; col++ {
					if sudoku[row][col] == number {
						delete(possibilityMatrix[i][j], number)
					}
				}
			}
		}
	} else if i < 3 && j < 6 {
		for number := 1; number <= 9; number++ {
			for row := 0; row < 3; row++ {
				for col := 3; col < 6; col++ {
					if sudoku[row][col] == number {
						delete(possibilityMatrix[i][j], number)
					}
				}
			}
		}
	} else if i < 3 {
		for number := 1; number <= 9; number++ {
			for row := 0; row < 3; row++ {
				for col := 6; col < 9; col++ {
					if sudoku[row][col] == number {
						delete(possibilityMatrix[i][j], number)
					}
				}
			}
		}
	} else if i < 6 && j < 3 {
		for number := 1; number <= 9; number++ {
			for row := 3; row < 6; row++ {
				for col := 0; col < 3; col++ {
					if sudoku[row][col] == number {
						delete(possibilityMatrix[i][j], number)
					}
				}
			}
		}
	} else if i < 6 && j < 6 {
		for number := 1; number <= 9; number++ {
			for row := 3; row < 6; row++ {
				for col := 3; col < 6; col++ {
					if sudoku[row][col] == number {
						delete(possibilityMatrix[i][j], number)
					}
				}
			}
		}
	} else if i < 6 {
		for number := 1; number <= 9; number++ {
			for row := 3; row < 6; row++ {
				for col := 6; col < 9; col++ {
					if sudoku[row][col] == number {
						delete(possibilityMatrix[i][j], number)
					}
				}
			}
		}
	} else if j < 3 {
		for number := 1; number <= 9; number++ {
			for row := 6; row < 9; row++ {
				for col := 0; col < 3; col++ {
					if sudoku[row][col] == number {
						delete(possibilityMatrix[i][j], number)
					}
				}
			}
		}
	} else if j < 6 {
		for number := 1; number <= 9; number++ {
			for row := 6; row < 9; row++ {
				for col := 3; col < 6; col++ {
					if sudoku[row][col] == number {
						delete(possibilityMatrix[i][j], number)
					}
				}
			}
		}
	} else {
		for number := 1; number <= 9; number++ {
			for row := 6; row < 9; row++ {
				for col := 6; col < 9; col++ {
					if sudoku[row][col] == number {
						delete(possibilityMatrix[i][j], number)
					}
				}
			}
		}
	}
}

func read() *[9][9]int {
	var sudoku [9][9]int = [9][9]int{
		[9]int{0, 0, 4, 0, 5, 0, 0, 0, 0},
		[9]int{9, 0, 0, 7, 3, 4, 6, 0, 0},
		[9]int{0, 0, 3, 0, 2, 1, 0, 4, 9},
		[9]int{0, 3, 5, 0, 9, 0, 4, 8, 0},
		[9]int{0, 9, 0, 0, 0, 0, 0, 3, 0},
		[9]int{0, 7, 6, 0, 1, 0, 9, 2, 0},
		[9]int{3, 1, 0, 9, 7, 0, 2, 0, 0},
		[9]int{0, 0, 9, 1, 8, 2, 0, 0, 3},
		[9]int{0, 0, 0, 0, 6, 0, 1, 0, 0},
	}
	// content, err := os.ReadFile("sudoku.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// k := 0
	// for i := 0; i < len(sudoku); i++ {
	// 	for j := 0; j < len(sudoku[i]); j++ {
	// 		sudoku[i][j], _ = strconv.Atoi(string(content[k]))
	// 		k++
	// 	}
	// }
	return &sudoku
}

func print(sudoku *[9][9]int) {
	for i := 0; i < len(sudoku); i++ {
		for j := 0; j < len(sudoku[i]); j++ {
			fmt.Print(sudoku[i][j], " ")
		}
		fmt.Println()
	}
}

func initPossibilityMatrix() {
	for i := 0; i < len(possibilityMatrix); i++ {
		for j := 0; j < len(possibilityMatrix[i]); j++ {
			possibilityMatrix[i][j] = make(map[int]bool)
			for k := 1; k <= 9; k++ {
				possibilityMatrix[i][j][k] = true
			}
		}
		fmt.Println()
	}
}
