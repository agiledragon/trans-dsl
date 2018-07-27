package action

import (
    "github.com/agiledragon/trans-dsl"
    "github.com/agiledragon/trans-dsl/test/context"
)

type StubAttachSomething struct {

}


func (this *StubAttachSomething) Exec(transInfo *transdsl.TransInfo) error {
    stubInfo := transInfo.AppInfo.(*context.StubInfo)
    stubApplyCertainResource(&stubInfo.Y)
    return nil
}

func (this *StubAttachSomething) RollBack(transInfo *transdsl.TransInfo) {
    stubInfo := transInfo.AppInfo.(*context.StubInfo)
    stubReleaseCertianResource(&stubInfo.Y)
}

func stubApplyCertainResource(a *int) {
    *a = 8
}

func stubReleaseCertianResource(a *int) {
    *a = 0
}