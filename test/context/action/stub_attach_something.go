package action

import (
	"errors"
	"github.com/agiledragon/trans-dsl"
	"github.com/agiledragon/trans-dsl/test/context"
)

type StubAttachSomething struct {
}

func (this *StubAttachSomething) Exec(transInfo *transdsl.TransInfo) error {
	stubInfo := transInfo.AppInfo.(*context.StubInfo)
	stubInfo.Y = 8
	stubInfo.SpecialNum = 10
	if stubInfo.Abc == "break" && transInfo.LoopIdx == 1 {
		return errors.New("break2")
	}
	if stubInfo.Abc == "continue" && transInfo.LoopIdx == 2 {
		return errors.New("continue2")
	}
	return nil
}

func (this *StubAttachSomething) Rollback(transInfo *transdsl.TransInfo) {
	stubInfo := transInfo.AppInfo.(*context.StubInfo)
	stubInfo.Y = 0
}
