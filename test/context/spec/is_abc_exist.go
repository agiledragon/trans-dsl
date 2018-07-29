package spec

import (
	"github.com/agiledragon/trans-dsl"
	"github.com/agiledragon/trans-dsl/test/context"
)

type IsAbcExist struct {
}

func (this *IsAbcExist) Ok(transInfo *transdsl.TransInfo) bool {
	stubInfo := transInfo.AppInfo.(*context.StubInfo)
	return stubInfo.Abc == "abc"
}
