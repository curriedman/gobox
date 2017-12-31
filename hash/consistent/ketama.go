package consistent

import (
	"container/heap"
	"fmt"
	"github.com/recursivecurry/gobox/container"
	"hash/fnv"
	"log"
	"sort"
)

type Host struct {
	Key  uint64
	Name string
}

type Hosts []Host

func (h Hosts) Len() int {
	return len(h)
}

func (h Hosts) Less(i, j int) bool {
	return h[i].Key < h[j].Key
}

func (h *Hosts) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

type Ketama struct {
	factor  int
	hosts   Hosts
	hostSet map[string]struct{}
}

func NewKetama(factor int) *Ketama {
	return &Ketama{
		factor:  factor,
		hosts:   Hosts{},
		hostSet: map[string]struct{}{},
	}
}

func (k *Ketama) AddHost(host string) {

	if _, ok := k.hostSet[host]; ok {
		return
	}

	k.hostSet[host] = struct{}{}

	//for r :=0; r < k.factor; r++ {
	//	vhost := fmt.Sprintf("%s-%d", host, r)
	//	//for {
	//		//vhKey := internalHash(vhost)
	//		//found := sort.Search(k.hosts, func(i int) bool {
	//		//	return k.hosts[i] >= vhKey
	//		//})
	//		//heap.Push(&k.cacheMap, container.ValueWithDistance{vhKey, host})
	//	//}
	//}
}

func (k *Ketama) RemoveHost(host string) {
	if _, ok := k.hostSet[host]; !ok {
		return
	}

	delete(k.hostSet, host)
}

func (k *Ketama) Hash(object string) uint64 {
}

func internalHash(key string) uint64 {
	h := fnv.New64a()
	if _, err := h.Write([]byte(key)); err != nil {
		log.Fatalf("Fail to write\n")
	}
	return h.Sum64()
}
