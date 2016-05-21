package yuichango

import (
	. "github.com/weimingtom/yuichango/gui"
)

type Event struct {
	mSource *Widget
}

func NewEvent(source *Widget) *Event {
	w := &Event{}
	w.mSource = source;
	return w
}
   
func (self *Event) GetSource() *Widget {
	return self.mSource
}


