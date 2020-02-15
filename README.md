# `avg`

[![GoDoc](https://godoc.org/github.com/tidwall/avg?status.svg)](https://godoc.org/github.com/tidwall/avg)

A simple little rolling average package for Go.

## Getting Started

### Installing

To start using `avg`, install Go and run go get:

```
$ go get -u github.com/tidwall/avg
```

This will retrieve the library.

### Usage

There's only one type `Avg`, which has one function `Add`.

Here's an example that continually prints the read performance of crypto 
`rand.Read` over a period of five seconds.

```go
package main

import "crypto/rand"

func main(){
    buf := make([]byte, 0xFFFF)
    start := time.Now()
    var a Avg
    for time.Since(start) < time.Second*5 {
        n, err := rand.Read(buf[:])
        if err != nil {
            panic(err)
        }
        avg := a.Add(n)
        fmt.Printf("\r%.0f MBs/sec ", avg/1024/1024)
    }
    fmt.Printf("\n")
}

// output: 615 MBs/sec
```

### Contact

Josh Baker [@tidwall](https://twitter.com/tidwall)

### License

`avg` source code is available under the MIT License.
