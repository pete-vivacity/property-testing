package generator

import (
	"fmt"
	"math/rand"
	. "reflect"
)

type Generator struct {
	rand *rand.Rand
}

func NewGenerator(seed int64) Generator {
	return Generator{
		rand.New(rand.NewSource(seed)),
	}
}

func (g Generator) fill(v Value, size float64) {
	switch v.Type().Kind() {
	case Bool:
		g.fillBool(v, size)
	case Int, Int8, Int16, Int32, Int64:
		g.fillInt(v, size)
	case Uint, Uint8, Uint16, Uint32, Uint64:
		g.fillUint(v, size)
	case Uintptr:
		g.fillUintptr(v, size)
	case Float32, Float64:
		g.fillFloat(v, size)
	case Complex64, Complex128:
		g.fillComplex(v, size)
	case Array:
		g.fillArray(v, size)
	case Chan:
		g.fillChan(v, size)
	case Func:
		g.fillFunc(v, size)
	case Interface:
		g.fillInterface(v, size)
	case Map:
		g.fillMap(v, size)
	case Pointer:
		g.fillPointer(v, size)
	case Slice:
		g.fillSlice(v, size)
	case String:
		g.fillString(v, size)
	case Struct:
		g.fillStruct(v, size)
	case UnsafePointer:
		g.fillUnsafePointer(v, size)
	default:
		panic(fmt.Sprintf("don't know how to handle type %v kind %v", v.Type(), v.Type().Kind()))
	}
}

func (g Generator) Generate(t Type, size float64) Value {
	v := New(t).Elem()
	g.fill(v, size)
	return v
}
