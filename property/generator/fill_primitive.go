package generator

import . "reflect"

// maxPercentTrue: the highest chance of getting True for arbitrarily large N
// sizeToReachHalfway: the size required for the chance of getting True to reach maxPercentTrue/2
func (g Generator) randBool(maxPercentTrue, sizeToReachHalfway float64, size float64) bool {
	threshold := maxPercentTrue * size / (size + sizeToReachHalfway)
	return g.rand.Float64() < threshold
}

func (g Generator) fillBool(v Value, size float64) {
	v.SetBool(g.randBool(0.5, 2, size))
}

func (g Generator) fillInt(v Value, size float64) {
	v.SetInt(int64(g.rand.NormFloat64() * size))
}

func (g Generator) fillUint(v Value, size float64) {
	v.SetUint(uint64(g.rand.ExpFloat64() * size))
}

func (g Generator) fillUintptr(v Value, size float64) {
	v.SetUint(uint64(uintptr(g.Generate(TypeOf((*int)(nil)), size).UnsafePointer())))
}

func (g Generator) fillComplex(v Value, size float64) {
	v.SetComplex(complex(g.rand.NormFloat64()*size, g.rand.NormFloat64()*size))
}

func (g Generator) fillFloat(v Value, size float64) {
	v.SetFloat(g.rand.NormFloat64() * size)
}
