package enum

type Interface interface {
	ToEnum(a int64) Interface
	FromEnum(a Interface) int64
}
