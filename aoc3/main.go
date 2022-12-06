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

	sum := 0
	for fileScanner.Scan() {
		line := strings.Trim(fileScanner.Text(), " \n")

		first := make(map[byte]byte)

		for i := 0; i < len(line)/2; i++ {

			first[line[i]] = 1
		}

		for i := len(line) / 2; i < len(line); i++ {

			if first[line[i]] > 0 {
				first[line[i]] = 2
			}
		}

		for k, v := range first {
			if v == 2 {
				if k < 91 {
					sum += int(k - 64 + 26)
				} else {
					sum += int(k - 96)
				}
			}

		}

	}
	readFile.Close()

	fmt.Println(sum)
}
