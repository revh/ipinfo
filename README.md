# ipinfo
A wrapper for http://ipinfo.io in Go

```go
package main

import (
	"fmt"
	"github.com/RevH/ipinfo"
)

func main() {
    fmt.Println(ipinfo.MyInfo())
	fmt.Println(ipinfo.ForeignInfo("8.8.8.8"))
}
```
