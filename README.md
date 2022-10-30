# cache

```golang
go get github.com/dmytrodemianchuk/cache
```

```golang
package main

import (
	"fmt"

	"github.com/dmytrodemianchuk/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId, _ := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId, _ = cache.Get("userId")

	fmt.Println(userId)
}
