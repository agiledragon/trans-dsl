package transdsl

type Concurrent struct {
    Fragments []Fragment
}

func (this *Concurrent) Exec(transInfo *TransInfo) error {
    chs := make([]chan int, len(this.Fragments))
    for i, fragment := range this.Fragments {
        chs[i] = make(chan int)
        go func(ch chan int) {
            fragment.Exec(transInfo)
            ch <- 1
        }(chs[i])
        
    }
    
    for _, ch := range chs {
        <-ch
    }
    return nil
}

func (this *Concurrent) Rollback(transInfo *TransInfo) {

}