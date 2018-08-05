package test

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
