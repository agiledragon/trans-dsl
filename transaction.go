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
