package transdsl

type TransInfo struct {
	//framework paras
	Times   int
	LoopIdx int
	EventId string
	Ch      chan struct{}

	//user info
	AppInfo interface{}
}
