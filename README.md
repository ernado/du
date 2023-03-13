# du

Disk usage package for Go.

```go
package main

import (
    "fmt"
	
    "github.com/ernado/du"
)

func main() {
	info, err := du.Get(".")
	if err != nil {
        panic(err)
    }
	fmt.Printf("%+v", info)
}
```