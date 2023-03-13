# du [![Go Reference](https://pkg.go.dev/badge/github.com/ernado/du.svg)](https://pkg.go.dev/github.com/ernado/du)

Disk usage package for Go.

```go
go get github.com/ernado/du
```

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