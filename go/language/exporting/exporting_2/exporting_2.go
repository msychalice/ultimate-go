// ------------------------------------------
// Access a value of an unexported identifier
// ------------------------------------------

package main

import (
	"fmt"

	"github.com/hoanhan101/ultimate-go/go/language/exporting/exporting_2/counters"
)

func main() {

	// Create a variable of the unexported type using the exported New function
	// from the package counters.
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)

	// Can't access the unexported type directly
	// value := counters.user{name: "msy", email: "msy@gmail.com"}

	// But the user type can be using through a factory function
	value := counters.NewUser("msy", "msy@gmail.com")
	fmt.Println(value)
}
