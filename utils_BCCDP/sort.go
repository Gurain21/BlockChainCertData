package utils_BCCDP

import (
	"fmt"
)

func MaoPaoSort(arr []int, left, right int) {
	if left >= right {
		fmt.Println("请再次确认您传入的下标值无误！")
	}

	for i := left; i < right; i++ {
		for j := left; j < right-left-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

	}
}
func QuickSort(arr []int, left, right int) {
	if left > right {
		fmt.Println("请再次确认您传入的下标值无误！")
	}
	//设置基准值
	base := arr[left]
	i := left
	j := right
	for i != j {
		for arr[j] >= base && i < j {
			j--
		}
		for arr[i] <= base && i < j {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], base = base, arr[i]
	QuickSort(arr, left, i-1)
	QuickSort(arr, j+1, right)

}
