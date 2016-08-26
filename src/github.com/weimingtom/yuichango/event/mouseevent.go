package yuichango;

import (
	. "github.com/weimingtom/yuichango/gui"
)

const (
    MOVED = 0
    PRESSED = 1
    RELEASED = 2
    WHEEL_MOVED_DOWN = 3
    WHEEL_MOVED_UP = 4
    CLICKED = 5
    EXITED = 7
    DRAGGED = 8
	
    MOUSEEVENT_EMPTY = 0
	MOUSEEVENT_LEFT = 1  
	MOUSEEVENT_RIGHT = 2
	MOUSEEVENT_MIDDLE = 3
)

type MouseEvent struct {
	* InputEvent
	mType int
	mButton int
	MX int
	MY int
	mClickCount int
}

func NewMouseEvent(source *Widget, isShiftPressed bool,
    	isControlPressed bool, isAltPressed bool, isMetaPressed bool,
        type_ int, button int,
        x int, y int, clickCount int) *MouseEvent {
	e := &MouseEvent{}
	e.InputEvent = NewInputEvent(source, 
		isShiftPressed, 
		isControlPressed, 
		isAltPressed, 
		isMetaPressed)
    e.mType = type_;
    e.mButton = button;
    e.MX = x;
    e.MY = y;
    e.mClickCount = clickCount;
	return e;
}

func (self *MouseEvent) GetButton() int {
   	return self.mButton
}
    
func (self *MouseEvent) GetX() int {
	return self.MX
}

func (self *MouseEvent) GetY() int {
	return self.MY
}

func (self *MouseEvent) GetClickCount() int {
	return self.mClickCount
}

func (self *MouseEvent) GetType() int {
	return self.mType
}
	
	