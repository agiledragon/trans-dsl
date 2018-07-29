package test

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

func newAnyOfTrans() *transdsl.Transaction {
	trans := &transdsl.Transaction{
		Fragments: []transdsl.Fragment{
			&transdsl.Optional{
				Spec: &transdsl.AnyOf{
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

func newNotTrans() *transdsl.Transaction {
	trans := &transdsl.Transaction{
		Fragments: []transdsl.Fragment{
			&transdsl.Optional{
				Spec: &transdsl.Not{
					Spec: new(spec.IsAbcExist),
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

func TestAnyOfTrans(t *testing.T) {
	trans := newAnyOfTrans()
	Convey("TestAnyOfTrans", t, func() {

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
			So(transInfo.AppInfo.(*context.StubInfo).Y, ShouldEqual, 2)
		})
	})
}

func TestNotTrans(t *testing.T) {
	trans := newNotTrans()
	Convey("TestNotTrans", t, func() {

		Convey("spec is true", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					Abc: "def",
					Y: 1,
				},
			}
			err := trans.Start(transInfo)
			So(err, ShouldEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).Y, ShouldEqual, 2)
		})
		
		Convey("spec is false", func() {
			transInfo := &transdsl.TransInfo{
				AppInfo: &context.StubInfo{
					Abc: "abc",
					Y: 1,
				},
			}
			err := trans.Start(transInfo)
			So(err, ShouldEqual, nil)
			So(transInfo.AppInfo.(*context.StubInfo).Y, ShouldEqual, 1)
		})
	})
}
