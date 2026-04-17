# In-Memory Cache

## Overview

`in-memory-cache` is a lightweight, thread-safe key-value cache with TTL (time-to-live) support in Go. Expired entries are removed lazily on access. No external dependencies.

## Features

- Thread-safe (uses `sync.Mutex`)
- TTL per key with lazy expiration
- Set, Get, and Delete operations
- No external dependencies

## Installation

```sh
go get github.com/dzmitry-sazonau/in-memory-cache
```

## Usage

```go
package main

import (
	"fmt"
	"log"
	"time"

	cache "github.com/dzmitry-sazonau/in-memory-cache"
)

func main() {
	c := cache.New()

	// Set a value with 5 second TTL
	c.Set("userId", 42, time.Second*5)

	// Get a value
	val, err := c.Get("userId")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val) // 42

	// Delete a key
	c.Delete("userId")
}
```

## API Reference

### `New() *Cache`
Creates a new cache instance.

### `Set(key string, value any, ttl time.Duration)`
Stores a value with a given key and TTL duration.

### `Get(key string) (any, error)`
Retrieves a value by key. Returns an error if the key is not found or has expired. Expired entries are deleted on access.

### `Delete(key string)`
Removes a value by key.

## License

This project is licensed under the MIT License.