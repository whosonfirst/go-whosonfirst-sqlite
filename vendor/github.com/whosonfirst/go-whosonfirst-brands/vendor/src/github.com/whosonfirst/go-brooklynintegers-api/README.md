# go-brooklynintegers-api

Go package for the Brooklyn Integers API.

## Install

You will need to have both `Go` and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Usage

## Simple

```
package main

import (
	"fmt"
	api "github.com/whosonfirst/go-brooklynintegers-api"
)

func main() {

	client := api.NewAPIClient()
	i, _ := client.CreateInteger()

	fmt.Println(i)
}
```

## Less simple

```
import (
       "fmt"
       api "github.com/whosonfirst/go-brooklynintegers-api"
)

client := api.NewAPIClient()

method := "brooklyn.integers.create"
params := url.Values{}

rsp, err := client.ExecuteMethod(method, &params)

if err != nil {
	return 0, err
}

ints, _ := rsp.Parsed.S("integers").Children()

if len(ints) == 0 {
	return 0, errors.New("Failed to generate any integers")
}

first := ints[0]

f, ok := first.Path("integer").Data().(float64)

if !ok {
	return 0, errors.New("Failed to parse response")
}

i := int64(f)
return i, nil
```

## Tools

### int

Mint one or more Brooklyn Integers.

```
$> ./bin/int -h
Usage of ./bin/int:
  -count int
    	The number of Brooklyn Integers to mint (default 1)
  -tts
    	Output integers to a text-to-speak engine.
  -tts-engine string
    	A valid go-writer-tts text-to-speak engine. Valid options are: osx.
```

### proxy-server

Proxy, pre-load and buffer requests to the Brooklyn Integers `brooklyn.integers.create` API method. No, really.

```
$> ./bin/proxy-server -h
Usage of ./bin/proxy-server:
  -cors=false: Enable CORS headers
  -loglevel="info": Log level
  -min=5: The minimum number of Brooklyn Integers to keep on hand at all times
  -port=8080: Port to listen
```

As in:

```
./bin/proxy-server -min 100
[big-integer] 02:47:44.029143 [error] failed to create new integer, because invalid character '<' looking for beginning of value
...remaining errors excluded for brevity

[big-integer] 02:47:45.431095 [info] time to refill the pool with 100 integers (success: 70 failed: 30): 1.728106507s (pool length is now 61)
[big-integer] 02:47:48.703314 [status] pool length: 61
[big-integer] 02:47:53.704234 [status] pool length: 61
[big-integer] 02:47:54.144543 [info] time to refill the pool with 39 integers (success: 39 failed: 0): 441.226293ms (pool length is now 81)
[big-integer] 02:47:58.704465 [status] pool length: 81
[big-integer] 02:48:03.704680 [status] pool length: 81
[big-integer] 02:48:06.929803 [info] time to refill the pool with 19 integers (success: 19 failed: 0): 3.226286242s (pool length is now 94)
[big-integer] 02:48:08.704911 [status] pool length: 94
[big-integer] 02:48:13.705098 [status] pool length: 94
[big-integer] 02:48:13.904573 [info] time to refill the pool with 6 integers (success: 6 failed: 0): 200.858368ms (pool length is now 100)
[big-integer] 02:48:18.705313 [status] pool length: 100
[big-integer] 02:48:23.705487 [status] pool length: 100
[big-integer] 02:48:28.705684 [status] pool length: 100
```

And then:

```
$> curl http://localhost:8080
404733361
$> curl http://localhost:8080
404733359
```

And then:

```
[big-integer] 02:48:33.705859 [status] pool length: 98
[big-integer] 02:48:33.886058 [info] time to refill the pool with 2 integers (success: 2 failed: 0): 181.959167ms (pool length is now 100)
[big-integer] 02:48:38.706063 [status] pool length: 100
[big-integer] 02:48:43.706231 [status] pool length: 100
```

For event more reporting set the `-loglevel` flag to `debug`.

## See also

* http://brooklynintegers.com/
* http://brooklynintegers.com/api
