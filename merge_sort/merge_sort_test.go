package merge_sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeSort(t *testing.T) {
	Convey("merge sort", t, func() {
		Convey("normal test", func() {
			for i := 0; i < 20; i++ {
				nums := make([]int, 1000)
				rand.Seed(time.Now().UnixNano())
				for i := 0; i < 100; i++ {
					nums[i] = rand.Intn(100)
				}
				MergeSort(nums)
				So(sort.IntsAreSorted(nums), ShouldResemble, true)
			}
		})
	})
}
