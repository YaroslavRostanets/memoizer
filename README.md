# Memoizer

Added ttl function. Outdated cache values ​​are removed to avoid memory usage.

## Features

- Simple API
- Stores values of any type
- Fast in-memory access
- No external dependencies

## Installation

```bash
go get github.com/YaroslavRostanets/memoizer
```

## Usage

```go
package main

import (
"fmt"
"github.com/YaroslavRostanets/cache"
)

func main() {
c := cache.New()

	c.Set("answer", 42, time.Second)
	c.Set("name", "golang", time.Second)

	fmt.Println(c.Get("answer")) // 42
	fmt.Println(c.Get("name"))   // golang
	fmt.Println(c.Get("unknown")) // <nil> cache item is not set
	time.Sleep(10 * time.Second)
	fmt.Println(c.Get("answer")) // <nil> cache item is not set
	
	c.Delete("answer")
}

```

## API

* New() *Cache — creates a new cache instance
* Get(key string) any — returns value by key
* Set(key string, value any) — stores a value
* Delete(key string) — removes a key
* Close() - stops the channel to check for outdated records

## UPD
