package bounded

type Interface interface {
	MinBound() Interface
	MaxBound() Interface
}
