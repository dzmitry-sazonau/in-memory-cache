# In-Memory Cache 🎯🚀💾

## Overview 📌💡🔍

`in-memory-cache` is a simple, lightweight key-value cache implementation in Go. It provides basic operations to store, retrieve, and delete data in memory without external dependencies. 🎯✅📦

## Features ⚡✨🛠️

- Simple in-memory key-value storage
- Set, Get, and Delete operations
- No external dependencies
- Easy to integrate into any Go project

## Installation 📥🔧💻

To use `in-memory-cache`, simply include the package in your project:

```sh
 go get github.com/dzmitry-sazonau/in-memory-cache
```

## Usage 🏗️📜⚙️

### Import the package 📦📌👨‍💻

```go
package main

import (
	"fmt"
	"github.com/dzmitry-sazonau/in-memory-cache"
)

func main() {
	// Create a new cache instance
	c := cache.New()

	// Set a value
	c.Set("username", "JohnDoe")

	// Get a value
	value := c.Get("username")
	fmt.Println("Username:", value)

	// Delete a key
	c.Delete("username")

	// Try to get a deleted key
	fmt.Println("Deleted key:", c.Get("username"))
}
```

## API Reference 📚📑💡

### `New() *Cache`
Creates a new cache instance. 🆕🗂️🚀

### `Set(key string, value interface{})`
Stores a value associated with a given key. 📝🔑📌

### `Get(key string) interface{}`
Retrieves a value by its key. Returns `"unknown key"` if the key is not found. 🔍📦❓

### `Delete(key string)`
Removes a value by its key. 🗑️🚮✅

## License 📜⚖️🔖

This project is licensed under the MIT License. ✅🔓📄

