package pool

import (
	"strconv"
	"sync"
	"sync/atomic"
)

type PoolItem interface {
	StringValue() string
	IntValue() int64
}

type PoolInt struct {
	PoolItem
	Int int64
}

func (i PoolInt) StringValue() string {
	return strconv.FormatInt(i.Int, 10)
}

func (i PoolInt) IntValue() int64 {
	return i.Int
}

type PoolString struct {
	PoolItem
	String string
}

func (s PoolString) StringValue() string {
	return s.String
}

func (s PoolString) IntValue() int64 {
	return int64(0)
}

// https://github.com/SimonWaldherr/golang-examples/blob/2be89f3185aded00740a45a64e3c98855193b948/advanced/lifo.go

type LIFOPool struct {
	nodes []PoolItem
	count int64
	mutex *sync.Mutex
}

func NewLIFOPool() *LIFOPool {

	mu := new(sync.Mutex)
	nodes := make([]PoolItem, 0)

	return &LIFOPool{
		mutex: mu,
		nodes: nodes,
		count: 0,
	}
}

func (pl *LIFOPool) Length() int64 {

	pl.mutex.Lock()
	defer pl.mutex.Unlock()

	return pl.count
}

func (pl *LIFOPool) Push(i PoolItem) {

	pl.mutex.Lock()
	defer pl.mutex.Unlock()

	pl.nodes = append(pl.nodes[:pl.count], i)
	atomic.AddInt64(&pl.count, 1)
}

func (pl *LIFOPool) Pop() (PoolItem, bool) {

	pl.mutex.Lock()
	defer pl.mutex.Unlock()

	if pl.count == 0 {
		return nil, false
	}

	atomic.AddInt64(&pl.count, -1)
	i := pl.nodes[pl.count]

	return i, true
}
