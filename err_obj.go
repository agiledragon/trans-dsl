package transdsl

import (
    "errors"
)

var(
    ErrSucc = errors.New("trans succ")
    ErrContinue = errors.New("loop continue")
    ErrBreak = errors.New("loop break")
)

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

type Continue struct {

}

func (this *Continue) Exec(transInfo *TransInfo) error {
    return ErrContinue
}

func (this *Continue) Rollback(transInfo *TransInfo) {

}

type Break struct {

}

func (this *Break) Exec(transInfo *TransInfo) error {
    return ErrBreak
}

func (this *Break) Rollback(transInfo *TransInfo) {

}