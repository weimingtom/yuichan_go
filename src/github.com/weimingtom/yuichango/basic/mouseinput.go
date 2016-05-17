package yuichango

const (
    MouseInput_MOVED = 0
    MouseInput_PRESSED = 1
    MouseInput_RELEASED = 2
    MouseInput_WHEEL_MOVED_DOWN = 3
    MouseInput_WHEEL_MOVED_UP = 4
    
    MouseInput_EMPTY = 0
    MouseInput_LEFT = 1
    MouseInput_RIGHT = 2
    MouseInput_MIDDLE = 3
)
	
type MouseInput struct {
	mType int
    mButton int
    mTimeStamp int
    mX int
    mY int
}

func NewMouseInput() *MouseInput {
	return &MouseInput{}
}

func NewMouseInputByAttr(button int, type_ int,
    x int, y int, timeStamp int) *MouseInput {
    m := &MouseInput{}
	m.mType = type_
    m.mButton = button
    m.mTimeStamp = timeStamp
    m.mX = x
    m.mY = y
	return m
}

func (self *MouseInput) SetType(type_ int) {
	self.mType = type_;
}

func (self *MouseInput) GetType() int {
	return self.mType;
}

func (self *MouseInput) SetButton(button int) {
	self.mButton = button;
}

func (self *MouseInput) GetButton() int {
	return self.mButton;
}

func (self *MouseInput) SetTimeStamp(timeStamp int) {
   	self.mTimeStamp = timeStamp;
}
   
func (self *MouseInput) GetTimeStamp() int {
	return self.mTimeStamp;
}
   
func (self *MouseInput) SetX(x int) {
   	self.mX = x;
}
   
func (self *MouseInput) GetX() int {
	return self.mX;
}
   
func (self *MouseInput) SetY(y int) {
	self.mY = y;
}

func (self *MouseInput) GetY() int {
	return self.mY;
}

	
	