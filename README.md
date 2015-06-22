# ipinfo
A wrapper for http://ipinfo.io written in Go language

```go
package main

import (
	"log"
	"github.com/RevH/ipinfo"
)

func main() {
	myIP, err := ipinfo.MyIP()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(myIP)

	foreignIP, err := ipinfo.ForeignIP("8.8.8.8")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(foreignIP)
}
```
