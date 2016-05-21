package yuichango

import (
	. "github.com/weimingtom/yuichango/gui"
)

type ActionEvent struct {
	*Event
	mId string
}

func NewActionEvent(source *Widget, id string) *ActionEvent {
	e := &ActionEvent{};
	e.Event = NewEvent(source);
	e.mId = id;
	return e
}

func (self *ActionEvent) GetId() string {
	return self.mId
}
