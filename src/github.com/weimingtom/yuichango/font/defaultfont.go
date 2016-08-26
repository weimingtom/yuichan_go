package yuichango;

import (
	. "github.com/weimingtom/yuichango/basic"
	. "github.com/weimingtom/yuichango/platform"
)

type DefaultFont struct {
	
}

func (self *DefaultFont) GetWidth(text string) int {
	return 0
}

func (self *DefaultFont) GetHeight() int {
	return 0
}	

func (self *DefaultFont) DrawString(graphics Graphics, text string, 
	x int, y int) (ex *GuichanException) {
	ex = nil
	return
}
