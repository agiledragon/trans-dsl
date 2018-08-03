package transdsl

import (
    "time"
)

type Wait struct {
    EventId string
    Timeout time.Duration //ms
    Fragment Fragment
}


func (this *Wait) Exec(transInfo *TransInfo) error {
    transInfo.EventId = this.EventId
    tc := make(chan struct{})
    go func() {
        <-time.After(this.Timeout * time.Millisecond)
        tc <- struct{}{}
    }()
    select {
    case <-transInfo.Ch:
        return this.Fragment.Exec(transInfo)
    case <-tc:
        return ErrTimeout
    }


}

func (this *Wait) Rollback(transInfo *TransInfo) {
    this.Fragment.Rollback(transInfo)
}