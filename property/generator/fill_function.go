package generator

import . "reflect"

func makeKeyOfArgsFunc(v Value) func([]Value) interface{} {
	fields := make([]StructField, v.Type().NumIn())
	for i := 0; i < v.Type().NumIn(); i++ {
		fields[i] = StructField{
			Name: string(rune('A' + i)),
			Type: v.Type().In(i),
		}
	}
	argsStruct := StructOf(fields)

	keyOfArgs := func(values []Value) interface{} {
		s := New(argsStruct).Elem()
		for i := 0; i < len(values); i++ {
			s.Field(i).Set(values[i])
		}
		return s.Interface()
	}

	return keyOfArgs
}

func memoise(v Value) Value {
	keyOfArgs := makeKeyOfArgsFunc(v)
	cache := map[interface{}][]Value{}
	return MakeFunc(v.Type(), func(args []Value) []Value {
		key := keyOfArgs(args)
		if results, ok := cache[key]; ok {
			return results
		}
		results := v.Call(args)
		cache[key] = results
		return results
	})
}

func Memoise[T any](f T) T {
	memoised := New(TypeOf(f)).Elem()
	memoised.Set(memoise(ValueOf(f)))
	return memoised.Interface().(T)
}

func (g Generator) randomFunc(v Value, size float64) Value {
	return MakeFunc(v.Type(), func([]Value) []Value {
		results := make([]Value, v.Type().NumOut())
		for i := 0; i < v.Type().NumOut(); i++ {
			results[i] = g.Generate(v.Type().Out(i), size)
		}
		return results
	})
}

func (g Generator) fillFunc(v Value, size float64) {
	v.Set(memoise(g.randomFunc(v, size)))
}
