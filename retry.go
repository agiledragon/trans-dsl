package transdsl

import "time"

type Retry struct {
	MaxTimes int
	TimeLen  time.Duration //ms
	Fragment Fragment
	Errs     []error
}

func (this *Retry) Exec(transInfo *TransInfo) error {
	var err error
	for i := 0; i < this.MaxTimes; i++ {
		err = this.Fragment.Exec(transInfo)
		if err == nil {
			return nil
		}

		for _, e := range this.Errs {
			if isEqual(e, err) {
				return err
			}
		}

		<-time.After(this.TimeLen * time.Millisecond)
	}
	return err
}

func (this *Retry) Rollback(transInfo *TransInfo) {
	this.Fragment.Rollback(transInfo)
}
