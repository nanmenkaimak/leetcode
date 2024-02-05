package valid_mountain_array

func validMountainArray(arr []int) bool {
	maxNum, idx := 0, 0

	for i := 0; i < len(arr); i++ {
		if maxNum < arr[i] {
			maxNum = arr[i]
			idx = i
		}
	}

	if idx == 0 || idx == len(arr)-1 {
		return false
	}

	for i := 0; i < idx; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	for i := len(arr) - 1; i > idx; i-- {
		if arr[i] >= arr[i-1] {
			return false
		}
	}
	return true
}
