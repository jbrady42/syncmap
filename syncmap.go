package syncmap

import "sync"

type Map struct {
	data map[interface{}]interface{}
	lock *sync.RWMutex
}

type Tuple struct {
	Key interface{}
	Val interface{}
}

func New() *Map {
	lock := new(sync.RWMutex)
	tmp := make(map[interface{}]interface{})
	return &Map{data: tmp, lock: lock}
}

func (t *Map) Set(key, value interface{}) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.data[key] = value
}

func (t *Map) Delete(key interface{}) {
	t.lock.Lock()
	defer t.lock.Unlock()
	delete(t.data, key)
}

func (t *Map) Has(key interface{}) (found bool) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	_, found = t.data[key]
	return found
}

func (t *Map) Get(key interface{}) (value interface{}, found bool) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	value, found = t.data[key]
	return value, found
}

func (t *Map) Len() int {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return len(t.data)
}

func (t *Map) Iter() <-chan Tuple {
	outChan := make(chan Tuple)
	go func() {
		t.lock.RLock()
		for key, value := range t.data {
			outChan <- Tuple{key, value}
		}
		t.lock.RUnlock()
		close(outChan)
	}()
	return outChan
}
