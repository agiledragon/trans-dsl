package test

import (
	"errors"
	"github.com/agiledragon/trans-dsl"
	"github.com/agiledragon/trans-dsl/test/context"
	"github.com/agiledragon/trans-dsl/test/context/action"
	"github.com/agiledragon/trans-dsl/test/context/spec"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func newSpecialTrans() *transdsl.Transaction {
	trans := &transdsl.Transaction{
		Fragments: []transdsl.Fragment{
			new(action.StubAttachSomething),
			&transdsl.Optional{
				Spec:   new(spec.IsSpecialConditionSatisfied),
				IfFrag: new(transdsl.Succ),
			},
			new(action.StubActivateSomething),
		},
	}
	return trans
}

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

func TestSpecialTrans(t *testing.T) {
	trans := newSpecialTrans()
	Convey("TestSpecialTrans", t, func() {

		Convey("spec is true", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					X:          "special",
					SpecialNum: 1,
				},
			}
			err := trans.Start(transInfo)
			So(err, ShouldEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).SpecialNum, ShouldEqual, 10)
		})

		Convey("spec is false", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					X:          "normal",
					SpecialNum: 1,
				},
			}
			err := trans.Start(transInfo)
			So(err, ShouldEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).SpecialNum, ShouldEqual, 20)
		})
	})
}

func TestFailTrans(t *testing.T) {
	trans := newFailTrans()
	Convey("TestFailTrans", t, func() {

		Convey("spec is true", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					X:          "insufficient",
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
					X:          "sufficient",
					SpecialNum: 1,
				},
			}
			err := trans.Start(transInfo)
			So(err, ShouldEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).SpecialNum, ShouldEqual, 20)
		})
	})
}
