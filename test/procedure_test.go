package test

import (
    "github.com/agiledragon/trans-dsl"
    "github.com/agiledragon/trans-dsl/test/context"
    "github.com/agiledragon/trans-dsl/test/context/action"
    "github.com/agiledragon/trans-dsl/test/context/spec"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func newApplyProcedure() transdsl.Fragment {
    procedure := &transdsl.Procedure{
        Fragments: []transdsl.Fragment{
            new(action.StubGetSomething),
            new(action.StubModifySomething),
            new(action.StubApplySomething),
        },
    }
    return procedure
}

func newProcedureTrans() *transdsl.Transaction {
    trans := &transdsl.Transaction{
        Fragments: []transdsl.Fragment{
            &transdsl.Optional{
                Spec:   new(spec.IsAbcExist),
                IfFrag: newApplyProcedure(),
            },
            new(action.StubActivateSomething),
        },
    }
    return trans
}

func TestProcedureTrans(t *testing.T) {
    trans := newProcedureTrans()
    Convey("TestProcedureTrans", t, func() {

        Convey("trans exec succ when spec is true", func() {
            transInfo := &transdsl.TransInfo{
                AppInfo: &context.StubInfo{
                    Abc: "abc",
                    P1: 1,
                    P2: 2,
                    P3: 3,
                },
            }
            err := trans.Start(transInfo)
            So(err, ShouldEqual, nil)
            stubInfo := transInfo.AppInfo.(*context.StubInfo)
            So(stubInfo.P1, ShouldEqual, 21)
            So(stubInfo.P2, ShouldEqual, 22)
            So(stubInfo.P3, ShouldEqual, 23)
        })

        Convey("trans exec succ when spec is false", func() {
            transInfo := &transdsl.TransInfo{
                AppInfo: &context.StubInfo{
                    Abc: "def",
                    P1: 1,
                    P2: 2,
                    P3: 3,
                },
            }
            err := trans.Start(transInfo)
            So(err, ShouldEqual, nil)
            stubInfo := transInfo.AppInfo.(*context.StubInfo)
            So(stubInfo.P1, ShouldEqual, 1)
            So(stubInfo.P2, ShouldEqual, 2)
            So(stubInfo.P3, ShouldEqual, 3)
        })

        Convey("iffrag rollback", func() {
            transInfo := &transdsl.TransInfo{
                AppInfo: &context.StubInfo{
                    X: "test",
                    Abc: "abc",
                    P1: 1,
                    P2: 2,
                    P3: 3,
                },
            }
            err := trans.Start(transInfo)
            So(err, ShouldNotEqual, nil)
            stubInfo := transInfo.AppInfo.(*context.StubInfo)
            So(stubInfo.P1, ShouldEqual, 0)
            So(stubInfo.P2, ShouldEqual, 0)
            So(stubInfo.P3, ShouldEqual, 0)
        })
    })
}
