package transdsl

import (
    "errors"
)

var(
    ErrSucc = errors.New("trans succ")
    ErrContinue = errors.New("loop continue")
    ErrBreak = errors.New("loop break")
    ErrUnexpectedEvent = errors.New("unexpected event")
    ErrTimeout = errors.New("timeout")
)

func isEqual(leftErr, rightErr error) bool {
    return leftErr.Error() == rightErr.Error()
}

type Succ struct {

}

func (this *Succ) Exec(transInfo *TransInfo) error {
    return ErrSucc
}

func (this *Succ) Rollback(transInfo *TransInfo) {

}

type Fail struct {
    ErrCode error
}

func (this *Fail) Exec(transInfo *TransInfo) error {
    return this.ErrCode
}

func (this *Fail) Rollback(transInfo *TransInfo) {

}
