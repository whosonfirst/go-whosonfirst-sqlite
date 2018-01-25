package main

import (
       "fmt"
       pool "github.com/whosonfirst/go-whosonfirst-pool"
)

func main() {

     p := pool.NewLIFOPool()

     f := pool.PoolInt{Int:int64(123)}

     p.Push(f)
     v, _ := p.Pop()

     fmt.Printf("%d", v.IntValue())
}
