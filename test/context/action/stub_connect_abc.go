package action

import (
	"github.com/agiledragon/trans-dsl"
	"github.com/agiledragon/trans-dsl/test/context"
)

type StubConnectAbc struct {
}

func (this *StubConnectAbc) Exec(transInfo *transdsl.TransInfo) error {
	stubInfo := transInfo.AppInfo.(*context.StubInfo)
	stubInfo.Y = 2
	return nil
}

func (this *StubConnectAbc) Rollback(transInfo *transdsl.TransInfo) {
	stubInfo := transInfo.AppInfo.(*context.StubInfo)
	stubInfo.Y = 0
}
