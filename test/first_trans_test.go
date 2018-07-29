package test

import (
	"github.com/agiledragon/trans-dsl"
	"github.com/agiledragon/trans-dsl/test/context"
	"github.com/agiledragon/trans-dsl/test/context/action"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func newFirstTrans() *transdsl.Transaction {
	trans := &transdsl.Transaction{
		Fragments: []transdsl.Fragment{
			new(action.StubAttachSomething),
			new(action.StubActivateSomething),
		},
	}
	return trans
}

func TestFirstTrans(t *testing.T) {
	trans := newFirstTrans()
	Convey("TestFirstTrans", t, func() {

		Convey("trans exec succ", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					X: "first",
					Y: 1,
				},
			}
			err := trans.Start(transInfo)
			So(err, ShouldEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).Y, ShouldEqual, 8)
		})

		Convey("trans exec failed", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					X: "test",
					Y: 1,
				},
			}
			err := trans.Start(transInfo)
			So(err, ShouldNotEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).Y, ShouldEqual, 0)
		})
	})
}
