With `gocached`, Golang developers can simply add in-memory caching to their applications without having to worry about the underlying implementation. It is lightweight and easy-to-use. It is thread-safe and provides a simple API for storing and retrieving any type of data.

⚠️ This was created as an exercise and prototype for an internal tool. It's not fully-featured and PRs are welcome.

## Getting Started

### Installation

```bash
go get github.com/lazharichir/gocached
```

### Usage

```golang
import (
    "context"
    "fmt"
    "time"

    "github.com/lazharichir/gocached"
)

func main() {
    ctx := context.Background()

    // Add an optional Eviction Policy that will be used to evict items from the cache
    // - LRU: Least Recently Used
    // - LFU: Least Frequently Used
    lru := gocached.NewLRUEvicter[string]()

    // Create a new cache instance:
    // - the cache keys are of type string
    // - the cache values are of type string
    cache := gocached.NewCache[string, int](lru, gocached.WithDefaultTTL(5*time.Second))

    // Set a value in the cache (key-value)
    cache.Set(ctx, "the_key_1", 123456789, gocached.WithTTL(10*time.Second))

    // Has checks if a key exists in the cache
    if cache.Has(ctx, "the_key_1") {
        fmt.Println("the_key_1 exists in the cache")
    } else {
        fmt.Println("the_key_1 does not exist in the cache [BIG PROBLEM]")
    }

    // Get a value from the cache
    // - value is automatically typed to int
    // - found is a boolean indicating if the key was found in the cache
    // - err is an error if something went wrong
    value, found, err := cache.Get(ctx, "the_key_1")
    if err != nil {
        panic(err)
    }
    if found {
        fmt.Println("the_key_1 was found:", value) // prints "the_key_1 was found: 123456789"
    } else {
        fmt.Println("the_key_1 was not found [ANOTHER BIG PROBLEM]")
    }

    // Delete a value from the cache
    cache.Delete(ctx, "the_key_1")
}
```

## Generics

`gocached` is written in Golang, and makes use of Generics to ensure type safety at compile time. This means that you can store any type of data in the cache, and retrieve it without having to worry about type conversions, by creating specific cache instances.

```golang
type User struct {
    ID string
    Name string
    Age int
}

type Product struct {
    ID string
    Name string
    Price float64
}

userCache := gocached.NewCache[string, User](nil)
productCache := gocached.NewCache[string, Product](nil)

usr, found, err := userCache.Get(ctx, "user-1")
// usr is of type User

product, found, err := productCache.Get(ctx, "product-1")
// product is of type Product
```

If you prefer to store bytes in the cache, you can use the `[]byte` type and then simply convert it to the type you want.

```golang
bytesCache := gocached.NewCache[string, []byte](nil)
```

## Evictions

`gocached` supports two types of eviction policies:
* Least Recently Used (LRU)
* Least Frequently Used (LFU)

By default, there is no eviction policy, and the cache will grow indefinitely. You can add an eviction policy when creating a new cache instance.

```golang
lru := gocached.NewLRUEvicter[string]()
lruCache := gocached.NewCache[string, int](lru)

lfu := gocached.NewLFUEvicter[string]()
lfuCache := gocached.NewCache[string, int](lfu)
```

## Options

`gocached` supports a few options that can be used to customize the default cache behavior, as well as specific cached entries.

### Default TTL

The default TTL (Time To Live) is a cache-wide option that can be used to set a default expiration time for all cached entries. It can be overridden when setting a new entry in the cache.

```golang
cache := gocached.NewCache[string, int](nil, gocached.WithDefaultTTL(5*time.Second))
```

### Entry TTL

The entry TTL is an option that can be used to set a specific expiration time for a cached entry. It overrides the default TTL.

```golang
cache := gocached.NewCache[string, int](nil, gocached.WithDefaultTTL(5*time.Second))
cache.Set(ctx, "the_key_1", 123456789, gocached.WithTTL(10*time.Second))
```

### Entry Expiry Date

The entry expiry date is an option that can be used to set a specific expiry date for a cached entry.

```golang
cache := gocached.NewCache[string, int](nil)
cache.Set(ctx, "the_key_1", 123456789, gocached.WithExpiryDate(time.Now().Add(10*time.Second)))
```

## Thread Safety

`gocached` is thread-safe, and can be used in a multi-threaded environment. The four main entry points (Set, Get, Has, Del) make use of a sync.RWMutex.

## Feature Ideas

`gocached` is still in its early stages, and there is a lot of room for improvement. The following features would make for great additions:

* [ ] Add a maximum size option to the cache
* [ ] Add optional callbacks for cache events (e.g. on evict, on set, on get, on del)
* [ ] Batch operations (e.g. BatchSet, BatchGet, BatchHas, BatchDel)
* [ ] Measure and improve performance
* [ ] Add optional metrics (e.g. number of hits, misses, evictions)

## Miscellanous

### But Is It Super Fast Though?

[First make it work, then make it right, and finally make it fast.](https://wiki.c2.com/?MakeItWorkMakeItRightMakeItFast) `gocached` is still young and performance has not been scientifically measured and benchmarked yet. There are already a few low-hanging fruits to improve performance.

### Contribution

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make to `gocached` are greatly appreciated. Whether it be fixing bugs, proposing new features, or discussing potential improvements, your input helps shape the library into a better tool for everyone.

Steps to contribute:

1. **Fork** the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a **Pull Request**

### License

Distributed under the MIT License. See LICENSE for more information.