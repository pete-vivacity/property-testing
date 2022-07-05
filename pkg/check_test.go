package pkg

import (
	"testing"
)

func TestFindCounterExample(t *testing.T) {
	cases := []struct {
		name     string
		function any
	}{
		{
			name: "Easy",
			function: func(a int) bool {
				return a < 5
			},
		}, {
			name: "Harder",
			function: func(a, b uint) bool {
				return a < 100 || b > 1
			},
		}, {
			name: "Palindrome",
			function: func(a []int) bool {
				l := len(a)
				for i := 0; i < l/2; i++ {
					if a[i] != a[l-i-1] {
						return false
					}
				}
				return true
			},
		}, {
			name: "Sum",
			function: func(a []int) bool {
				acc := 0
				for i := range a {
					acc += a[i]
				}
				return acc < 1000
			},
		}, {
			name: "MapSum",
			function: func(a map[int]int) bool {
				acc := 0
				for k, v := range a {
					acc += k + v
				}
				return acc < 1000
			},
		}, {
			name: "FuncFive",
			function: func(f func(int) int) bool {
				return f(5) < 100
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			counterExample := findCounterExample(c.function, 1000)
			if counterExample != nil {
				t.Logf("Success, found counterexample: %s", sprintValues(counterExample))
			} else {
				t.Logf("Failed to find a counterexample")
				t.FailNow()
			}
		})
	}
}
