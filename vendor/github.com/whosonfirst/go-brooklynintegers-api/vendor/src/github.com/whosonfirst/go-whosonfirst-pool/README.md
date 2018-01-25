# go-whosonfirst-pool

A generic LIFO pool derived from Simon Waldherr's [example code](https://github.com/SimonWaldherr/golang-examples/blob/2be89f3185aded00740a45a64e3c98855193b948/advanced/lifo.go). This implementation is safe to use with goroutines.

## Usage

### Simple

```
import (
       "fmt"
       pool "github.com/whosonfirst/go-whosonfirst-pool"
)

func main() {

     p := pool.NewLIFOPool()
     i := pool.PoolInt{Int:int64(123)}

     p.Push(i)
     v, ok := p.Pop()

     if ok {
     	fmt.Printf("%d", v.IntValue())
     }
}
```

### Less simple

The LIFOPool expects and returns object-struct-things that conform to the `PoolItem` interface. These must define both an `IntValue` and a `StringValue` method which return `int64` and `string` values respectively.

The `pool` package comes with it's own a `PoolInt` and a `PoolString` object-struct-things if you'd like to use those.

## Isn't there a better way to do this?

Maybe. I'd love to hear about it.
 
## See also

* https://github.com/SimonWaldherr/golang-examples/blob/2be89f3185aded00740a45a64e3c98855193b948/advanced/lifo.go
