package action

import (
    "github.com/agiledragon/trans-dsl"
)

type StubMoneyReq struct {
}

func (this *StubMoneyReq) Exec(transInfo *transdsl.TransInfo) error {
    return nil
}

func (this *StubMoneyReq) Rollback(transInfo *transdsl.TransInfo) {

}

