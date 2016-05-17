package yuichango

type InputEvent struct {
	Event
    mShiftPressed bool
    mControlPressed bool
    mAltPressed bool
    mMetaPresse bool
    mIsConsumed bool
}


func NewInputEvent(source *Widget,
	isShiftPressed bool,
	isControlPressed bool,
	isAltPressed bool,
	isMetaPressed bool) {
	e := NewEvent(source)
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
  
func (self *InputEvent) IsControlPressed() {
  	return self.mControlPressed;
}
  
func (self *InputEvent) IsAltPressed() {
  	return self.mAltPressed;
}
  
func (self *InputEvent) isMetaPressed() bool {
	return self.mMetaPressed;
}
  
func (self *InputEvent) Consume() {
  	self.mIsConsumed = true;
}
  
func (self *InputEvent) IsConsumed() bool {
  	return mIsConsumed;
}

