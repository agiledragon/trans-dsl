package spec

import (
    "github.com/agiledragon/trans-dsl"
    "github.com/agiledragon/trans-dsl/test/context"
)

type IsSomeResourceInsufficient struct {
}

func (this *IsSomeResourceInsufficient) Ok(transInfo *transdsl.TransInfo) bool {
    stubInfo := transInfo.AppInfo.(*context.StubInfo)
    return stubInfo.X == "insufficient"
}
