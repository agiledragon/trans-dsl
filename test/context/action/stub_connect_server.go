package action

import (
	"errors"
	"github.com/agiledragon/trans-dsl"
	"github.com/agiledragon/trans-dsl/test/context"
)

type StubConnectServer struct {
}

func (this *StubConnectServer) Exec(transInfo *transdsl.TransInfo) error {
	stubInfo := transInfo.AppInfo.(*context.StubInfo)
	if stubInfo.Y == -1 {
		return errors.New("panic")
	}
	if stubInfo.FailTimes > 0 {
		stubInfo.FailTimes--
		return errors.New("failed")
	}
	return nil
}

func (this *StubConnectServer) Rollback(transInfo *transdsl.TransInfo) {

}
