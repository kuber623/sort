package merge_sort

import (
	"math"
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
				for i := 0; i < len(nums); i++ {
					nums[i] = rand.Intn(math.MaxInt64)
				}
				photo := make([]int, len(nums))
				for i, v := range nums {
					photo[i] = v
				}
				MergeSort(nums)
				sort.Ints(photo)
				So(nums, ShouldResemble, photo)
			}
		})
	})
}
