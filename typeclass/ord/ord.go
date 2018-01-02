package ord

import "github.com/recursivecurry/gobox/typeclass/eq"

type Interface interface{
	eq.Interface
	Less(other Interface) bool
}
