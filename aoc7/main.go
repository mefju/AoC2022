package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	command_pattern := regexp.MustCompile("\\$ (ls|cd)(.*)")
	ls_pattern := regexp.MustCompile("^([0-9]*|dir) ([a-z.]+)")
	var current_dir []string
	var dir_size map[string]int
	dir_size = make(map[string]int)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if command_pattern.MatchString(line) {
			tmp := command_pattern.FindStringSubmatch(line)
			command := tmp[1]
			argument := tmp[2]

			switch command {
			case "cd":
				if argument == " .." {
					current_dir = current_dir[:len(current_dir)-1]
				} else {
					current_dir = append(current_dir, strings.Trim(argument, " "))
				}
				// fmt.Print("===")
				// fmt.Println(current_dir)
			case "ls":
			}

		}
		if ls_pattern.MatchString(line) {
			tmp := ls_pattern.FindStringSubmatch(line)

			if tmp[1] == "dir" {
				// fmt.Println(strings.Join(current_dir, "_") + "_" + tmp[2])
			} else {

				for i := 0; i < len(current_dir); i++ {
					num, err := strconv.Atoi(tmp[1])
					if err != nil {
						panic(tmp)
					}
					dir_size[strings.Join(current_dir[:len(current_dir)-i], "/")] += num
				}

			}
		}
	}

	sum := 0
	limit := dir_size["/"] - 40000000

	current_to_delete := dir_size["/"]
	for _, v := range dir_size {

		if v < 100000 {
			sum += v
		}

		if v >= limit && v < current_to_delete {
			current_to_delete = v
		}

	}
	fmt.Println(sum)
	fmt.Println(current_to_delete)
	readFile.Close()
}
