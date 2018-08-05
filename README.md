# Transaction DSL

trans-dsl is a transaction model framework in golang,  can be used in complex business scenes.

## Introduction
In some complex domains, a business process may involve the interaction of multiple messages, which may be synchronous messages, asynchronous messages, or even system calls. We explicitly model the operation behavior of message interaction, and call the operation behavior of a message interaction as Action, then the flow chart of the business corresponds to an Action sequence. In general, a single scenario business process corresponds to a transaction process. Based on the transaction model framework, business developers can not only express all business processes in a simple and complete manner at a higher level, but also apply and maintain transaction models at low cost.

## Features
+ support synchronous messages
+ support system calls
+ support asynchronous messages

## Installation
------------
	$ go get github.com/agiledragon/trans-dsl
	
## Keyword
+ optional
+ loop
+ retry
+ wait
+ procedure
+ succ
+ fail
+ not
+ allof
+ anyof

## Using Transaction DSL
Here just make some tests as typical examples.
**Please refer to the test cases, very complete and detailed.**


### optional
```go
import (
	"github.com/agiledragon/trans-dsl"
	"github.com/agiledragon/trans-dsl/test/context"
	"github.com/agiledragon/trans-dsl/test/context/action"
	"github.com/agiledragon/trans-dsl/test/context/spec"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func newIfTrans() *transdsl.Transaction {
	trans := &transdsl.Transaction{
		Fragments: []transdsl.Fragment{
			&transdsl.Optional{
				Spec:   new(spec.IsAbcExist),
				IfFrag: new(action.StubConnectAbc),
			},
			new(action.StubActivateSomething),
		},
	}
	return trans
}

func newElseTrans() *transdsl.Transaction {
	trans := &transdsl.Transaction{
		Fragments: []transdsl.Fragment{
			&transdsl.Optional{
				Spec:     new(spec.IsAbcExist),
				IfFrag:   new(action.StubConnectAbc),
				ElseFrag: new(action.StubConnectDef),
			},
			new(action.StubActivateSomething),
		},
	}
	return trans
}

func TestIfTrans(t *testing.T) {
	trans := newIfTrans()
	Convey("TestIfTrans", t, func() {

		Convey("trans exec succ when spec is true", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					Abc: "abc",
					Y: 1,
				},
			}
			err := trans.Start(transInfo)
			So(err, ShouldEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).Y, ShouldEqual, 2)
		})

		Convey("trans exec succ when spec is false", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					Abc: "def",
					Y: 1,
				},
			}
			err := trans.Start(transInfo)
			So(err, ShouldEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).Y, ShouldEqual, 1)
		})

		Convey("iffrag rollback", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					X: "test",
					Abc: "abc",
					Y: 1,
				},
			}
			err := trans.Start(transInfo)
			So(err, ShouldNotEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).Y, ShouldEqual, 0)
		})
	})
}

func TestElseTrans(t *testing.T) {
	trans := newElseTrans()
	Convey("TestElseTrans", t, func() {

		Convey("trans exec succ when spec is false", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					Abc: "def",
					Y: 1,
				},
			}
			err := trans.Start(transInfo)
			So(err, ShouldEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).Y, ShouldEqual, 3)
		})

		Convey("elsefrag rollback", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					X: "test",
					Abc: "def",
					Y: 1,
				},
			}
			err := trans.Start(transInfo)
			So(err, ShouldNotEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).Y, ShouldEqual, -1)
		})
	})
}

```

