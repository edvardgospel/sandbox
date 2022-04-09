// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"
// )

// var possibilityMatrix [4][4]map[int]bool

// func main() {
// 	start := time.Now()
// 	sudoku := read()
// 	initPossibilityMatrix()
// 	print(sudoku)
// 	solve(sudoku)
// 	fmt.Println()
// 	print(sudoku)
// 	elapsed := time.Since(start)
// 	fmt.Printf("Execution time %s", elapsed)
// }

// func solve(sudoku *[4][4]int) {
// 	solved := checkIfSolved(sudoku)
// 	for !solved {
// 		for row := 0; row < len(sudoku); row++ {
// 			for col := 0; col < len(sudoku[row]); col++ {
// 				if sudoku[row][col] != 0 {
// 					continue
// 				}
// 				checkRow(sudoku, row, col)
// 				checkCol(sudoku, row, col)
// 				checkSquare(sudoku, row, col)
// 			}
// 		}
// 		stillSolvable := isStillSolvable(sudoku)
// 		if !stillSolvable {
// 			log.Fatalln("Sudoku is not solvable")
// 		}
// 		updateSudoku(sudoku)
// 		solved = checkIfSolved(sudoku)
// 	}
// }

// func updateSudoku(sudoku *[4][4]int) {
// 	for i := 0; i < len(sudoku); i++ {
// 		for j := 0; j < len(sudoku[i]); j++ {
// 			if sudoku[i][j] == 0 && len(possibilityMatrix[i][j]) == 1 {
// 				counter := 1
// 				for key := range possibilityMatrix[i][j] {
// 					if counter == 1 {
// 						sudoku[i][j] = key
// 						delete(possibilityMatrix[i][j], key)
// 					}
// 					counter++
// 				}
// 			}
// 		}
// 	}
// }

// func checkIfSolved(sudoku *[4][4]int) bool {
// 	for i := 0; i < len(sudoku); i++ {
// 		for j := 0; j < len(sudoku[i]); j++ {
// 			if sudoku[i][j] == 0 {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func isStillSolvable(sudoku *[4][4]int) bool {
// 	for row := 0; row < len(sudoku); row++ {
// 		for col := 0; col < len(sudoku[row]); col++ {
// 			if sudoku[row][col] == 0 && len(possibilityMatrix[row][col]) == 1 {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func checkRow(sudoku *[4][4]int, i, j int) {
// 	for number := 1; number <= 4; number++ {
// 		for col := 0; col < len(sudoku[i]); col++ {
// 			if sudoku[i][col] == number {
// 				delete(possibilityMatrix[i][j], number)
// 			}
// 		}
// 	}
// }

// func checkCol(sudoku *[4][4]int, i, j int) {
// 	for number := 1; number <= 4; number++ {
// 		for row := 0; row < len(sudoku); row++ {
// 			if sudoku[row][j] == number {
// 				delete(possibilityMatrix[i][j], number)
// 			}
// 		}
// 	}
// }

// func checkSquare(sudoku *[4][4]int, i, j int) {
// 	if i < 2 && j < 2 {
// 		for number := 1; number <= 4; number++ {
// 			for row := 0; row < 2; row++ {
// 				for col := 0; col < 2; col++ {
// 					if sudoku[row][col] == number {
// 						delete(possibilityMatrix[i][j], number)
// 					}
// 				}
// 			}
// 		}
// 	} else if i < 2 {
// 		for number := 1; number <= 4; number++ {
// 			for row := 0; row < 2; row++ {
// 				for col := 2; col < 4; col++ {
// 					if sudoku[row][col] == number {
// 						delete(possibilityMatrix[i][j], number)
// 					}
// 				}
// 			}
// 		}
// 	} else if j < 2 {
// 		for number := 1; number <= 4; number++ {
// 			for row := 2; row < 4; row++ {
// 				for col := 0; col < 2; col++ {
// 					if sudoku[row][col] == number {
// 						delete(possibilityMatrix[i][j], number)
// 					}
// 				}
// 			}
// 		}
// 	} else {
// 		for number := 1; number <= 4; number++ {
// 			for row := 2; row < 4; row++ {
// 				for col := 2; col < 4; col++ {
// 					if sudoku[row][col] == number {
// 						delete(possibilityMatrix[i][j], number)
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// func read() *[4][4]int {
// 	var sudoku [4][4]int = [4][4]int{
// 		[4]int{1, 0, 0, 0},
// 		[4]int{0, 0, 0, 4},
// 		[4]int{0, 0, 2, 0},
// 		[4]int{0, 3, 0, 0},
// 	}
// 	// content, err := os.ReadFile("sudoku2.txt")
// 	// fmt.Println("----------------")
// 	// fmt.Println(string(content))
// 	// fmt.Println("----------------")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// k := 0
// 	// for i := 0; i < len(sudoku); i++ {
// 	// 	for j := 0; j < len(sudoku[i]); j++ {
// 	// 		sudoku[i][j], _ = strconv.Atoi(string(content[k]))
// 	// 		k++
// 	// 	}
// 	// }
// 	return &sudoku
// }

// func print(sudoku *[4][4]int) {
// 	for i := 0; i < len(sudoku); i++ {
// 		for j := 0; j < len(sudoku[i]); j++ {
// 			fmt.Print(sudoku[i][j], " ")
// 		}
// 		fmt.Println()
// 	}
// }

// func initPossibilityMatrix() {
// 	for i := 0; i < len(possibilityMatrix); i++ {
// 		for j := 0; j < len(possibilityMatrix[i]); j++ {
// 			possibilityMatrix[i][j] = make(map[int]bool)
// 			for k := 1; k <= 4; k++ {
// 				possibilityMatrix[i][j][k] = true
// 			}
// 		}
// 		fmt.Println()
// 	}
// }
