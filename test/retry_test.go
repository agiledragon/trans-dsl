package test

import (
	"errors"
	"github.com/agiledragon/trans-dsl"
	"github.com/agiledragon/trans-dsl/test/context"
	"github.com/agiledragon/trans-dsl/test/context/action"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func newTryTrans() *transdsl.Transaction {
	trans := &transdsl.Transaction{
		Fragments: []transdsl.Fragment{
			&transdsl.Retry{
				MaxTimes: 3,
				TimeLen:  100,
				Fragment: new(action.StubConnectServer),
				Errs:     []error{errors.New("fatal"), errors.New("panic")},
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
					Y:         -1,
				},
			}
			err := trans.Start(transInfo)
			So(err.Error(), ShouldEqual, "panic")
		})
	})
}
