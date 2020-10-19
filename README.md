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
	if fileName, err := conf.Open(os.Getenv("EXAMPLE_CONF"), "stream.dev.conf", "stream.conf"); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Load config from %s", fileName)
	}

	fmt.Println(conf.GetString("rtsp.timeout.txt"))
	fmt.Println(conf.GetTimeDuration("rtsp.timeout.rtp"))
}
```
