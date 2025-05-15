package dbapi_test

import (
	"fmt"
	"log"

	"github.com/HashemJaafar7/dbapi"
)

func Example() {
	// Initialize and open database
	var database dbapi.DB
	err := dbapi.Open(&database, "example_db")
	if err != nil {
		log.Fatal(err)
	}
	database.DropAll()
	defer database.Close()
	fmt.Println("1. Database opened successfully")

	// Add a new key-value pair
	err = dbapi.Add(database, []byte("name"), []byte("John"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("2. Added key 'name' with value 'John'")

	// Try adding duplicate key
	err = dbapi.Add(database, []byte("name"), []byte("Jane"))
	fmt.Printf("3. Error adding duplicate: %v\n", err)

	// Add another value
	err = dbapi.Add(database, []byte("city"), []byte("London"))
	if err != nil {
		log.Fatal(err)
	}

	// Get the value
	value, err := dbapi.Get(database, []byte("city"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("4. Got value for 'city': %s\n", value)

	// Try getting non-existent key
	_, err = dbapi.Get(database, []byte("country"))
	fmt.Printf("5. Error getting missing key: %v\n", err)

	// Update a value
	err = dbapi.Update(database, []byte("name"), []byte("Jane"))
	if err != nil {
		log.Fatal(err)
	}
	value, _ = dbapi.Get(database, []byte("name"))
	fmt.Printf("6. Updated name to: %s\n", value)

	// Delete a key
	err = dbapi.Delete(database, []byte("city"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("7. Deleted 'city' key")

	// Try to get deleted value
	_, err = dbapi.Get(database, []byte("city"))
	fmt.Printf("8. Error getting deleted key: %v\n", err)

	// Add more data for viewing
	err = dbapi.Add(database, []byte("email"), []byte("jane@example.com"))
	if err != nil {
		log.Fatal(err)
	}

	// View all entries
	fmt.Println("9. All database entries:")
	err = dbapi.View(database, func(key, value []byte) {
		fmt.Printf("   %s: %s\n", key, value)
	})
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	// 1. Database opened successfully
	// 2. Added key 'name' with value 'John'
	// 3. Error adding duplicate: ErrKeyIsUsed : key [110 97 109 101] is used
	// 4. Got value for 'city': London
	// 5. Error getting missing key: ErrKeyNotFound : key [99 111 117 110 116 114 121] not found
	// 6. Updated name to: Jane
	// 7. Deleted 'city' key
	// 8. Error getting deleted key: ErrKeyNotFound : key [99 105 116 121] not found
	// 9. All database entries:
	//    email: jane@example.com
	//    name: Jane
}
