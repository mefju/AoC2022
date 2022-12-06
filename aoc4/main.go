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

	pattern := regexp.MustCompile("([0-9]*)-([0-9]*),([0-9]*)-([0-9]*)")

	sum := 0
	sum_overlap := 0

	for fileScanner.Scan() {
		inside := 0
		overlap := 1
		line := strings.Trim(fileScanner.Text(), " \n")
		result := pattern.FindStringSubmatch(line)
		min_a, _ := strconv.Atoi(result[1])
		max_a, _ := strconv.Atoi(result[2])
		min_b, _ := strconv.Atoi(result[3])
		max_b, _ := strconv.Atoi(result[4])

		if min_b >= min_a && max_b <= max_a {
			inside = 1
		} else {
			if min_a >= min_b && max_a <= max_b {
				inside = 1
			}
		}

		if inside > 0 {
			sum += 1
		}

		if min_b > max_a && min_a < min_b {
			overlap = 0
		}
		if min_a > max_b && min_b < min_a {
			overlap = 0
		}

		if overlap == 0 {
			fmt.Println(line)
		}
		sum_overlap += overlap

	}
	fmt.Println(sum)
	fmt.Println(sum_overlap)
	readFile.Close()

}
