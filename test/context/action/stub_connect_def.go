package action

import (
	"github.com/agiledragon/trans-dsl"
	"github.com/agiledragon/trans-dsl/test/context"
)

type StubConnectDef struct {
}

func (this *StubConnectDef) Exec(transInfo *transdsl.TransInfo) error {
	stubInfo := transInfo.AppInfo.(*context.StubInfo)
	stubInfo.Y = 3
	return nil
}

func (this *StubConnectDef) Rollback(transInfo *transdsl.TransInfo) {
	stubInfo := transInfo.AppInfo.(*context.StubInfo)
	stubInfo.Y = -1
}
