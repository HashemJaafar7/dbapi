package go_database_api_test

import (
	"fmt"
	"log"

	db "github.com/HashemJaafar7/go_database_api"
)

func Example() {
	// Initialize and open database
	var database db.DB
	err := db.Open(&database, "example_db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	fmt.Println("1. Database opened successfully")

	// Add a new key-value pair
	err = db.Add(database, []byte("name"), []byte("John"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("2. Added key 'name' with value 'John'")

	// Try adding duplicate key
	err = db.Add(database, []byte("name"), []byte("Jane"))
	fmt.Printf("3. Error adding duplicate: %v\n", err)

	// Add another value
	err = db.Add(database, []byte("city"), []byte("London"))
	if err != nil {
		log.Fatal(err)
	}

	// Get the value
	value, err := db.Get(database, []byte("city"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("4. Got value for 'city': %s\n", value)

	// Try getting non-existent key
	_, err = db.Get(database, []byte("country"))
	fmt.Printf("5. Error getting missing key: %v\n", err)

	// Update a value
	err = db.Update(database, []byte("name"), []byte("Jane"))
	if err != nil {
		log.Fatal(err)
	}
	value, _ = db.Get(database, []byte("name"))
	fmt.Printf("6. Updated name to: %s\n", value)

	// Delete a key
	err = db.Delete(database, []byte("city"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("7. Deleted 'city' key")

	// Try to get deleted value
	_, err = db.Get(database, []byte("city"))
	fmt.Printf("8. Error getting deleted key: %v\n", err)

	// Add more data for viewing
	err = db.Add(database, []byte("email"), []byte("jane@example.com"))
	if err != nil {
		log.Fatal(err)
	}

	// View all entries
	fmt.Println("9. All database entries:")
	err = db.View(database, func(key, value []byte) {
		fmt.Printf("   %s: %s\n", key, value)
	})
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	// 1. Database opened successfully
	// 2. Added key 'name' with value 'John'
	// 3. Error adding duplicate: key [110 97 109 101] is used
	// 4. Got value for 'city': London
	// 5. Error getting missing key: key [99 111 117 110 116 114 121] not found
	// 6. Updated name to: Jane
	// 7. Deleted 'city' key
	// 8. Error getting deleted key: key [99 105 116 121] not found
	// 9. All database entries:
	//    email: jane@example.com
	//    name: Jane
}
