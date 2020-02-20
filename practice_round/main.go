package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toNumbers(s []string) []int {
	result := []int{}
	for _, v := range s {
		if n, err := strconv.Atoi(v); err == nil {
			result = append(result, n)
		}
	}
	return result
}

func toStrings(s []int) []string {
	result := []string{}
	for _, v := range s {
		n := strconv.Itoa(v)
		result = append(result, n)
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input1 := toNumbers(strings.Fields(scanner.Text()))
	scanner.Scan()
	input2 := toNumbers(strings.Fields(scanner.Text()))

	M, N := input1[0], input1[1]

	result := Solve(M, N, input2)

	Reverse(result)

	fmt.Println(len(result))
	fmt.Println(strings.Join(toStrings(result), " "))
}
