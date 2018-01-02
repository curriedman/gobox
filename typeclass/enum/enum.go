package enum

type Interface interface {
	FuncToEnum(a int64) Interface
	FromEnum() int64
}
