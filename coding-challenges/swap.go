package challenges

// Minimum number of swaps required to sort an array
// Given an array of n distinct elements, find the minimum number of swaps required to sort the array.

// Examples:

// Input: {4, 3, 2, 1}
// Output: 2
// Explanation: Swap index 0 with 3 and 1 with 2 to
//               form the sorted array {1, 2, 3, 4}.

// Input: {1, 5, 4, 3, 2}
// Output: 2

func SelectionSort(arr []int) (int, []int) {

	var swaps int
	var index int
	var zindex int

	arrSize := len(arr)
	for index = 0; index < arrSize-1; index++ {
		for zindex = index + 1; zindex < arrSize; zindex++ {
			if arr[index] < arr[zindex] {
				continue
			} else {
				auxValue := arr[index]
				arr[index] = arr[zindex]
				arr[zindex] = auxValue
				swaps++
			}
		}
	}

	return swaps, arr
}

func partition(vector *[]int, start int, end int) int {

	pivot := (*vector)[start]
	top := start

	for i := start + 1; i <= end; i++ {
		if pivot > (*vector)[i] {
			(*vector)[top] = (*vector)[i]
			(*vector)[i] = (*vector)[top+1]
			top++
		}
	}
	(*vector)[top] = pivot

	return top
}

func QuickSort(arr *[]int, start int, end int) {
	if start < end {
		mid := partition(arr, start, end)
		QuickSort(arr, start, mid)
		QuickSort(arr, mid+1, end)
	}
}
