package action

import (
    "github.com/agiledragon/trans-dsl"
    "github.com/agiledragon/trans-dsl/test/context"
)

type StubGetSomething struct {
}

func (this *StubGetSomething) Exec(transInfo *transdsl.TransInfo) error {
    stubInfo := transInfo.AppInfo.(*context.StubInfo)
    stubInfo.P1 = 21
    return nil
}

func (this *StubGetSomething) Rollback(transInfo *transdsl.TransInfo) {
    stubInfo := transInfo.AppInfo.(*context.StubInfo)
    stubInfo.P1 = 0
}
