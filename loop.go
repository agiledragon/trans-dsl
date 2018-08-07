package transdsl

type Loop struct {
	FuncVar      func() Fragment
	BreakErrs    []error
	ContinueErrs []error
	fragments    []Fragment
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
	this.fragments = make([]Fragment, this.times)
	for i := 0; i < this.times; i++ {
		transInfo.LoopIdx = i
		this.fragments[i] = this.FuncVar()
		err := this.fragments[i].Exec(transInfo)
		if err != nil {
			if match(err, this.BreakErrs) {
				break
			}
			if match(err, this.ContinueErrs) {
				continue
			}
			for j := i - 1; j >= 0; j-- {
				transInfo.LoopIdx = j
				this.fragments[j].Rollback(transInfo)
			}
			return err
		}
	}
	return nil
}

func (this *Loop) Rollback(transInfo *TransInfo) {
	for i := this.times - 1; i >= 0; i-- {
		transInfo.LoopIdx = i
		this.fragments[i].Rollback(transInfo)
	}
}
