// Package twofer is short for two for one.
package main

import "fmt"

// ShareWith should share with.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}

func main() {

}
