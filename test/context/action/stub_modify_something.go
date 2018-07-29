package action

import (
    "github.com/agiledragon/trans-dsl"
    "github.com/agiledragon/trans-dsl/test/context"
)

type StubModifySomething struct {
}

func (this *StubModifySomething) Exec(transInfo *transdsl.TransInfo) error {
    stubInfo := transInfo.AppInfo.(*context.StubInfo)
    stubInfo.P2 = 22
    return nil
}

func (this *StubModifySomething) Rollback(transInfo *transdsl.TransInfo) {
    stubInfo := transInfo.AppInfo.(*context.StubInfo)
    stubInfo.P2 = 0
}