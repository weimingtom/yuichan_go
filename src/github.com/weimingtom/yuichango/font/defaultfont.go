package yuichango;

import (
	. "github.com/weimingtom/yuichango/basic"
	. "github.com/weimingtom/yuichango/platform"
)

type DefaultFont struct {
	
}

func (self *DefaultFont) GetWidth(text string) int {
	return 8 * len(text);
}

func (self *DefaultFont) GetHeight() int {
	return 8;
}

func (self *DefaultFont) DrawGlyph(graphics Graphics, 
	glyph rune, x int, y int) int {
    graphics.DrawRectangle(NewRectangleByAttr(x, y, 8, 8));
    return 8
}

func (self *DefaultFont) DrawString(graphics Graphics, text string, 
	x int, y int) (ex *GuichanException) {
	ex = nil
	text2 := []rune(text)
	for i := 0; i < len(text2); i++ {
		self.DrawGlyph(graphics, text2[i], x, y)
        x += self.GetWidth(string(text2[i])) //FIXME:
    }
	return
}

func (self *DefaultFont) GetStringIndexAt(text string, x int) int {
    if x > len([]rune(text)) * 8 {
        return len([]rune(text));
    }
    return x / 8;
}
