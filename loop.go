package transdsl

type Loop struct {
	Spec      Specification
	Fragments []Fragment
	FuncVar   func() Fragment
	times int
}

func (this *Loop) Exec(transInfo *TransInfo) error {
	if this.Spec == nil {
		return this.doWithTimes(transInfo)
	}
	return this.doWithoutTimes(transInfo)
}

func (this *Loop) RollBack(transInfo *TransInfo) {
	for i := this.times - 1; i >= 0; i-- {
		transInfo.LoopIdx = i
		this.Fragments[i].Rollback(transInfo)
	}
}

func (this *Loop) doWithTimes(transInfo *TransInfo) error {
	this.times = transInfo.Times
	this.Fragments = make([]Fragment, this.times)
	for i := 0; i < this.times; i++ {
		transInfo.LoopIdx = i
		this.Fragments[i] = this.FuncVar()
		err := this.Fragments[i].Exec(transInfo)
		if err != nil {
			for j := i - 1; j >= 0; j-- {
				transInfo.LoopIdx = j
				this.Fragments[j].Rollback(transInfo)
			}
			return err
		}
	}
	return nil
}

func (this *Loop) doWithoutTimes(transInfo *TransInfo) error {
	this.Fragments = make([]Fragment, this.times)
	for {
		transInfo.LoopIdx = this.times
		this.times++
		this.Fragments = append(this.Fragments, this.FuncVar())
		err := this.Fragments[transInfo.LoopIdx].Exec(transInfo)
		if err != nil {
			for j := transInfo.LoopIdx - 1; j >= 0; j-- {
				transInfo.LoopIdx = j
				this.Fragments[j].Rollback(transInfo)
			}
			return err
		}
		if this.Spec.Ok(transInfo) {
			break
		}
	}
	return nil
}
