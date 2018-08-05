package spec

import (
	"github.com/agiledragon/trans-dsl"
	"github.com/agiledragon/trans-dsl/test/context"
)

type IsSpecialConditionSatisfied struct {
}

func (this *IsSpecialConditionSatisfied) Ok(transInfo *transdsl.TransInfo) bool {
	stubInfo := transInfo.AppInfo.(*context.StubInfo)
	return stubInfo.X == "special"
}
