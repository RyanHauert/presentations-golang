package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	Convey("Given a pool of five workers", t, func() {
		pool := NewWorkerPool(5)

		Convey("When given a batch of work", func() {
			for i := 0; i < 20; i++ {
				pool.Queue <- i
			}
			close(pool.Queue)

			Convey("Work will be balanced", func() {
				counts := make(map[string]int)
				for res := range pool.Results {
					counts[res.Worker]++
				}

				for _, v := range counts {
					So(v, ShouldEqual, 4)
				}
			})
		})
	})
}
