package pkg

import (
	"fmt"
	"github.com/pete-vivacity/property-testing/pkg/generator"
	. "reflect"
	"strings"
	"testing"
	"time"
)

func findCounterExample(f any, trials int) []Value {
	typeOfF := TypeOf(f)
	numIn := typeOfF.NumIn()
	args := make([]Value, numIn)

	g := generator.NewGenerator(time.Now().UnixNano())

	size := 0.0
	for t := 0; t < trials; t++ {
		for i := 0; i < numIn; i++ {
			args[i] = g.Generate(typeOfF.In(i), size)
		}

		if !ValueOf(f).Call(args)[0].Interface().(bool) {
			return args
		}

		size += 0.1
	}
	return nil
}

func sprintValues(values []Value) string {
	argStrings := make([]string, len(values))
	for i := 0; i < len(values); i++ {
		argStrings[i] = fmt.Sprint(values[i])
	}
	return strings.Join(argStrings, ", ")
}

func Check(t *testing.T, f any) {
	typeOfF := TypeOf(f)
	if typeOfF.Kind() != Func || typeOfF.NumOut() != 1 || typeOfF.Out(0).Kind() != Bool {
		t.Fatalf("Check expects to be called with `func(<args>) bool` but was called with `%v`", typeOfF)
	}

	counterExample := findCounterExample(f, 1000)
	if counterExample != nil {
		t.Fatalf("Found a counterexample: {%s}", sprintValues(counterExample))
	}

	t.Logf("Found no counterexample.")
}
