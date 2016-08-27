package yuichango

import (
	. "github.com/weimingtom/yuichango/gui"
)

type Button struct {
	*Widget
}

func NewButton() *Button {
	b := &Button{}
	return b
}
