package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var forrest [99][99]int
	var visibleL [99][99]int
	var visibleR [99][99]int
	var visibleT [99][99]int
	var visibleB [99][99]int

	linenum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		for colnum, v := range line {
			forrest[linenum][colnum] = int(v - '0')

		}
		linenum += 1

	}

	for i := 0; i < 99; i++ {
		current := -1
		for j := 0; j < 99; j++ {
			if forrest[i][j] > current {
				current = forrest[i][j]
				visibleL[i][j] = 1
			} else {
				visibleL[i][j] = 0
			}
		}
	}

	for i := 0; i < 99; i++ {
		current := -1
		for j := 98; j > -1; j-- {
			if forrest[i][j] > current {
				current = forrest[i][j]
				visibleR[i][j] = 1
			} else {
				visibleR[i][j] = 0
			}
		}
	}

	for j := 0; j < 99; j++ {
		current := -1
		for i := 0; i < 99; i++ {
			if forrest[i][j] > current {
				current = forrest[i][j]
				visibleT[i][j] = 1
			} else {
				visibleT[i][j] = 0
			}
		}
	}

	for j := 0; j < 99; j++ {
		current := -1
		for i := 98; i > -1; i-- {
			if forrest[i][j] > current {
				current = forrest[i][j]
				visibleB[i][j] = 1
			} else {
				visibleB[i][j] = 0
			}
		}
	}

	visible := 0
	scenic := 0
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {

			if visibleB[i][j]+visibleT[i][j]+visibleL[i][j]+visibleR[i][j] > 0 {
				visible += 1
			}

			countT := 0
			for t := j - 1; t > -1; t-- {
				countT += 1
				if forrest[i][t] >= forrest[i][j] {
					break
				}
			}

			countB := 0
			for t := j + 1; t < 99; t++ {
				countB += 1
				if forrest[i][t] >= forrest[i][j] {
					break
				}
			}

			countL := 0
			for t := i - 1; t > -1; t-- {
				countL += 1
				if forrest[t][j] >= forrest[i][j] {
					break
				}
			}
			countR := 0
			for t := i + 1; t < 99; t++ {
				countR += 1
				if forrest[t][j] >= forrest[i][j] {
					break
				}
			}

			if scenic < (countL * countR * countT * countB) {
				scenic = countL * countR * countT * countB
			}

		}
	}

	fmt.Println(visible)
	fmt.Println(scenic)

	readFile.Close()
}
