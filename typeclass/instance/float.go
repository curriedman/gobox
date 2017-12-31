package instance

import "github.com/recursivecurry/gobox/typeclass/eq"

type Float float64

const FLOAT Float = Float(0.0)

func (f Float) Eq(other eq.Interface) bool {
	va, ok := other.(Float)
	if !ok {
		return false
	}
	return float64(va) == float64(f)
}

func (f Float) Ne(other eq.Interface) bool {
	return !f.Eq(other)
}
