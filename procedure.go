package transdsl

type Procedure struct {
    Fragments []Fragment
}

func (this *Procedure) Exec(transInfo *TransInfo) error {
    index, err := forEachFragments(this.Fragments, transInfo)
    if err != nil {
        if index <= 0 {
            return err
        }
        backEachFragments(this.Fragments, transInfo, index)
    }
    return err
}

func (this *Procedure) Rollback(transInfo *TransInfo) {
    backEachFragments(this.Fragments, transInfo, len(this.Fragments))
}

