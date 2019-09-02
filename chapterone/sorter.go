package chapterone

import "fmt"

//MergeSort ...
func MergeSort(array []int) {

	fmt.Println("The length of array -> ", len(array))
	mergeSorterRecursive(array, 0, len(array)-1)

}

func mergeSorterRecursive(array []int, leftIndex int, rightIndex int) {
	if leftIndex >= rightIndex {
		return
	}

	middle := (leftIndex + rightIndex) / 2
	mergeSorterRecursive(array, leftIndex, middle)
	mergeSorterRecursive(array, middle+1, rightIndex)

	merge(array, leftIndex, middle, rightIndex)
}

func merge(array []int, leftIndex int, middle int, rightIndex int) {
	temp := make([]int, rightIndex-leftIndex+1)

	leftPtr := leftIndex
	rightPtr := middle + 1
	index := 0
	for leftPtr <= middle && rightPtr <= rightIndex {
		if array[leftPtr] < array[rightPtr] {
			temp[index] = array[leftPtr]
			index++
			leftPtr++
		} else if array[leftPtr] > array[rightPtr] {
			temp[index] = array[rightPtr]
			rightPtr++
			index++
		} else {
			temp[index] = array[leftPtr]
			index++
			temp[index] = array[rightPtr]
			index++
			leftPtr++
			rightPtr++
		}
	}

	for leftPtr <= middle {
		temp[index] = array[leftPtr]
		index++
		leftPtr++
	}

	for rightPtr <= rightIndex {
		temp[index] = array[rightPtr]
		index++
		rightPtr++
	}

	for i, k := 0, leftIndex; i < cap(temp); {
		array[k] = temp[i]
		k++
		i++
	}
}
