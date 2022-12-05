package main

import (
	"bufio"
	"fmt"
	"os"
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

	sum := 0
	sumb := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		switch line {
		case "A X":
			sum += 4
			sumb += 3
		case "A Y":
			sum += 8
			sumb += 4
		case "A Z":
			sum += 3
			sumb += 8
		case "B X":
			sum += 1
			sumb += 1
		case "B Y":
			sum += 5
			sumb += 5
		case "B Z":
			sum += 9
			sumb += 9
		case "C X":
			sum += 7
			sumb += 2
		case "C Y":
			sum += 2
			sumb += 6
		case "C Z":
			sum += 6
			sumb += 7
		}

	}

	fmt.Println(sum)
	fmt.Println(sumb)
	readFile.Close()

}