### loop
```go
import (
	"errors"
	"github.com/agiledragon/trans-dsl"
	"github.com/agiledragon/trans-dsl/test/context"
	"github.com/agiledragon/trans-dsl/test/context/action"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func newLoopTrans() *transdsl.Transaction {
	trans := &transdsl.Transaction{
		Fragments: []transdsl.Fragment{
			new(action.StubGetSomething),
			&transdsl.Loop{
				FuncVar:      newLoopProcedure,
				BreakErrs:    []error{errors.New("break1"), errors.New("break2")},
				ContinueErrs: []error{errors.New("continue1"), errors.New("continue2")},
			},
		},
	}
	return trans
}

func newLoopProcedure() transdsl.Fragment {
	procedure := &transdsl.Procedure{
		Fragments: []transdsl.Fragment{
			new(action.StubAttachSomething),
			new(action.StubActivateSomething),
		},
	}
	return procedure
}

func TestLoopTrans(t *testing.T) {
	trans := newLoopTrans()
	Convey("TestLoopTrans", t, func() {

		Convey("trans exec succ when loop 3 times", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					LoopValue: 1,
				},
			}
			err := trans.Start(transInfo)
			So(err, ShouldEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).LoopValue, ShouldEqual, 4)
		})

		Convey("trans exec succ when second time break", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					LoopValue: 1,
					Abc:       "break",
				},
			}
			err := trans.Start(transInfo)
			So(err.Error(), ShouldEqual, "break2")
			So(transInfo.AppInfo.(*context.StubInfo).LoopValue, ShouldEqual, 2)

		})

		Convey("trans exec succ when third time continue", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					LoopValue: 1,
					Abc:       "continue",
				},
			}
			err := trans.Start(transInfo)
			So(err.Error(), ShouldEqual, "continue2")
			So(transInfo.AppInfo.(*context.StubInfo).LoopValue, ShouldEqual, 3)

		})

		Convey("trans exec fail", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					LoopValue: 1,
					X:         "test",
				},
			}
			err := trans.Start(transInfo)
			So(err, ShouldNotEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).LoopValue, ShouldEqual, 1)
		})
	})
}

```

### retry
```go
import (
    "github.com/agiledragon/trans-dsl"
    "github.com/agiledragon/trans-dsl/test/context"
    "github.com/agiledragon/trans-dsl/test/context/action"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
    "errors"
)

func newTryTrans() *transdsl.Transaction {
    trans := &transdsl.Transaction{
        Fragments: []transdsl.Fragment{
            &transdsl.Retry{
                MaxTimes: 3,
                TimeLen: 100,
                Fragment: new(action.StubConnectServer),
                Errs: []error{errors.New("fatal"), errors.New("panic")},
            },
        },
    }
    return trans
}

func TestTryTrans(t *testing.T) {
    trans := newTryTrans()
    Convey("TestTryTrans", t, func() {

        Convey("trans exec succ when fail time is 1", func() {
            transInfo := &transdsl.TransInfo{
                AppInfo: &context.StubInfo{
                    FailTimes: 1,
                },
            }
            err := trans.Start(transInfo)
            So(err, ShouldEqual, nil)
        })

        Convey("trans exec succ when fail time is 2", func() {
            transInfo := &transdsl.TransInfo{
                AppInfo: &context.StubInfo{
                    FailTimes: 2,
                },
            }
            err := trans.Start(transInfo)
            So(err, ShouldEqual, nil)
        })

        Convey("trans exec fail when fail time is 3", func() {
            transInfo := &transdsl.TransInfo{
                AppInfo: &context.StubInfo{
                    FailTimes: 3,
                },
            }
            err := trans.Start(transInfo)
            So(err, ShouldNotEqual, nil)
        })

        Convey("trans exec fail when error string is panic", func() {
            transInfo := &transdsl.TransInfo{
                AppInfo: &context.StubInfo{
                    FailTimes: 1,
                    Y: -1,
                },
            }
            err := trans.Start(transInfo)
            So(err.Error(), ShouldEqual, "panic")
        })
    })
}

```

