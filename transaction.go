package transdsl

type Transaction struct {
	Fragments []Fragment
}

func (this *Transaction) Start(transInfo *TransInfo) error {
	index, err := forEachFragments(this.Fragments, transInfo)
	if err != nil {
		backEachFragments(this.Fragments, transInfo, index)
	}
	return err
}

func (this *Transaction) HandleEvent(EventId string, transInfo *TransInfo) error {
	if EventId != transInfo.EventId {
		return ErrUnexpectedEvent
	}
	transInfo.Ch <- struct{}{}
	transInfo.EventId = ""
	return nil
}
