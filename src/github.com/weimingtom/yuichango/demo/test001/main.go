package main

import (
	"fmt"
	. "github.com/weimingtom/yuichango/basic"
)

func main() {
	fmt.Print("yuichango start...\n")
	fmt.Printf("yuichango version : %s\n", GcnGuichanVersion())
	NewColor()
	NewColorByInt(0)
	NewColorByRGB(0, 0, 0)
	c1 := NewColorByRGBA(1, 2, 3, 4)
	fmt.Printf("color = %s\n", c1)
}
