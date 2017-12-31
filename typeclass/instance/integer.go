package instance

import (
	"math"
	"github.com/recursivecurry/gobox/typeclass/eq"
	"github.com/recursivecurry/gobox/typeclass/enum"
	"github.com/recursivecurry/gobox/typeclass/bounded"
)

type Integer int64

const INTEGER Integer = Integer(0)

func (i Integer) ToEnum(a int64) enum.Interface {
	return Integer(a)
}

func (i Integer) FromEnum(a enum.Interface) int64 {
	val := a.(Integer)

	return int64(val)
}

func (i Integer) Eq(other eq.Interface) bool {
	va, ok := other.(Integer)
	if !ok {
		return false
	}
	return int64(va) == int64(i)
}

func (i Integer) Ne(other eq.Interface) bool {
	return !i.Eq(other)
}

func (i Integer) MinBound() bounded.Interface {
	return Integer(math.MinInt64)
}

func (i Integer) MaxBound() bounded.Interface {
	return Integer(math.MaxInt64)
}

func (i Integer) Less(other eq.Interface) bool {
	i2, ok := other.(Integer)
	if !ok {
		return false
	}
	return int64(i) < int64(i2)
}
