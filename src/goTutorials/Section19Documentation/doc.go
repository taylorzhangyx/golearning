// Package doc The package to demonstrate how the documentation works on GO
package doc

import (
	"fmt"
)

// Show Print out the documentation for Show.
func Show() {
	fmt.Println("Showing the documentation for doc package")

}

// WriteDown Save the doc in the default file.
func WriteDown() {
	fmt.Println("Write down the documentation")
}

// godoc - run the local doc server on localhost:6060
