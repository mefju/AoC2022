package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	stacks := make(map[int][]byte)
	stacks2 := make(map[int][]byte)

	stack_patters := regexp.MustCompile("(^[[])")
	move_patterns := regexp.MustCompile("^move ([0-9]*) from ([0-9]*) to ([0-9]*)")
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if stack_patters.MatchString(line) {

			for i := 0; i < 9; i++ {
				if line[1+i*4] != ' ' {
					stacks[i+1] = append(stacks[i+1], line[1+i*4])
					stacks2[i+1] = append(stacks2[i+1], line[1+i*4])
				}
			}
		}

		if move_patterns.MatchString(line) {
			tmp := move_patterns.FindStringSubmatch(line)
			number, err := strconv.Atoi(tmp[1])
			if err != nil {
				panic(err)
			}
			from, err := strconv.Atoi(tmp[2])
			if err != nil {
				panic(err)
			}
			to, err := strconv.Atoi(tmp[3])

			if err != nil {
				panic(err)
			}

			for i := 0; i < number; i++ {
				stacks[to] = append([]byte{stacks[from][0]}, stacks[to]...)
				stacks[from] = stacks[from][1:len(stacks[from])]
			}

			fr_move := make([]byte, number)
			fr_stay := make([]byte, len(stacks2[from])-number)

			copy(fr_move, stacks2[from][:number])
			copy(fr_stay, stacks2[from][number:])

			stacks2[to] = append(fr_move, stacks2[to]...)
			stacks2[from] = fr_stay

			// for _, k := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
			// 	fmt.Printf("%d:", k)
			// 	for _, v := range stacks2[k] {
			// 		fmt.Printf("%c", v)
			// 	}
			// 	fmt.Println("|")
			// }

		}

	}

	for _, k := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		fmt.Printf("%c", stacks[k][0])
	}

	fmt.Println()
	for _, k := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		fmt.Printf("%c", stacks2[k][0])
	}
	fmt.Println()
	readFile.Close()

}
