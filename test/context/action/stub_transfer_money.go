package action

import (
    "github.com/agiledragon/trans-dsl"
)

type StubTransferMoney struct {
}

func (this *StubTransferMoney) Exec(transInfo *transdsl.TransInfo) error {
    return nil
}

func (this *StubTransferMoney) Rollback(transInfo *transdsl.TransInfo) {

}

