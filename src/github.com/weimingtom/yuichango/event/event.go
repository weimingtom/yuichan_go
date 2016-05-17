package yuichango

type Event struct {
	mSource *Widget
}

func NewEvent(source *Widget) *Widget {
	w := &Widget{}
	w.mSource = source;
	return w
}
   
func (self *Widget) GetSource() *Widget {
	return self.mSource
}


