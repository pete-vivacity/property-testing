package generator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	. "reflect"
	"strings"
	"testing"
	"time"
	"unsafe"
)

func TestGenerator(t *testing.T) {

	g := NewGenerator(time.Now().UnixNano())

	cases := []struct {
		name      string
		prototype any
		desired   []any
	}{
		{
			name:      "Bool",
			prototype: false,
			desired:   []any{true},
		}, {
			name:      "Int",
			prototype: 0,
			desired:   []any{1, -1},
		}, {
			name:      "Uint",
			prototype: uint(0),
			desired:   []any{uint(1)},
		}, {
			name:      "Uintptr",
			prototype: uintptr(0),
		}, {
			name:      "Float",
			prototype: 0.0,
		}, {
			name:      "Complex",
			prototype: complex(0, 0),
		}, {
			name:      "Array",
			prototype: [4]int{},
		}, {
			name:      "Chan",
			prototype: (<-chan bool)(nil),
		}, {
			name:      "Interface",
			prototype: struct{ any }{}, // the obvious prototype is `any(nil)` but TypeOf would open that, so we need to wrap it in a struct
		}, {
			name:      "Map",
			prototype: map[int]int{},
			desired:   []any{map[int]int{0: 0}},
		}, {
			name:      "Pointer",
			prototype: (*int)(nil),
		}, {
			name:      "Slice",
			prototype: []int{},
			desired:   []any{[]int{0}},
		}, {
			name:      "String",
			prototype: "",
		}, {
			name:      "Struct",
			prototype: struct{ int }{0},
		}, {
			name:      "UnsafePointer",
			prototype: unsafe.Pointer(nil),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			desiredMap := map[string]bool{}
			desiredMap[fmt.Sprint(c.prototype)] = true
			for i := range c.desired {
				desiredMap[fmt.Sprint(c.desired[i])] = true
			}
			type_ := TypeOf(c.prototype)
			for size := 0.0; size < 5; size++ {
				for i := 0; i < 20; i++ {
					v := g.Generate(type_, size)
					delete(desiredMap, fmt.Sprint(v))
				}
			}
			if len(desiredMap) > 0 {
				missing := make([]string, 0, len(desiredMap))
				for k := range desiredMap {
					missing = append(missing, fmt.Sprint(k))
				}
				assert.Fail(t, "Missing desired values", "[%s]", strings.Join(missing, ", "))
			}
		})
	}
}

// test function generation differently because functions do not have zero values

func TestNilFunc(t *testing.T) {
	var f func()
	g := NewGenerator(time.Now().UnixNano())
	type_ := TypeOf(f)
	f = g.Generate(type_, 5).Interface().(func())
	f()
}

func TestFuncInputs(t *testing.T) {
	var f func(int, int)
	g := NewGenerator(time.Now().UnixNano())
	type_ := TypeOf(f)
	f = g.Generate(type_, 5).Interface().(func(int, int))
	f(1, 2)
}

func TestFuncOutputs(t *testing.T) {
	var f func() (int, int)
	g := NewGenerator(time.Now().UnixNano())
	type_ := TypeOf(f)
	f = g.Generate(type_, 5).Interface().(func() (int, int))
	_, _ = f()
}

func TestFuncInputsAndOutputs(t *testing.T) {
	var f func(int, int) (int, int)
	g := NewGenerator(time.Now().UnixNano())
	type_ := TypeOf(f)
	f = g.Generate(type_, 5).Interface().(func(int, int) (int, int))
	_, _ = f(1, 2)
}
