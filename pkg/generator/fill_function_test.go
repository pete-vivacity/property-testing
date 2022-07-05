package generator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoise(t *testing.T) {
	count := 0
	var fibonacci func(int) int
	fibonacci = Memoise(func(n int) int {
		if n == 0 || n == 1 {
			return 1
		}
		count++
		return fibonacci(n-1) + fibonacci(n-2)
	})

	const n = 1000
	fmt.Println(fibonacci(n))
	assert.Equal(t, count, n-1)
}
