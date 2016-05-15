package yuichango

import (
	"strconv"
)

type Color struct {
	R int
	G int
	B int
	A int
}

func NewColor() *Color {
	c := &Color{}
	c.R = 0
	c.G = 0
	c.B = 0
	c.A = 255
	return c;
}

func NewColorByInt(color int) *Color {
	c := &Color{}
    c.R = (color >> 16) & 0xFF
    c.G = (color >>  8) & 0xFF
    c.B =  color        & 0xFF
    c.A =  255;	
	return c;
}

func NewColorByRGB(r int, g int, b int) *Color {
    return NewColorByRGBA(r, g, b, 255)
}

func NewColorByRGBA(ar int, ag int, ab int, aa int) *Color {
    c := &Color{}
	c.R = ar
    c.G = ag
    c.B = ab
    c.A = aa
	return c;
}

func (self *Color) Plus (color *Color) *Color {
    result := NewColorByRGBA(self.R + color.R, 
        self.G + color.G, 
        self.B + color.B, 
        255)
	if result.R > 255 {
    	result.R = 255
    } else if result.R < 0 {
		result.R = 0
    }
	if result.G > 255 {
    	result.G = 255
    } else if result.G < 0 {
		result.G = 0
    }
	if result.B > 255 {
    	result.B = 255
    } else if result.B < 0 {
		result.B = 0
    }
    return result
}

func (self *Color) Minus (color *Color) *Color {
    result := NewColorByRGBA(self.R - color.R, 
        self.G - color.G, 
        self.B - color.B, 
        255)
	if result.R > 255 {
    	result.R = 255
    } else if result.R < 0 {
		result.R = 0
    }
	if result.G > 255 {
    	result.G = 255
    } else if result.G < 0 {
		result.G = 0
    }
	if result.B > 255 {
    	result.B = 255
    } else if result.B < 0 {
		result.B = 0
    }
    return result
}

func (self *Color) Multi (value float32) *Color {
    result := NewColorByRGBA(int(float32(self.R) * value), 
        int(float32(self.G) * value), 
        int(float32(self.B) * value), 
        self.A)
	if result.R > 255 {
    	result.R = 255
    } else if result.R < 0 {
		result.R = 0
    }
	if result.G > 255 {
    	result.G = 255
    } else if result.G < 0 {
		result.G = 0
    }
	if result.B > 255 {
    	result.B = 255
    } else if result.B < 0 {
		result.B = 0
    }
    return result
}

func (self *Color) IsEqual(color *Color) bool {
    return self.R == color.R && 
		self.G == color.G && 
		self.B == color.B && 
		self.A == color.A
}

func (self *Color) IsNotEqual(color *Color) bool {
    return !(self.R == color.R && 
		self.G == color.G && 
		self.B == color.B && 
		self.A == color.A)
}

func (self *Color) String() string {
	return "Color [R = " +
		strconv.Itoa(self.R) +
		", G = " +
		strconv.Itoa(self.G) +
		", B = " +
		strconv.Itoa(self.B) +
		", A = " + 
		strconv.Itoa(self.A) + 
		"]"
}
