package yuichango

const (
    SPACE = int(' ')
    TAB = int('\t')
    ENTER = int('\n')
    LEFT_ALT = 1000
    RIGHT_ALT = 1001
    LEFT_SHIFT = 1002
    RIGHT_SHIFT = 1003
    LEFT_CONTROL = 1004
    RIGHT_CONTROL = 1005
    LEFT_META = 1006
    RIGHT_META = 1007
    LEFT_SUPER = 1008
    RIGHT_SUPER = 1009
    INSERT = 1010
    HOME = 1011
    PAGE_UP = 1012
    DELETE = 1013
    END = 1014
    PAGE_DOWN = 1015
    ESCAPE = 1016
    CAPS_LOCK = 1017
    BACKSPACE = 1018
    F1 = 1019
    F2 = 1020
    F3 = 1021
    F4 = 1022
    F5 = 1023
    F6 = 1024
    F7 = 1025
    F8 = 1026
    F9 = 1027
    F10 = 1028
    F11 = 1029
    F12 = 1030
    F13 = 1031
    F14 = 1032
    F15 = 1033
    PRINT_SCREEN = 1034
    SCROLL_LOCK = 1035
    PAUSE = 1036
    NUM_LOCK = 1037
    ALT_GR = 1038
    LEFT = 1039
    RIGHT = 1040
    UP = 1041
    DOWN = 1042
)


type Key struct {
	mValue int
}

func NewKey() *Key {
	k := &Key{}
	k.mValue = 0
	return k 
}

func NewKeyByValue(value int) *Key {
	k := &Key{}
	k.mValue = value
	return k 
}

func (self *Key) IsCharacter() bool {
    return (self.mValue >= 32 && self.mValue <= 126) || 
		(self.mValue >= 162 && self.mValue <= 255) || 
		(self.mValue == 9)
}

func (self *Key) IsNumber() bool {
	return self.mValue >= 48 && self.mValue <= 57
}

func (self *Key) IsLetter() bool {
    return (((self.mValue >= 65 && self.mValue <= 90) || 
			(self.mValue >= 97 && self.mValue <= 122) || 
			(self.mValue >= 192 && self.mValue <= 255)) && 
			(self.mValue != 215) && (self.mValue != 247));
}

func (self *Key) GetValue() int {
	return self.mValue;
}

func (self *Key) IsEqual(key *Key) bool {
	return self.mValue == key.mValue;
}

func (self *Key) IsNotEqual(key *Key) bool {
	return self.mValue != key.mValue;
}
	