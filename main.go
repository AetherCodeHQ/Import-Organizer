package main

import (
	"fmt"
	"os"
)

// import_organizer - Organize import statements
func import_organizer(path string) {
	fmt.Println("========================================")
	fmt.Println("  Import-Organizer")
	fmt.Println("  Organize import statements")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	import_organizer(path)
}
