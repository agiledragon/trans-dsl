package transdsl

type Optional struct {
	Spec     Specification
	IfFrag   Fragment
	ElseFrag Fragment
	ifFlag   bool
	elseFlag bool
}

func (this *Optional) Exec(transInfo *TransInfo) error {
	if this.Spec.Ok(transInfo) {
		this.ifFlag = true
		return this.IfFrag.Exec(transInfo)
	}
	if this.ElseFrag != nil {
		this.elseFlag = true
		return this.ElseFrag.Exec(transInfo)
	}
	return nil
}

func (this *Optional) Rollback(transInfo *TransInfo) {
	if this.ifFlag {
		this.IfFrag.Rollback(transInfo)
		return
	}

	if this.elseFlag {
		this.ElseFrag.Rollback(transInfo)
	}
}