### wait
```go
import (
	"github.com/agiledragon/trans-dsl"
	"github.com/agiledragon/trans-dsl/test/context"
	"github.com/agiledragon/trans-dsl/test/context/action"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

var eventId = "assign cmd"

func newWaitTrans() *transdsl.Transaction {
	trans := &transdsl.Transaction{
		Fragments: []transdsl.Fragment{
			new(action.StubTransferMoney),
			&transdsl.Wait{
				EventId:  eventId,
				Timeout:  100,
				Fragment: new(action.StubAssignCmd),
			},
			new(action.StubAttachSomething),
			new(action.StubActivateSomething),
		},
	}
	return trans
}

type TransObj struct {
	trans     *transdsl.Transaction
	transInfo *transdsl.TransInfo
}

var transIds map[string]string
var transObjs map[string]TransObj
var key string

func handleEvent(eventId string, eventContent []byte) {
	transId := transIds[key]
	transObj := transObjs[transId]
	trans := transObj.trans
	transInfo := transObj.transInfo
	stubInfo := transInfo.AppInfo.(*context.StubInfo)
	stubInfo.EventContent = eventContent
	<-time.After(50 * time.Millisecond)
	trans.HandleEvent(eventId, transInfo)
}

func TestWaitTrans(t *testing.T) {
	trans := newWaitTrans()
	transIds = make(map[string]string)
	key = "business id"
	transId := "123456"
	transIds[key] = transId

	transObjs = make(map[string]TransObj)
	transInfo := &transdsl.TransInfo{
		Ch: make(chan struct{}),
		AppInfo: &context.StubInfo{
			TransId: "",
			X:       "info",
			Y:       1,
		},
	}
	transObjs[transId] = TransObj{trans: trans, transInfo: transInfo}
	Convey("TestWaitTrans", t, func() {

		Convey("wait succ", func() {
			go handleEvent(eventId, nil)
			err := trans.Start(transInfo)
			So(err, ShouldEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).Y, ShouldEqual, 8)
		})

		Convey("wait timeout", func() {
			transInfo := &transdsl.TransInfo{
				Ch: make(chan struct{}),
				AppInfo: &context.StubInfo{
					X: "info",
					Y: 1,
				},
			}
			err := trans.Start(transInfo)
			So(err.Error(), ShouldEqual, transdsl.ErrTimeout.Error())
		})
	})
}

```

### allof
```go
import (
	"github.com/agiledragon/trans-dsl"
	"github.com/agiledragon/trans-dsl/test/context"
	"github.com/agiledragon/trans-dsl/test/context/action"
	"github.com/agiledragon/trans-dsl/test/context/spec"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func newAllOfTrans() *transdsl.Transaction {
	trans := &transdsl.Transaction{
		Fragments: []transdsl.Fragment{
			&transdsl.Optional{
				Spec: &transdsl.AllOf{
					Specs: []transdsl.Specification{
						new(spec.IsAbcExist),
						new(spec.IsDefExist),
						new(spec.IsGhiExist),
					},
				},
				IfFrag: new(action.StubConnectAbc),
			},
			new(action.StubActivateSomething),
		},
	}
	return trans
}

func TestAllOfTrans(t *testing.T) {
	trans := newAllOfTrans()
	Convey("TestAllOfTrans", t, func() {

		Convey("all specs are true", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					Abc: "abc",
					Y: 1,
				},
			}
			err := trans.Start(transInfo)
			So(err, ShouldEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).Y, ShouldEqual, 2)
		})

		Convey("one of specs is false", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					Abc: "def",
					Y: 1,
				},
			}
			err := trans.Start(transInfo)
			So(err, ShouldEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).Y, ShouldEqual, 1)
		})
	})
}
```

### fail
```go
import (
    "github.com/agiledragon/trans-dsl"
    "github.com/agiledragon/trans-dsl/test/context"
    "github.com/agiledragon/trans-dsl/test/context/action"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
    "github.com/agiledragon/trans-dsl/test/context/spec"
    "errors"
)

var ErrResourceInsufficient = errors.New("resource insufficient")

func newFailTrans() *transdsl.Transaction {
    trans := &transdsl.Transaction{
        Fragments: []transdsl.Fragment{
            new(action.StubAttachSomething),
            &transdsl.Optional{
                Spec: new(spec.IsSomeResourceInsufficient),
                IfFrag: &transdsl.Fail{
                    ErrCode: ErrResourceInsufficient,
                },
            },
            new(action.StubActivateSomething),
        },
    }
    return trans
}

func TestFailTrans(t *testing.T) {
    trans := newFailTrans()
    Convey("TestFailTrans", t, func() {

        Convey("spec is true", func() {
            transInfo := &transdsl.TransInfo{
                AppInfo: &context.StubInfo{
                    X: "insufficient",
                    SpecialNum: 1,
                },
            }
            err := trans.Start(transInfo)
            So(err, ShouldEqual, ErrResourceInsufficient)
            So(transInfo.AppInfo.(*context.StubInfo).SpecialNum, ShouldEqual, 10)
        })

        Convey("spec is false", func() {
            transInfo := &transdsl.TransInfo{
                AppInfo: &context.StubInfo{
                    X: "sufficient",
                    SpecialNum: 1,
                },
            }
            err := trans.Start(transInfo)
            So(err, ShouldEqual, nil)
            So(transInfo.AppInfo.(*context.StubInfo).SpecialNum, ShouldEqual, 20)
        })
    })
}

```
