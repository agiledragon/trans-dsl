package action

import (
    "github.com/agiledragon/trans-dsl"
)

type StubAssignCmd struct {
}

func (this *StubAssignCmd) Exec(transInfo *transdsl.TransInfo) error {
    return nil
}

func (this *StubAssignCmd) Rollback(transInfo *transdsl.TransInfo) {

}
