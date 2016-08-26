package yuichango;

import (
	. "github.com/weimingtom/yuichango/gui"
)

type SelectEvent struct {
	*Event
}

func NewSelectEvent(source *Widget) *SelectEvent {
	e := &SelectEvent{}
	e.Event = NewEvent(source)
	return e
}

