## air 开发时改动自动重启
---
```shell
// air 修改自动重新编译启动
go get -u github.com/cosmtrek/air
```

## container/heap(最小堆)
---
堆(heap)通常是一个可以被看做一棵树的数组对象。堆总是满足下列性质：
1. 堆中某个节点的值总是不大于或不小于其父节点的值；
2. 堆总是一棵完全二叉树。
3. 将根节点最大的堆叫做最大堆或大根堆，根节点最小的堆叫做最小堆或小根堆。

想要使用 container/heap,目标类型需要包含这些方法: `Len` `Less` `Swap` `Push` `Pop`

具体使用可以参考 LeetCode TOP K相关问题解法

## container/list(双向链表)
---
```go
func main() {
	list := list.New()

	e1 := list.PushBack(4)
	e2 := list.PushFront(1)
	list.InsertBefore(3, e1)
	list.InsertAfter(2, e2)

	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
```