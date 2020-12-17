```go
package main

import "fmt"
import "math/rand"

// 随机数据
func randomArray(n int) []int {
	slice := rand.Perm(100)[:n]
	return slice
}

// 快速排序
func quickSort(arr []int) []int {
	median := arr[rand.Intn(len(arr))]

	lowPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			lowPart = append(lowPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			highPart = append(highPart, item)
		}
	}

	lowPart = quickSort(lowPart)
	highPart = quickSort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}

// 插入排序
func insertSort(arr []int){
  for i :=0; i < len(arr); i++{
    
    tmp := arr[i]
    j := i - 1
    for ; j > 0 && arr[j] > tmp; j--{
      arr[j + 1] = arr[j]
    }
  }
}

func main() {
	arr := randomArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("")
	fmt.Println("quick sorted array is: ", quickSort(arr))
}

```