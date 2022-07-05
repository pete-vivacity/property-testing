package generator

import (
	"fmt"
	. "reflect"
	"unsafe"
)

func (g Generator) randLength(size float64) int {
	return g.rand.Intn(int(size) + 1)
}

func (g Generator) fillArray(v Value, size float64) {
	for i := 0; i < v.Len(); i++ {
		g.fill(v.Index(i), size)
	}
}

func (g Generator) fillChan(v Value, size float64) {
	if g.randBool(1, 2, size) {
		v.Set(MakeChan(ChanOf(BothDir, v.Type().Elem()), g.randLength(size)))
	}
}

func (g Generator) fillInterface(v Value, size float64) {
	if v.NumMethod() > 0 {
		panic(fmt.Sprintf("can't generate %v because reflect can't generate methods -> https://github.com/golang/go/issues/4146", v.Type()))
	}
	if g.randBool(1, 2, size) {
		v.Set(g.Generate(TypeOf(0), size))
	}
}

func (g Generator) fillMap(v Value, size float64) {
	length := g.randLength(size)
	v.Set(MakeMap(v.Type()))
	for i := 0; i < length; i++ {
		v.SetMapIndex(
			g.Generate(v.Type().Key(), size),
			g.Generate(v.Type().Elem(), size))
	}
}

func (g Generator) fillPointer(v Value, size float64) {
	if g.randBool(1, 2, size) {
		v.Set(g.Generate(v.Type().Elem(), size).Addr())
	} else {
		v.Set(Zero(v.Type()))
	}
}

func (g Generator) fillSlice(v Value, size float64) {
	length := g.randLength(size)
	v.Set(MakeSlice(v.Type(), length, length))
	for i := 0; i < length; i++ {
		g.fill(v.Index(i), size)
	}
}

func (g Generator) fillString(v Value, size float64) {
	length := g.randLength(size)
	bytes := make([]byte, length)
	g.rand.Read(bytes)
	v.SetString(string(bytes))
}

func forceSettable(v Value) Value {
	return NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}

func (g Generator) fillStruct(v Value, size float64) {
	for i := 0; i < v.Type().NumField(); i++ {
		g.fill(forceSettable(v.Field(i)), size)
	}
}

func (g Generator) fillUnsafePointer(v Value, size float64) {
	v.SetPointer(g.Generate(TypeOf((*int)(nil)), size).UnsafePointer())
}
