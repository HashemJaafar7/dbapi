package go_database_api

import (
	"fmt"

	badger "github.com/dgraph-io/badger/v4"
)

type DB = *badger.DB

// Open initializes and opens a Badger database at the specified path.
// It takes a pointer to a DB instance and a string representing the database path.
// The function modifies the provided DB pointer to point to the opened database.
// Returns an error if the database fails to open.
//
// Parameters:
//   - db: A pointer to a DB instance that will be initialized and opened.
//   - path: The file path where the Badger database is located.
//
// Returns:
//   - error: An error object if the database fails to open, otherwise nil.
func Open(db *DB, path string) error {
	var err error
	*db, err = badger.Open(badger.DefaultOptions(path))
	return err
}

// Delete removes a key-value pair from the database.
// It takes a DB instance and a key as input parameters.
// The function performs a database update operation to delete the specified key.
// If the operation is successful, it returns nil; otherwise, it returns an error.
//
// Parameters:
//   - db: The database instance implementing the DB interface.
//   - key: The key to be deleted, represented as a byte slice.
//
// Returns:
//   - error: An error if the delete operation fails, or nil if it succeeds.
//
// if the key is not found, it returns nil.
func Delete(db DB, key []byte) error {
	err := db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
	return err
}

// Add inserts a key-value pair into the database if the key does not already exist.
// If the key is already in use, it returns an error indicating the key conflict.
//
// Parameters:
//   - db: The database instance implementing the DB interface.
//   - key: The key to be added as a byte slice.
//   - value: The value to be associated with the key as a byte slice.
//
// Returns:
//   - error: An error if the key already exists or if there is an issue during the update operation.
func Add(db DB, key []byte, value []byte) error {
	err := db.Update(func(txn *badger.Txn) error {
		_, err := txn.Get(key)
		if err == nil {
			return fmt.Errorf("key %v is used", key)
		}
		return txn.Set(key, value)
	})
	return err
}

// Update updates the value associated with the given key in the database.
// It performs the update operation within a transaction.
//
// Parameters:
//   - db: The database instance implementing the DB interface.
//   - key: The key for which the value needs to be updated.
//   - value: The new value to be associated with the key.
//
// Returns:
//   - error: An error if the update operation fails, otherwise nil.
func Update(db DB, key []byte, value []byte) error {
	err := db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
	return err
}

// Get retrieves the value associated with the given key from the provided database.
// It performs a read-only transaction to fetch the value.
//
// Parameters:
//   - db: The database instance implementing the DB interface.
//   - key: A byte slice representing the key to look up.
//
// Returns:
//   - A byte slice containing the value associated with the key, or an error if the key is not found
//     or if any other issue occurs during the transaction.
//
// Errors:
//   - Returns an error if the key is not found, with a message indicating the missing key.
//   - Propagates any other errors encountered during the transaction.
func Get(db DB, key []byte) ([]byte, error) {
	var valCopy []byte

	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		valCopy, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return nil, fmt.Errorf("key %v not found", key)
		}
		return nil, err
	}

	return valCopy, nil
}

// View iterates over all key-value pairs in the provided BadgerDB instance
// and applies the given function to each pair. The function takes two
// arguments: the key and the value as byte slices.
//
// Parameters:
//   - db: The BadgerDB instance to perform the read-only transaction on.
//   - function: A callback function that processes each key-value pair.
//
// Returns:
//   - error: An error if the transaction fails, or nil if successful.
//
// Example usage:
//
//	err := View(db, func(key, value []byte) {
//	    fmt.Printf("Key: %s, Value: %s\n", key, value)
//	})
//	if err != nil {
//	    log.Fatalf("Error viewing database: %v", err)
//	}
func View(db DB, function func(key, value []byte)) error {
	err := db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			item.Value(func(value []byte) error {
				function(item.Key(), value)
				return nil
			})
		}
		return nil
	})
	return err
}
