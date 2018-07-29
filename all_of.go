package transdsl

type AllOf struct {
	Specs []Specification
}

func (this *AllOf) Ok(transInfo *TransInfo) bool {
	for _, spec := range this.Specs {
		if !spec.Ok(transInfo) {
			return false
		}
	}
	return true
}
