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