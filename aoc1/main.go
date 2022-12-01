package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var inputFile = "input.txt"
	readFile, err := os.Open(inputFile)

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var data []int

	for fileScanner.Scan() {
		initVar, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			data = append(data, 0)
		} else {
			data = append(data, initVar)
		}
	}

	readFile.Close()

	var max int = 0
	var sum int = 0
	max = sum
	for _, number := range data {
		if number == 0 {
			if sum > max {
				max = sum
			}
			sum = 0
		} else {
			sum += number
		}

	}

	fmt.Println(max)

	var top3 [3]int = [3]int{0, 0, 0}

	for _, number := range data {
		if number == 0 {

			switch {
			case sum > top3[2]:
				top3[2] = sum
			case sum > top3[1]:
				top3[1] = sum
			case sum > top3[0]:
				top3[0] = sum
			}

			sum = 0
		} else {
			sum += number
		}

	}
	fmt.Println(top3[0] + top3[1] + top3[2])

}
