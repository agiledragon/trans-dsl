package transdsl

type Not struct {
	Spec Specification
}

func (this *Not) Ok(transInfo *TransInfo) bool {
	return !this.Spec.Ok(transInfo)
}
