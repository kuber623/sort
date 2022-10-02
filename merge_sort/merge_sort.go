package merge_sort

func MergeSort(nums []int) {
	mergeSort(&nums, 0, len(nums)-1)
}

func mergeSort(nums *[]int, low, high int) {
	if low >= high {
		return
	}
	// divide
	mid := (low + high) / 2
	mergeSort(nums, low, mid)
	mergeSort(nums, mid+1, high)
	// merge and sort
	merge(nums, low, high)
}

func merge(nums *[]int, low, high int) {
	if low >= high {
		return
	}
	tmp := make([]int, 0, high-low+1)
	mid := (low + high) / 2
	i, j := low, mid+1
	for i <= mid && j <= high {
		if (*nums)[i] <= (*nums)[j] {
			tmp = append(tmp, (*nums)[i])
			i++
		} else {
			tmp = append(tmp, (*nums)[j])
			j++
		}
	}
	// place remaining elements
	switch {
	case i <= mid:
		tmp = append(tmp, (*nums)[i:mid+1]...)
	case j <= high:
		tmp = append(tmp, (*nums)[j:high+1]...)
	}
	// write back sorted elements
	for i := range tmp {
		(*nums)[low+i] = tmp[i]
	}
}
