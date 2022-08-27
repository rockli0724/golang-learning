package main

import "fmt"

//查找
/*
二分查找的思路: 比如我们要查找的数是 findVal
1. arr 是一个有序数组，并且是从小到大排序
2. 先找到 中间的下标 middle = (leftIndex + rightIndex) / 2, 然后让 中间下标的值和 findVal 进行
比较
2.1 如果 arr[middle] > findVal , 就应该向 leftIndex ---- (middle - 1)
2.2 如果 arr[middle] < findVal , 就应该向 middel+1---- rightIndex
2.3 如果 arr[middle] == findVal ， 就找到
2.4 上面的 2.1 2.2 2.3 的逻辑会递归执行
3. 想一下，怎么样的情况下，就说明找不到[分析出退出递归的条件!!]
*/

func main() {
	//var arr = [6]int{1, 45, 46, 50, 100, 123}
	//binaryFind(&arr, 0, len(arr), 50)

	actions := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	batchSize := 3
	batches := make([][]int, 0, (len(actions)+batchSize-1)/batchSize)

	for batchSize < len(actions) {
		actions, batches = actions[batchSize:], append(batches, actions[0:batchSize:batchSize])

	}
	batches = append(batches, actions)
	fmt.Println(batches)

}

//1) 对有序数列进行二分查找{1,8, 10, 89, 1000, 1234}
func binaryFind(array *[6]int, leftIndex int, rightIndex int, findValue int) {
	//判断 leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex {
		return
	}
	//先找到 中间的下标
	middle := (leftIndex + rightIndex) / 2
	//中间下标值>目标值
	if (*array)[middle] > findValue {
		//说明我们要查找的数在 [leftIndex,middle)
		binaryFind(array, leftIndex, middle-1, findValue)
	} else if (*array)[middle] < findValue {
		//说明我们要查找的数在 (middle,rightIndex]
		binaryFind(array, middle+1, rightIndex, findValue)
	} else {
		fmt.Printf("找到了，下标=%v", middle)
	}
}
