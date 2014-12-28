package deferr_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/marconi/deferr"
)

type FakeTodoRepo struct {
	deferr.TodoRepo
}

func (ftr *FakeTodoRepo) List() []*deferr.Todo {
	return nil
}

func (ftr *FakeTodoRepo) Push(t *deferr.Todo) error {
	return nil
}

func (ftr *FakeTodoRepo) Pop() (*deferr.Todo, error) {
	return nil, nil
}

func (ftr *FakeTodoRepo) Defer() error {
	return nil
}

func TestTodoManagerSpec(t *testing.T) {
	manager := deferr.NewTodoManager(&FakeTodoRepo{})

	Convey("testing todo manager", t, func() {
		Convey("should be able to push item", func() {
			t := &deferr.Todo{Name: "Wash clothes"}
			err := manager.Push(t)
			So(err, ShouldBeNil)
			So(t.Slug, ShouldNotBeBlank)
		})
	})
}
