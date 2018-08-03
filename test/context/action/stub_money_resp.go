package action

import (
    "github.com/agiledragon/trans-dsl"
)

type StubMoneyResp struct {
}

func (this *StubMoneyResp) Exec(transInfo *transdsl.TransInfo) error {
    return nil
}

func (this *StubMoneyResp) Rollback(transInfo *transdsl.TransInfo) {

}
