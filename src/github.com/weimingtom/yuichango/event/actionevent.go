package yuichango

type ActionEvent struct {
	Event
	mId string
}

func (self *ActionEvent) GetId() string {
	return self.mId
}
