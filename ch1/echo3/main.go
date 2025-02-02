// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {
	var _value []string
	for i, arg := range os.Args {
		_value = append(_value, fmt.Sprintf("%d %s", i, arg))
	}
	fmt.Println(strings.Join(_value, "\n"))
}

//!-
