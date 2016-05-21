package yuichango

import (
	. "github.com/weimingtom/yuichango/basic"
	. "github.com/weimingtom/yuichango/gui"
)

const (
	KeyEvent_PRESSED = 0
    KeyEvent_RELEASED = 1
)

type KeyEvent struct {
	*InputEvent
    mType int
    mIsNumericPad bool
    mKey *Key
}

func NewKeyEvent(source *Widget,
	isShiftPressed bool,
	isControlPressed bool,
	isAltPressed bool,
	isMetaPressed bool,
	type_ int, 
	isNumericPad bool, 
	key *Key) *KeyEvent {
	e := &KeyEvent{}
	e.InputEvent = NewInputEvent(source,
                isShiftPressed,
                isControlPressed,
                isAltPressed,
                isMetaPressed)
	e.mType = type_;
	e.mIsNumericPad = isNumericPad;
	e.mKey = key;
	return e
}

func (self *KeyEvent) GetType() int {
	return self.mType
}

func (self *KeyEvent) IsNumericPad() bool {
	return self.mIsNumericPad
}

func (self *KeyEvent) GetKey() *Key {
	return self.mKey
}   


