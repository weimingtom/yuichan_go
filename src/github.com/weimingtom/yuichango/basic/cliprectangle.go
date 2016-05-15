package yuichango

import (
	"strconv"
)

type ClipRectangle struct {
    Rectangle
    XOffset int
	YOffset	int
}

func NewClipRectangle(x int, y int, 
	width int, height int,
    xOffset int, yOffset int) *ClipRectangle {
	r := &ClipRectangle{}
	r.X = x
	r.Y = y
	r.Width = width
    r.Height = height
    r.XOffset = xOffset
    r.YOffset = yOffset
	return r
}

func NewClipRectangleByOther(other *Rectangle) {
    r := &ClipRectangle{}
	r.X = other.X
    r.Y = other.Y
    r.Width = other.Width
    r.Height = other.Height
}

func (self *ClipRectangle) String() string {
    return "ClipRectangle [X = " + strconv.Itoa(self.X) +
    	", Y = " + strconv.Itoa(self.Y) +
        ", Width = " + strconv.Itoa(self.Width) + 
        ", Height = " + strconv.Itoa(self.Height) +
        ", XOffset = " + strconv.Itoa(self.XOffset) + 
        ", YOffset = " + strconv.Itoa(self.YOffset) +
        "]"
}
