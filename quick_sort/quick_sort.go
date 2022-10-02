package quick_sort

func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	quickSort(&nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums *[]int, low, high int) {
	if low >= high {
		return
	}
	pivot := (*nums)[high]
	i, j, direct := low, high, true
	for i < j {
		switch direct {
		case true:
			if (*nums)[i] < pivot {
				i++
				continue
			}
		case false:
			if (*nums)[j] >= pivot {
				j--
				continue
			}
		}
		(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
		direct = !direct
	}
	quickSort(nums, low, i-1)
	quickSort(nums, i+1, high)
}
