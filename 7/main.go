package main

import (
	"fmt"
	"strconv"
	"sync"
)

type CMap[V any] struct {
	sync.RWMutex
	Data map[string]V
}

func (cm *CMap[V]) Get(key string) (V, bool) {
	cm.RLock()
	defer cm.RUnlock()
	v, ok := cm.Data[key]
	return v, ok
}

func (cm *CMap[V]) Set(key string, value V) {
	cm.Lock()
	defer cm.Unlock()
	cm.Data[key] = value
}

func (cm *CMap[V]) Has(key string) bool {
	cm.RLock()
	defer cm.RUnlock()
	_, ok := cm.Data[key]
	return ok
}

func (cm *CMap[V]) Delete(key string) {
	cm.Lock()
	defer cm.Unlock()
	delete(cm.Data, key)
}

func NewCMap[V any]() CMap[V] {
	return CMap[V]{
		Data: make(map[string]V),
	}
}

func main() {
	balance := NewCMap[int]()
	wg := sync.WaitGroup{}
	mp := make(map[string]int)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			key := strconv.Itoa(index)
			value := index * 2
			balance.Set(key, value)
		}(i)
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			key := strconv.Itoa(index)
			value := index * 2
			mp[key] = value
		}(i)
	}
	wg.Wait()
	for k, v := range mp {
		fmt.Printf("%s: %d\n", k, v)
	}

}
