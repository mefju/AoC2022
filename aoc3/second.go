package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var counter byte = 0
	sum := 0
	first := make(map[byte]byte)

	for fileScanner.Scan() {
		fmt.Println(counter)

		line := strings.Trim(fileScanner.Text(), " \n")

		for i := 0; i < len(line); i++ {

			if first[line[i]] == counter {
				first[line[i]] = counter + 1
			}
		}

		if counter == 2 {

			for k, v := range first {
				if v == 3 {
					if k < 91 {
						sum += int(k - 64 + 26)
					} else {
						sum += int(k - 96)
					}
				}

			}

			counter = 0
			first = make(map[byte]byte)

		} else {
			counter += 1
		}

	}
	readFile.Close()

	fmt.Println(sum)
}
