package transdsl

type Specification interface {
	Ok(transInfo *TransInfo) bool
}

type Not struct {
	Spec Specification
}

func (this *Not) Ok(transInfo *TransInfo) bool {
	return !this.Spec.Ok(transInfo)
}

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
