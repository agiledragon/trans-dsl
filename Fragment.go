package transdsl

type Fragment interface {
	Exec(transInfo *TransInfo) error
	Rollback(transInfo *TransInfo)
}

func forEachFragments(fragments []Fragment, transInfo *TransInfo) (int, error) {
	for i, fragment := range fragments {
		err := fragment.Exec(transInfo)
		if err == ErrSucc {
			return 0, nil
		}
		
		if err != nil {
			return i, err
		}
	}
	return 0, nil
}

func backEachFragments(fragments []Fragment, transInfo *TransInfo, index int) {
	if index <= 0 {
		return
	}
	index--
	for ; index >= 0; index-- {
		fragments[index].Rollback(transInfo)
	}
}
