package transdsl

import (
    "errors"
)

var(
    ErrSucc = errors.New("trans succ")
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