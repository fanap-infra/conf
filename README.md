[HOCON Docs](https://github.com/typesafehub/config/blob/master/HOCON.md).

`stream.conf`
```conf
rtsp {
    timeout {
        rtp = 5s
        txt = "Hossein"
    }
}
```

`example.go`
```go
package main

import (
	"fmt"
	"github.com/fanap-infra/conf"
)

func main() {
	if err := conf.Open("stream.conf"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println(conf.GetString("rtsp.timeout.txt"))
	fmt.Println(conf.GetTimeDuration("rtsp.timeout.rtp"))
}
```
