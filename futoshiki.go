// A futoshiki is a square made by n x n numbers.
// Possible numbers go from 1 to n.
// For a number to be valid, the same number must not
// appear in the same row and column.
//
// Configurable traits:
// - size of the square
// - number of numbers to display
// - number of greater than signs to display

package futoshiki

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Cell struct {
	X int
	Y int
}

type Futoshiki struct {
	Square [][]int
}

func (f *Futoshiki) Print() {
	for row := range f.Square {
		rowNumbers := f.Square[row]
		rowString := ""
		var betweenRowsSymbols []string
		for column := range rowNumbers {
			// skip first column
			str := ""
			if column > 0 {
				symbol := ">"
				if rowNumbers[column-1] < rowNumbers[column] {
					symbol = "<"
				}
				str += " " + symbol + " "
			}

			rowString += fmt.Sprintf("%s%v", str, rowNumbers[column])

			// skip last row
			if row != len(f.Square)-1 {
				symbol := "⋀"
				if f.Square[row][column] > f.Square[row+1][column] {
					symbol = "⋁"
				}
				betweenRowsSymbols = append(betweenRowsSymbols, symbol)
			}

		}

		fmt.Println(rowString)
		fmt.Println(strings.Join(betweenRowsSymbols, "   "))
	}
}

func GenerateFutoshiki(size int) *Futoshiki {
	rand.Seed(time.Now().UnixNano())

	sq := make([][]int, size)
	for i := 0; i < size; i++ {
		sq[i] = make([]int, size)
	}

	availableNumbers := generateSequence(size)

	for row := range sq {
		for column := range sq[row] {
			cell := Cell{
				X: column,
				Y: row,
			}
			numbers := availableNumbersInCell(cell, sq, availableNumbers)

			if len(numbers) == 0 {
				sq = make([][]int, 0)
				return GenerateFutoshiki(size)
			}

			number := numbers[rand.Intn(len(numbers))]

			sq[row][column] = number
		}
	}

	return &Futoshiki{
		Square: sq,
	}
}

func availableNumbersInCell(cell Cell, square [][]int, overallAvailableNumbers []int) []int {
	var availableNumbers []int

	canUseNumber := map[int]bool{}
	for _, n := range overallAvailableNumbers {
		canUseNumber[n] = true
	}

	for _, n := range numbersInRow(cell, square) {
		canUseNumber[n] = false
	}

	for _, n := range numbersInColumn(cell, square) {
		canUseNumber[n] = false
	}

	for number, canUse := range canUseNumber {
		if canUse {
			availableNumbers = append(availableNumbers, number)
		}
	}

	return availableNumbers
}

func numbersInRow(cell Cell, square [][]int) []int {
	return square[cell.Y]
}

func numbersInColumn(cell Cell, square [][]int) []int {
	var column []int

	for _, row := range square {
		column = append(column, row[cell.X])
	}

	return column
}

// generates a sequence of numbers from 1 to max
func generateSequence(max int) []int {
	var sequence []int
	for i := 1; i <= max; i++ {
		sequence = append(sequence, i)
	}

	return sequence
}
