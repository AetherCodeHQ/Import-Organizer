package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: import-organizer <file.go>")
		os.Exit(1)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// match "package/path" or "pkgname" inside import blocks
	re := regexp.MustCompile(`"([^"]+)"`)
	matches := re.FindAllStringSubmatch(string(data), -1)
	var stdlib, third []string
	for _, m := range matches {
		p := m[1]
		if strings.Contains(p, ".") {
			third = append(third, p)
		} else {
			stdlib = append(stdlib, p)
		}
	}
	sort.Strings(stdlib)
	sort.Strings(third)
	fmt.Println("=== stdlib ===")
	for _, p := range stdlib {
		fmt.Println(p)
	}
	fmt.Println("=== third-party ===")
	for _, p := range third {
		fmt.Println(p)
	}
	fmt.Printf("\n%d stdlib, %d third-party\n", len(stdlib), len(third))
}
