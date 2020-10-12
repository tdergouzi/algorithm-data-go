package search

// 二分查找
// -1 表示未找到该值
func Binary(arr []int64, n int64) int {
	if len(arr) <= 1 {
		return 0
	}
	//return binaryWithRecursion(arr, n, 0, len(arr)-1)
	//return binaryWithoutRecursion(arr, n)
	return binaryForFirst(arr, n)
}

// 非递归实现
func binaryWithoutRecursion(arr []int64, n int64) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == n {
			return mid
		} else if arr[mid] < n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// 递归实现
func binaryWithRecursion(arr []int64, n int64, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high - low)
	if arr[mid] == n {
		return mid
	} else if arr[mid] < n {
		low = mid + 1
		return binaryWithRecursion(arr, n, low, high)
	} else {
		high = mid - 1
		return binaryWithRecursion(arr, n, low, high)
	}
}

// 查找第一个给定值的元素
func binaryForFirst(arr []int64, n int64) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == n {
			for arr[mid-1] == arr[mid] {
				mid--
			}
			return mid
		} else if arr[mid] < n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return len(arr)
}
