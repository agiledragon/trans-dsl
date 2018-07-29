package spec

import (
	"github.com/agiledragon/trans-dsl"
)

type IsGhiExist struct {
}

func (this *IsGhiExist) Ok(transInfo *transdsl.TransInfo) bool {
	return true
}
