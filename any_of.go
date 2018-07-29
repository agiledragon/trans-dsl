package transdsl


type AnyOf struct {
	Specs []Specification
}

func (this *AnyOf) Ok(transInfo *TransInfo) bool {
	for _, spec := range this.Specs {
		if spec.Ok(transInfo) {
			return true
		}
	}
	return false
}


