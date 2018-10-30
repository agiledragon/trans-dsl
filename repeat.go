package transdsl

import (
    "time"
)

type Repeat struct {
    MaxTimes int
    TimeLen  time.Duration //ms
    Fragment Fragment
}

func (this *Repeat) Exec(transInfo *TransInfo) error {
    flag := false
    if this.MaxTimes < 0 {
        flag = true
    }
    
    for i := 0; flag || i < this.MaxTimes; i++ {
        this.Fragment.Exec(transInfo)
        <-time.After(this.TimeLen * time.Millisecond)
    }
    return nil
}

func (this *Repeat) Rollback(transInfo *TransInfo) {
    this.Fragment.Rollback(transInfo)
}

