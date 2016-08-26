package yuichango;

import (
	. "github.com/weimingtom/yuichango/basic"
)

type Font interface {
	GetWidth(text string) int
	GetHeight() int	
	DrawString(graphics Graphics, text string, 
		x int, y int) (ex *GuichanException)
}

type FontAbstract struct {
	
}

func (self *FontAbstract) GetStringIndexAt(font Font, text string, x int) int {
    var i int
    var size int
    for i = 0; i < len(text); i++ {
        size = font.GetWidth(text[0 : 0 + i])
        if size > x {
            return i;
        }
    }
    return len(text);
}
	
	
func (self *FontAbstract) Destroy(font Font) {
	
}
	