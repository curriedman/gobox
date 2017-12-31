package eq

type Interface interface {
	Eq(other Interface) bool
	Ne(other Interface) bool
}
