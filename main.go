package main

import "fmt"

func counting1Bits(value uint) int {
	n := 0
	for value != 0 {
		value &= value - 1
		n++
	}
	return n
}

func main() {
	z := counting1Bits(0x11001100)
	fmt.Println(z) // 4 1s in 0x11001100
}
