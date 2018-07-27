package transdsl


type Fragment interface {
    Exec(transInfo *TransInfo) error
    RollBack(transInfo *TransInfo)
}

func forEachFragments(fragments []Fragment, transInfo *TransInfo) (int, error) {
    for i, fragment := range fragments {
        err := fragment.Exec(transInfo)
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
        fragments[index].RollBack(transInfo)
    }
}