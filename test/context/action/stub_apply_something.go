package action

import (
	"github.com/agiledragon/trans-dsl"
	"github.com/agiledragon/trans-dsl/test/context"
)

type StubApplySomething struct {
}

func (this *StubApplySomething) Exec(transInfo *transdsl.TransInfo) error {
	stubInfo := transInfo.AppInfo.(*context.StubInfo)
	stubInfo.P3 = 23
	return nil
}

func (this *StubApplySomething) Rollback(transInfo *transdsl.TransInfo) {
	stubInfo := transInfo.AppInfo.(*context.StubInfo)
	stubInfo.P3 = 0
}
