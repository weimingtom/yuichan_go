package main

import (
	"fmt"
	. "github.com/weimingtom/yuichango/basic"
	. "github.com/weimingtom/yuichango/event"
	. "github.com/weimingtom/yuichango/gui"
	. "github.com/weimingtom/yuichango/platform"	
	. "github.com/weimingtom/yuichango/font"	
)

func main() {
	fmt.Print("yuichango start...\n")
	fmt.Printf("yuichango version : %s\n", GcnGuichanVersion())
	NewColor()
	NewColorByInt(0)
	NewColorByRGB(0, 0, 0)
	c1 := NewColorByRGBA(1, 2, 3, 4)
	fmt.Printf("color = %s\n", c1)
	NewClipRectangle(0, 0, 0, 0, 0, 0)
	GcnGuichanVersion()
	NewKey()
	NewKeyInput()
	NewMouseInput()
	NewRectangle()
	NewEvent(nil)
	iEvn := NewInputEvent(&Widget{}, false, false, false, false)
	fmt.Printf("%v\n", iEvn.GetSource())
	NewKeyEvent(nil, false, false, false, false, 0, false, nil)
	NewMouseEvent(nil, false, false, false, false, 0, 0, 0, 0, 0)
	NewSelectEvent(nil)
	
	var f Font = &DefaultFont{}
	f.GetWidth("")
}
