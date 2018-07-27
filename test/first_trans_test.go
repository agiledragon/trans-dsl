package test

import(
    "github.com/agiledragon/trans-dsl"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
    "github.com/agiledragon/trans-dsl/test/context/action"
    "github.com/agiledragon/trans-dsl/test/context"
)

func newSuccTrans() *transdsl.Transaction {
    trans := &transdsl.Transaction {
        Fragments: []transdsl.Fragment{
            new(action.StubAttachSomething),
            new(action.StubActivateSomething),
        },
    }
    return trans
}

func TestFirstTrans(t *testing.T) {

    Convey("first trrans test", t, func() {

        Convey("trans exec succ", func() {
            transInfo := &transdsl.TransInfo{
                AppInfo: &context.StubInfo{
                    X: "first",
                    Y: 1,
                },
            }
            trans := newSuccTrans()
            err := trans.Exec(transInfo)
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
            trans := newSuccTrans()
            err := trans.Exec(transInfo)
            So(err, ShouldNotEqual, nil)
            So(transInfo.AppInfo.(*context.StubInfo).Y, ShouldEqual, 0)
        })
    })
}
