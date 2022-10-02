package quick_sort

import (
	"math"
	"math/rand"
	"sort"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQuickSort(t *testing.T) {
	Convey("quick sort", t, func() {
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
				QuickSort(nums)
				sort.Ints(photo)
				So(nums, ShouldResemble, photo)
			}
		})
	})
}
