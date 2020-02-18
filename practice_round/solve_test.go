package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {

	assertSolve := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	}

	cases := []struct {
		m        int
		n        int
		payload  []int
		expected []int
	}{
		{m: 100, n: 1, payload: []int{4}, expected: []int{0}},
		{m: 100, n: 1, payload: []int{100}, expected: []int{0}},
		{m: 17, n: 4, payload: []int{2, 5, 6, 8}, expected: []int{0, 2, 3}},
		{m: 100, n: 10, payload: []int{4, 14, 15, 18, 29, 32, 36, 82, 95, 95}, expected: []int{3, 7}},
		{
			m: 4500, n: 50,
			payload:  []int{7, 12, 12, 13, 14, 28, 29, 29, 30, 32, 32, 34, 41, 45, 46, 56, 61, 61, 62, 63, 65, 68, 76, 77, 77, 92, 93, 94, 97, 103, 113, 114, 114, 120, 135, 145, 145, 149, 156, 157, 160, 169, 172, 179, 184, 185, 189, 194, 195, 195},
			expected: []int{0, 2, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49},
		},
	}

	for idx, c := range cases {
		t.Run(fmt.Sprintf("case#%v", idx), func(t *testing.T) {
			got := Solve(c.m, c.n, c.payload)
			Reverse(got)

			assertSolve(t, got, c.expected)
		})
	}
}
