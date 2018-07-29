package spec

import (
	"github.com/agiledragon/trans-dsl"
)

type IsDefExist struct {
}

func (this *IsDefExist) Ok(transInfo *transdsl.TransInfo) bool {
	return true
}
