# DBAPI - BadgerDB Wrapper Library

A simple and efficient wrapper around BadgerDB providing a clean API for basic key-value store operations.

## Features

- **Simple Interface**: Clean API for common database operations
- **Type-Safe**: Strong typing with Go's type system
- **Error Handling**: Consistent error handling with meaningful error messages
- **Transaction Support**: Built-in transaction handling for all operations
- **Full Iterator Support**: Easy iteration over all key-value pairs

## Installation

```bash
go get github.com/HashemJaafar7/dbapi
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/HashemJaafar7/dbapi"
)

func main() {
    // Open database
    var db dbapi.DB
    err := dbapi.Open(&db, "./my-db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Add a key-value pair
    err = dbapi.Add(db, []byte("key1"), []byte("value1"))
    if err != nil {
        panic(err)
    }

    // Get a value
    value, err := dbapi.Get(db, []byte("key1"))
    if err != nil {
        panic(err)
    }
    fmt.Printf("Value: %s\n", value)
}
```

## API Reference

### Open Database

```go
func Open(db *DB, path string) error
```

Opens or creates a database at the specified path.

### Add Key-Value Pair

```go
func Add(db DB, key []byte, value []byte) error
```

Adds a new key-value pair. Returns error if key already exists.

### Update Value

```go
func Update(db DB, key []byte, value []byte) error
```

Updates the value for an existing key.

### Get Value

```go
func Get(db DB, key []byte) ([]byte, error)
```

Retrieves the value associated with a key.

### Delete Key-Value Pair

```go
func Delete(db DB, key []byte) error
```

Removes a key-value pair from the database.

### View All Entries

```go
func View(db DB, function func(key, value []byte)) error
```

Iterates over all key-value pairs in the database.

## Error Handling

The library defines two main error types:

- `ErrKeyIsUsed`: Returned when attempting to add a key that already exists
- `ErrKeyNotFound`: Returned when trying to access a non-existent key

Example error handling:

```go
value, err := dbapi.Get(db, []byte("nonexistent"))
if err != nil {
    switch err.Error() {
    case dbapi.ErrKeyNotFound:
        fmt.Println("Key not found")
    default:
        fmt.Printf("Unknown error: %v\n", err)
    }
}
```

## Best Practices

1. **Database Path**: Use absolute paths when opening databases
2. **Resource Management**: Always close the database when done
3. **Error Handling**: Check errors returned from all operations
4. **Key Design**: Use consistent key naming conventions
5. **Batch Operations**: Use View for efficient batch reading

## Examples

### Iterating Over All Entries

```go
err := dbapi.View(db, func(key, value []byte) {
    fmt.Printf("Key: %s, Value: %s\n", key, value)
})
if err != nil {
    panic(err)
}
```

### Safe Key Addition

```go
err := dbapi.Add(db, []byte("key"), []byte("value"))
if err != nil {
    if err.Error() == dbapi.ErrKeyIsUsed {
        // Handle duplicate key
    } else {
        // Handle other errors
    }
}
```

### Update or Add Pattern

```go
value := []byte("new value")
if err := dbapi.Update(db, key, value); err != nil {
    if err.Error() == dbapi.ErrKeyNotFound {
        err = dbapi.Add(db, key, value)
    }
    if err != nil {
        // Handle error
    }
}
```

## Performance Considerations

- The library uses BadgerDB's default options which are optimized for SSD

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](LICENSE)
