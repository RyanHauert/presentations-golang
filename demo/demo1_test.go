package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestWorkersComplete(t *testing.T) {
	Convey("When a worker does work", t, func() {
		worker := make(chan string)
		go DoWork(worker)

		select {
		case res := <-worker:
			Convey("The worker should report it finished", func() {
				So(res, ShouldEqual, "Finished!")
			})
		case <-time.After(1 * time.Second):
			t.Error("worker did not complete")
		}
	})
}
