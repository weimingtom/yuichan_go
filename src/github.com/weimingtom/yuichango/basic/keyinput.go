package yuichango

const (
    KeyInput_PRESSED = 0
    KeyInput_RELEASED = 1
)

type KeyInput struct {
    mKey *Key
    mType int
    mShiftPressed bool
    mControlPressed bool
    mAltPressed bool
    mMetaPressed bool
    mNumericPad bool
}

func NewKeyInput() *KeyInput {
	return &KeyInput{}
}

func NewKeyInputByKeyType(key *Key, type_ int) *KeyInput {
    i := &KeyInput{}
	i.mKey = key
    i.mType = type_
    i.mShiftPressed = false
    i.mControlPressed = false
    i.mAltPressed = false
    i.mMetaPressed = false
    i.mNumericPad = false
	return i
}

func (self *KeyInput) setType(type_ int) {
	self.mType = type_;
}

func (self *KeyInput) GetType() int {
	return self.mType;
}

func (self *KeyInput) SetKey(key *Key) {
	self.mKey = key;
}

func (self *KeyInput) GetKey() *Key {
	return self.mKey;
}

func (self *KeyInput) IsShiftPressed() bool {
	return self.mShiftPressed;
}

func (self *KeyInput) SetShiftPressed(pressed bool) {
	self.mShiftPressed = pressed;
}

func (self *KeyInput) IsControlPressed() bool {
	return self.mControlPressed;
}

func (self *KeyInput) SetControlPressed(pressed bool) {
	self.mControlPressed = pressed;
}

func (self *KeyInput) IsAltPressed() bool {
	return self.mAltPressed;
}

func (self *KeyInput) SetAltPressed(pressed bool) {
	self.mAltPressed = pressed;
}

func (self *KeyInput) IsMetaPressed() bool {
	return self.mMetaPressed;
}

func (self *KeyInput) SetMetaPressed(pressed bool) {
	self.mMetaPressed = pressed;
}

func (self *KeyInput) IsNumericPad() bool {
	return self.mNumericPad;
}

func (self *KeyInput) SetNumericPad(numpad bool) {
	self.mNumericPad = numpad;
}
