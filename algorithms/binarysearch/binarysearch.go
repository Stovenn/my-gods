package binarysearch

func BS(arr []int, key int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] > key {
			hi = mid - 1
		} else if arr[mid] < key {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func BSR(arr []int, key int) int {
	for len(arr) > 1 {
		mid := arr[0] + (arr[len(arr)-1]-arr[0])/2
		if arr[mid] > key {
			return BSR(arr[:mid], key)
		} else if arr[mid] < key {
			return BSR(arr[mid:], key)
		} else {
			return mid
		}
	}
	return -1
}
