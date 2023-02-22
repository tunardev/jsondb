## 🌎 JSON Database

🚀 A simple JSON database for Go. It uses a JSON file as a database. It is very easy to use.

## 📦 Installation

```bash
go get github.com/tunardev/jsondb
```

## 📝 Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/tunardev/jsondb"
)

func main() {
	// Create a new database
	db, err := jsondb.New("db.json")
	if err != nil {
		panic(err)
	}

	// Set a value
	db.Set("person.name", "Tunar")
	db.Set("person.age", 14)
	db.Set("person.friends", []interface{}{"John", "Doe"})

	// Get a value
	name, err := db.Get("person.name")
	if err != nil {
		panic(err)
	}
	fmt.Println(name) // Output: Tunar

	// Increment a value
	db.Add("person.age", 1)
	age, err := db.Get("person.age") // Output: 15
	if err != nil {
		panic(err)
	}
	fmt.Println(age)

	// Decrement a value
	db.Sub("person.age", 1)
	age, err = db.Get("person.age") // Output: 14
	if err != nil {
		panic(err)
	}
	fmt.Println(age)

	// Check if a value exists
	fmt.Println(db.Has("person.name")) // Output: true

	// Push a value
	db.Push("person.friends", "Jane")
	if friends, err := db.Get("person.friends"); err == nil {
		fmt.Println(friends) // Output: [John Doe Jane]
	}

	// Delete a value
	db.Delete("person")

	fmt.Println(db.Has("person.name")) // Output: false
}
```