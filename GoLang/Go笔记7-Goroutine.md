> Goroutines可以被认为是轻量级的线程。与线程相比, 创建Goroutine的成本很小, 它就是一段代码, 一个函数入口, 以及在堆上为其分配的一个堆栈(初始大小为4K, 会随着程序的执行自动增长删除)。因此它非常廉价, Go应用程序可以并发运行数千个Goroutines。


Goroutines在线程上的优势:
1. 与线程相比, Goroutines非常便宜。它们只是堆栈大小的几个kb, 堆栈可以根据应用程序的需要增长和收缩, 而在线程的情况下, 堆栈大小必须指定并且是固定的。
2. Goroutines被多路复用到较少的OS线程。在一个程序中可能只有一个线程与数千个Goroutines。如果线程中的任何Goroutine都表示等待用户输入, 则会创建另一个OS线程, 剩下的Goroutines被转移到新的OS线程。所有这些都由运行时进行处理, 我们作为程序员从这些复杂的细节中抽象出来, 并得到了一个与并发工作相关的干净的API。
3. 当使用Goroutines访问共享内存时, 通过设计的通道可以防止竞态条件发生。通道可以被认为是Goroutines通信的管道。


## 主Gcoroutine
---
封装main函数的goroutine称为主goroutine。
1. 设定每一个goroutine所能申请的栈空间的最大尺寸。在32位的计算机系统中此最大尺寸为250MB, 而在64位的计算机系统中此尺寸为1GB. 如果有某个goroutine的栈空间尺寸大于这个限制, 那么运行时系统就会引发一个栈溢出(stack overflow)的运行时恐慌。随后, 这个go程序的运行也会终止。
2. 创建一个特殊的defer语句, 用于在主goroutine退出时做必要的善后处理。因为主goroutine也可能非正常的结束
3. 启动专用于在后台清扫内存垃圾的goroutine, 并设置GC可用的标识
4. 执行mian包中的init函数
5. 执行main函数
6. 执行完main函数后，它还会检查主goroutine是否引发了运行时恐慌，并进行必要的处理。最后主goroutine会结束自己以及当前进程的运行。

> 当新的Goroutine开始时, Goroutine调用立即返回。与函数不同, go不等待Goroutine执行结束。当Goroutine调用, 并且Goroutine的任何返回值被忽略之后, go立即执行到下一行代码。
main的Goroutine应该为其他的Goroutines执行。如果main的Goroutine终止了, 程序将被终止, 而其他Goroutine将不会运行。

## goroutine泄露
---
> 启动了一个 goroutine, 但没有符合预期的退出, 直到程序结束, 此goroutine才退出, 这种情况就是 goroutine 泄露。当 goroutine 泄露发生时, 该 goroutine 的栈一直被占用不能释放, goroutine 里的函数在堆上申请的空间也不能被垃圾回收器回收

goroutine泄漏的场景
```go
// 从 channel 里读 但是没有写
func leak() {
	ch := make(chan int)

	go func() {
		val := <-ch
		fmt.Println("We received a value:", val)
	}()
}

// goroutine进入死循环中 导致资源一直无法释放

func foo() {
	for {
		fmt.Println("fooo")
	}
}

// select操作在所有case上阻塞
// 向已满的 buffered channel 写 但是没有读
// 向 unbuffered channel 写 但是没有读
```