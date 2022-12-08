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

	fileScanner.Split(bufio.ScanBytes)
	characters := ""
	counter := 0
	for fileScanner.Scan() {
		counter = counter + 1
		character := fileScanner.Text()
		characters = character + characters
		if len(characters) > 13 {
			characterMap := make(map[rune]int)
			for _, k := range characters[0:14] {
				characterMap[k] = 1
			}
			if len(characterMap) > 13 {
				fmt.Println(counter)
				break
			}

		}
	}
	readFile.Close()

}
