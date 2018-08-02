package action

import (
	"errors"
	"github.com/agiledragon/trans-dsl"
	"github.com/agiledragon/trans-dsl/test/context"
)

type StubActivateSomething struct {
}

func (this *StubActivateSomething) Exec(transInfo *transdsl.TransInfo) error {
	stubInfo := transInfo.AppInfo.(*context.StubInfo)

	if stubInfo.X == "test" {
		return errors.New("something wrong")
	}
	stubInfo.SpecialNum = 20
	stubInfo.LoopValue++
	return nil
}

func (this *StubActivateSomething) Rollback(transInfo *transdsl.TransInfo) {

}
