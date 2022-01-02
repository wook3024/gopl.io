// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package main

import (
	"fmt"
	"time"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
func ModifiedPopCount(x uint64) int {
	var sum byte
	for i := 0; i < 8; i++ {
		sum += pc[byte(x>>(i*8))]
	}
	return int(sum)

}

func main() {
	var num uint64 = 25
	start := time.Now()
	fmt.Printf("PopCount result %d\n", PopCount(num))
	elapsed := time.Since(start)
	fmt.Printf("PopCount Elapsed time %s\n", elapsed)
	start = time.Now()
	fmt.Printf("ModifiedPopCount result %d\n", ModifiedPopCount(num))
	elapsed = time.Since(start)
	fmt.Printf("ModifiedPopCount Elapsed time %s\n", elapsed)
}

//!-
