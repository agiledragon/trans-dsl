package transdsl

type Loop struct {
	Spec      Specification
	Fragments []Fragment
	FuncVar   func() Fragment
}

func (this *Loop) Exec(transInfo *TransInfo) error {
	if this.Spec == nil {
		return this.doWithTimes(transInfo)
	}
	return this.doWithoutTimes(transInfo)
}

func (this *Loop) RollBack(transInfo *TransInfo) {
	for i := transInfo.Times - 1; i >= 0; i-- {
		transInfo.LoopIdx = i
		this.Fragments[i].Rollback(transInfo)
	}
}

func (this *Loop) doWithTimes(transInfo *TransInfo) error {
	this.Fragments = make([]Fragment, transInfo.Times)
	for i := uint(0); i < transInfo.Times; i++ {
		transInfo.LoopIdx = i
		this.Fragments[i] = this.FuncVar()
		err := this.Fragments[i].Exec(transInfo)
		if err != nil {

			if err == ErrBreak {
				break
			}

			if err == ErrContinue {
				continue
			}

			if i == 0 {
				return err
			}
			i--
			for j := i; j >= 0; j-- {
				transInfo.LoopIdx = j
				this.Fragments[j].Rollback(transInfo)
			}
			return err
		}
	}
	return nil
}

func (this *Loop) doWithoutTimes(transInfo *TransInfo) error {
	this.Fragments = make([]Fragment, transInfo.Times)
	for {
		transInfo.LoopIdx = transInfo.Times
		transInfo.Times++
		this.Fragments = append(this.Fragments, this.FuncVar())
		err := this.Fragments[transInfo.LoopIdx].Exec(transInfo)
		if err != nil {
			if this.Spec.Ok(transInfo) || err == ErrBreak {
				break
			}

			if err == ErrContinue {
				continue
			}

			if transInfo.LoopIdx == 0 {
				return err
			}
			transInfo.LoopIdx--
			for j := transInfo.LoopIdx; j >= 0; j-- {
				transInfo.LoopIdx = j
				this.Fragments[j].Rollback(transInfo)
			}
			return err
		}
	}
	return nil
}
