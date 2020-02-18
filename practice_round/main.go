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
		n, err := strconv.Atoi(v)
		if err == nil {
			result = append(result, n)
		}
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

	fmt.Println(M, N)
	fmt.Println(input2)
	result := Solve(M, N, input2)

	Reverse(result)

	fmt.Println(len(result))
	fmt.Println(result)
}
