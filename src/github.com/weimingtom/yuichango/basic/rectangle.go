package yuichango

import (
	"strconv"
)

type Rectangle struct {
    X int
    Y int
    Width int
    Height int	
}

func NewRectangle() *Rectangle {
	r := &Rectangle{}
	r.X = 0
	r.Y = 0
	r.Width = 0
	r.Height = 0
	return r;
}

func NewRectangleByRectangle(r_ *Rectangle) *Rectangle {
	r := &Rectangle{}
	r.X = r_.X
	r.Y = r_.Y
	r.Width = r_.Width
	r.Height = r_.Height
	return r
}

func NewRectangleByAttr(x_ int, y_ int, 
	width_ int, height_ int) *Rectangle {
	r := &Rectangle{}
	r.X = x_
	r.Y = y_
	r.Width = width_
	r.Height = height_
	return r
}

func (self *Rectangle) SetAll(x_ int, y_ int, 
	width_ int, height_ int) {
	self.X = x_
	self.Y = y_
	self.Width = width_
	self.Height = height_	
}

func (self *Rectangle) IsIntersecting(
	rectangle *Rectangle) bool {
    x_ := self.X;
    y_ := self.Y;
    width_ := self.Width;
    height_ := self.Height;
    x_ -= rectangle.X;
    y_ -= rectangle.Y;
    if (x_ < 0) {
        width_ += x_;
        x_ = 0;
    } else if (x_ + width_ > rectangle.Width) {
        width_ = rectangle.Width - x_;
    }
    if (y_ < 0) {
        height_ += y_;
        y_ = 0;
    } else if (y_ + height_ > rectangle.Height) {
        height_ = rectangle.Height - y_;
    }
    if (width_ <= 0 || height_ <= 0) {
        return false;
    }
    return true;
}

func (self *Rectangle) IsPointInRect(
	x_ int, y_ int) bool {
    return x_ >= self.X && 
   	   y_ >= self.Y && 
   	   x_ < self.X + self.Width && 
   	   y_ < self.Y + self.Height
}

func (self *Rectangle) String() string {
    return "Rectangle [x = " + strconv.Itoa(self.X) +
    	", y = " + strconv.Itoa(self.Y) +
        ", width = " + strconv.Itoa(self.Width) + 
        ", height = " + strconv.Itoa(self.Height) +
        "]";
}
	