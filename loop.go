package transdsl

type Loop struct {
	Fragments    []Fragment
	FuncVar      func() Fragment
	BreakErrs    []error
	ContinueErrs []error
	times        int
}

func match(expected error, actuals []error) bool {
	for _, actual := range actuals {
		if isEqual(expected, actual) {
			return true
		}
	}
	return false
}

func (this *Loop) Exec(transInfo *TransInfo) error {
	this.times = transInfo.Times
	this.Fragments = make([]Fragment, this.times)
	for i := 0; i < this.times; i++ {
		transInfo.LoopIdx = i
		this.Fragments[i] = this.FuncVar()
		err := this.Fragments[i].Exec(transInfo)
		if err != nil {
			if match(ErrBreak, this.BreakErrs) {
				break
			}
			if match(ErrContinue, this.ContinueErrs) {
				continue
			}
			for j := i - 1; j >= 0; j-- {
				transInfo.LoopIdx = j
				this.Fragments[j].Rollback(transInfo)
			}
			return err
		}
	}
	return nil
}

func (this *Loop) RollBack(transInfo *TransInfo) {
	for i := this.times - 1; i >= 0; i-- {
		transInfo.LoopIdx = i
		this.Fragments[i].Rollback(transInfo)
	}
}
