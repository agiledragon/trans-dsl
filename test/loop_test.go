package test

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
