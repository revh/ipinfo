# ipinfo
A wrapper for http://ipinfo.io written in Go language

```go
package main

import (
	"fmt"
	"github.com/RevH/ipinfo"
)

func main() {
	fmt.Println(ipinfo.MyIP())
	fmt.Println(ipinfo.ForeignIP("8.8.8.8"))
}
```
