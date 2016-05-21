package yuichango

import (
	. "github.com/weimingtom/yuichango/gui"
)

type InputEvent struct {
	*Event
    mShiftPressed bool
    mControlPressed bool
    mAltPressed bool
    mMetaPressed bool
    mIsConsumed bool
}


func NewInputEvent(source *Widget,
	isShiftPressed bool,
	isControlPressed bool,
	isAltPressed bool,
	isMetaPressed bool) *InputEvent {
	e := &InputEvent{}
	e.Event = NewEvent(source)
    e.mShiftPressed = isShiftPressed
    e.mControlPressed = isControlPressed
    e.mAltPressed = isAltPressed
    e.mMetaPressed = isMetaPressed
    e.mIsConsumed = false
	return e
}

func (self *InputEvent) IsShiftPressed() bool {
	return self.mShiftPressed;
}
  
func (self *InputEvent) IsControlPressed() bool {
  	return self.mControlPressed;
}
  
func (self *InputEvent) IsAltPressed() bool {
  	return self.mAltPressed;
}
  
func (self *InputEvent) isMetaPressed() bool {
	return self.mMetaPressed;
}
  
func (self *InputEvent) Consume() {
  	self.mIsConsumed = true;
}
  
func (self *InputEvent) IsConsumed() bool {
  	return self.mIsConsumed;
}

