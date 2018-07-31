package transdsl

import "time"

type Retry struct {
	MaxTimes uint
	TimeLen  time.Duration
	Frag     Fragment
	Errs     []error
}

func (this *Retry) Exec(transInfo *TransInfo) error {
	var err error
	for i := uint(0); i < this.MaxTimes; i++ {
		err = this.Frag.Exec(transInfo)
		if err == nil {
			return nil
		}

		for _, e := range this.Errs {
			if e.Error() == err.Error() {
				return err
			}
		}

		<-time.After(this.TimeLen * time.Millisecond)
	}
	return err
}

func (this *Retry) Rollback(transInfo *TransInfo) {

}
