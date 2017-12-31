package consistent

type Interface interface {
	AddHost(host string)
	RemoveHost(host string)
	Hash(object string) uint64
}
