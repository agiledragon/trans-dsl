package action

import (
    "github.com/agiledragon/trans-dsl"
    "github.com/agiledragon/trans-dsl/test/context"
    "errors"
)

type StubActivateSomething struct {

}


func (this *StubActivateSomething) Exec(transInfo *transdsl.TransInfo) error {
    stubInfo := transInfo.AppInfo.(*context.StubInfo)

    if stubInfo.X == "test" {
        return errors.New("something wrong")
    }
    return nil
}

func (this *StubActivateSomething) RollBack(transInfo *transdsl.TransInfo) {

}

