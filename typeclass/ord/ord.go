package ord

import "github.com/recursivecurry/gobox/typeclass/eq"

type Interface interface{
	Less(other eq.Interface) bool
}
