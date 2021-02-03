## 并发模型
---
无论语言层面何种并发模型, 到了操作系统层面, 一定是以线程的形态存在的。而操作系统根据资源访问权限的不同, 体系架构可分为用户空间和内核空间. 内核空间主要操作访问CPU资源 I/O资源、内存资源等硬件资源, 为上层应用程序提供最基本的基础资源, 用户空间呢就是上层应用程序的固定活动空间, 用户空间不可以直接访问资源, 必须通过 "系统调用", "库函数" "Shell脚本" 来调用内核空间提供的资源。
一般我们编写运行的线程是用户态线程, 与之对应的是系统本身内核态线程KSE(kernel support for user threads)KSE

Go并发编程模型在底层是由操作系统所提供的线程库支撑的.
> 线程可以视为进程中的控制流。一个进程至少会包含一个线程, 因为其中至少会有一个控制流持续运行。一个进程的第一个线程会随着这个进程的启动而创建(主线程)。一个进程也可以包含多个线程。这些线程都是由当前进程中已存在的线程创建出来的, 创建的方法就是调用系统调用(pthread create函数)。拥有多个线程的进程可以并发执行多个任务, 即使某个或某些任务被阻塞, 也不会影响其他任务正常执行, 这可以大大改善程序的响应时间和吞吐量。线程不可能独立于进程存在。它的生命周期不可能逾越其所属进程的生命周期。

线程的实现模型主要有3个，分别是:用户级线程模型、内核级线程模型和两级线程模型。它们之间最大的差异就在于线程与内核调度实体( Kernel Scheduling Entity,简称KSE)之间的对应关系上。顾名思义，内核调度实体就是可以被内核的调度器调度的对象。在很多文献和书中，它也称为内核级线程，是操作系统内核的最小调度单元。

### 1. 内核级线程模型
---
用户线程与KSE是1对1关系(1:1)。大部分编程语言的线程库(如linux的pthread，Java的java.lang.Thread，C++11的std::thread等等)都是对操作系统的线程(内核级线程)的一层封装, 每个线程与一个不同的KSE静态关联, 其调度完全由OS调度器来做。这种方式实现简单, 直接借助OS提供的线程能力, 并且不同用户线程之间一般也不会相互影响。但创建 销毁以及多个线程之间的上下文切换等操作都是直接由OS层面亲自来做, 在需要使用大量线程的场景下对OS的性能影响会很大。

### 2. 用户级线程模型
---
用户线程与KSE是多对1关系(M:1), 这种线程的创建 销毁以及多个线程之间的协调等操作都是由用户自己实现的线程库来负责, 一个进程中所有创建的线程都与同一个KSE在运行时动态关联。现在有许多语言实现的 协程 基本上都属于这种方式。这种实现方式相比内核级线程更加轻量级, 对系统资源的消耗会小很多, 因此可以创建的数量与上下文切换所花费的代价也会小得多。但该模型有个致命的缺点, 如果我们在某个用户线程上调用阻塞式系统调用(如用阻塞方式read网络IO), 那么一旦KSE因阻塞被内核调度出CPU的话, 剩下的所有对应的用户线程全都会变为阻塞状态(整个进程挂起)。所以这些语言的协程库会把自己一些阻塞的操作重新封装为完全的非阻塞形式, 然后在以前要阻塞的点上, 主动让出自己, 并通过某种方式通知或唤醒其他待执行的用户线程在该KSE上运行, 从而避免了内核调度器由于KSE阻塞而做上下文切换, 这样整个进程也不会被阻塞了。

+ 优点: 这种模型的好处是线程上下文切换都发生在用户空间, 避免的模态切换, 从而对于性能有积极的影响。
+ 缺点: 所有的线程基于一个内核调度实体即内核线程, 这意味着只有一个处理器可以被利用, 在多处理器环境下这是不能够被接受的, 本质上, 用户线程只解决了并发问题, 但是没有解决并行问题。如果线程因为 I/O 操作陷入了内核态, 内核态线程阻塞等待 I/O 数据, 则所有的线程都将会被阻塞, 用户空间也可以使用非阻塞而 I/O，但是不能避免性能及复杂度问题。

### 3. 两级线程模型
---
用户线程与KSE是多对多关系(M:N), 这种实现综合了前两种模型的优点, 为一个进程中创建多个KSE, 并且线程可以与不同的KSE在运行时进行动态关联, 当某个KSE由于其上工作的线程的阻塞操作被内核调度出CPU时, 当前与其关联的其余用户线程可以重新与其他KSE建立关联关系。当然这种动态关联机制的实现很复杂, 也需要用户自己去实现。Go语言中的并发就是使用的这种实现方式, Go为了实现该模型自己实现了一个运行时调度器来负责Go中的"线程"与KSE的动态关联。此模型有时也被称为 混合型线程模型, 即用户调度器实现用户线程到KSE的 "调度", 内核调度器实现KSE到CPU上的调度。

## Go并发调度: G-P-M模型
---
在操作系统提供的内核线程之上, Go搭建了一个特有的两级线程模型。goroutine机制实现了 M:N 的线程模型, goroutine机制是协程(coroutine)的一种实现, golang内置的调度器, 可以让多核CPU中每个CPU执行一个协程。

###Go语言中支撑整个scheduler实现的主要有4个重要结构, 分别是 M G P Sched

1. Sched结构就是调度器, 它维护有存储 M 和 G 的队列以及调度器的一些状态信息等。
2. M结构是 Machine, 系统线程, 它由操作系统管理的。goroutine就是跑在M之上的, M是一个很大的结构, 里面维护小对象内存cache(mcach), 当前执行的goroutine, 随机数发生器等.
3. P结构是Processor 处理器, 它的主要用途就是用来执行goroutine的, 它维护了一个goroutine队列: runqueue。Processor是让我们从N:1调度到M:N调度的重要部分。
4. G是goroutine实现的核心结构, 它包含了栈 指令指针以及其他对调度goroutine很重要的信息(其阻塞的channel等)

> 尽管 Go 编译器产生的是本地可执行代码, 这些代码仍旧运行在 Go 的 runtime上, 类似JAVA的虚拟机, 它负责管理包括内存分配, 垃圾回收, 栈处理, goroutine, channel, 切片, map, 反射等等。G P和M都是Go语言运行时系统（其中包括内存分配器, 并发调度器, 垃圾收集器等组件, 可以想象为Java中的JVM）抽象出来概念和数据结构对象

1. G: Goroutine的简称, 是对一个要并发执行的任务的封装, 也可以称作用户态线程。属于用户级资源, 对OS透明, 轻量级, 可以大量创建，上下文切换成本低
2. M: Machine的简称, 在linux平台上是用clone系统调用创建的，是利用系统调用创建出来的OS线程实体。M的作用就是执行G中包装的并发任务。Go运行时系统中的调度器的主要职责就是将G公平合理的安排到多个M上去执行。其属于OS资源，可创建的数量上也受限了OS, 通常情况下G的数量都多于活跃的M的。
3. P: Processor的简称, 逻辑处理器, 主要作用是管理G对象(每个P都有一个G队列), 并为G在M上的运行提供本地化资源。
4. 单核的场景下, 所有goroutine运行在同一个M系统线程中, 每一个 M 系统线程维护一个Processor, 一个Processor中只有一个goroutine, 其他 goroutine 在 runqueue 中等待。一个goroutine运行完自己的时间片后, 让出上下文, 回到runqueue中。 多核处理器的场景下，为了运行goroutines，每个M系统线程会持有一个Processor。
5. runqueue为空时: 没有goroutine可以调度。它会从另外一个上下文偷取一半的goroutine。
6. 全局队列(Global Queue): 存放等待运行的 G。

### Goroutine
---
#### 结构
---
G，取 goroutine 的首字母，主要保存 goroutine 的一些状态信息以及 CPU 的一些寄存器的值，例如 IP 寄存器，以便在轮到本 goroutine 执行时，CPU 知道要从哪一条指令处开始执行。
1. 当 goroutine 被调离 CPU 时，调度器负责把 CPU 寄存器的值保存在 g 对象的成员变量之中。
2. 当 goroutine 被调度起来运行时，调度器又负责把 g 对象的成员变量所保存的寄存器值恢复到 CPU 的寄存器。
```go
// src/runtime/runtime2.go#406
type g struct {
	stack       stack
	stackguard0 uintptr
	
	preempt       bool // 抢占信号
	preemptStop   bool // 抢占时将状态修改成 `_Gpreempted`
	preemptShrink bool // 在同步安全点收缩栈
	
	m              *m
	sched          gobuf
	atomicstatus   uint32
	goid           int64
}

// 描述栈的数据结构，栈的范围：[lo, hi)
type stack struct {
  // 栈顶，低地址
  lo uintptr
  // 栈低，高地址
  hi uintptr
}

// 这些内容会在调度器保存或者恢复上下文的时候用到，其中的栈指针和程序计数器会用来存储或者恢复寄存器中的值，改变程序即将执行的代码
type gobuf struct {
	// 存储 rsp 寄存器的值
	sp uintptr
	// 存储 rip 寄存器的值
	pc uintptr
	// 指向 goroutine
	g    guintptr
	ctxt unsafe.Pointer // this has to be a pointer so that gc scans it
	// 保存系统调用的返回值
	ret sys.Uintreg
	lr  uintptr
	bp  uintptr // for GOEXPERIMENT=framepointer
}
```
1. stack: 当前 Goroutine 的栈内存范围 [stack.lo, stack.hi)
2. stackguard0: 可以用于调度器抢占式调度
3. 每一个 Goroutine 上都持有两个分别存储 defer 和 panic 对应结构体的链表
4. m: 当前 Goroutine 占用的线程, 可能为空
5. atomicstatus: Goroutine 的状态
6. sched: 存储 Goroutine 的调度相关的数据
7. goid: Goroutine 的 ID，该字段对开发者不可见
8. sp: 栈指针
9. pc: 程序计数器
10. g: 持有 runtime.gobuf 的 Goroutine
11. ret: 系统调用的返回值；

#### 状态
---
结构体 runtime.g 的 atomicstatus 字段存储了当前 Goroutine 的状态。除了几个已经不被使用的以及与 GC 相关的状态之外，Goroutine 可能处于以下 9 种状态:
```go
var gStatusStrings = [...]string{
	_Gidle:      "idle",
	_Grunnable:  "runnable",
	_Grunning:   "running",
	_Gsyscall:   "syscall",
	_Gwaiting:   "waiting",
	_Gdead:      "dead",
	_Gcopystack: "copystack",
	_Gpreempted: "preempted",
}
```
+ _Gidle: 刚刚被分配并且还没有被初始化
+ _Grunnable: 没有执行代码，没有栈的所有权，存储在运行队列中
+ _Grunning: 可以执行代码，拥有栈的所有权，被赋予了内核线程 M 和处理器 P
+ _Gsyscall: 正在执行系统调用，拥有栈的所有权，没有执行用户代码，被赋予了内核线程 M 但是不在运行队列上
+ _Gwaiting: 由于运行时而被阻塞，没有执行用户代码并且不在运行队列上，但是可能存在于 Channel 的等待队列上
+ _Gdead: 没有被使用，没有执行代码，可能有分配的栈
+ _Gcopystack: 栈正在被拷贝，没有执行代码，不在运行队列上
+ _Gpreempted: 由于抢占而被阻塞，没有执行用户代码并且不在运行队列上，等待唤醒
+ _Gscan: GC 正在扫描栈空间，没有执行代码，可以与其他状态同时存在


虽然 Goroutine 在运行时中定义的状态非常多而且复杂，但是我们可以将这些不同的状态聚合成三种: `等待中` `可运行` `运行中` 运行期间会在这三种状态来回切换:

1. 等待中: Goroutine 正在等待某些条件满足，例如: 系统调用结束等，包括 _Gwaiting、_Gsyscall 和 _Gpreempted 几个状态
2. 可运行: Goroutine 已经准备就绪，可以在线程运行，如果当前程序中有非常多的 Goroutine，每个 Goroutine 就可能会等待更多的时间，即 _Grunnable
3. 运行中: Goroutine 正在某个线程上运行，即 _Grunning
![IMAGE](resources/B403448F4573CA50C1A18F3E3BC733DF.jpg =739x265)



### Machine
---
#### 结构
---
> Go 语言并发模型中的 M 是操作系统线程。调度器最多可以创建 10000 个线程，但是其中大多数的线程都不会执行用户代码（可能陷入系统调用），最多只会有 GOMAXPROCS 个活跃线程能够正常运行。它保存了 M 自身使用的栈信息，当前正在 M 上执行的 G 信息，与之绑定的 P 信等等。当 M 没有工作可做的时候，在它休眠前，会 `自旋` 地来找工作: 检查全局队列，查看 network poller，试图执行 gc 任务，或者 "偷" 工作。

```go
type m struct {
	g0   *g
	curg *g
	
	p             puintptr
	nextp         puintptr
	oldp          puintptr
}
```
1. g0: 是持有调度栈的 Goroutine
2. curg 是在当前线程上运行的用户 Goroutine，这也是操作系统线程唯一关心的两个 Goroutine。
3. runtime.m 结构体中还存在三个与处理器相关的字段，它们分别表示正在运行代码的处理器 p，暂存的处理器 nextp 和执行系统调用之前使用线程的处理器 oldp



### Processor
---
> 调度器中的处理器 P 是线程和 Goroutine 的中间层，它能提供线程需要的上下文环境，也会负责调度线程上的等待队列，通过处理器 P 的调度，每一个内核线程都能够执行多个 Goroutine，它能在 Goroutine 进行一些 I/O 操作时及时让出计算资源，提高线程的利用率。因为调度器在启动时就会创建 GOMAXPROCS 个处理器，所以 Go 语言程序的处理器数量一定会等于 GOMAXPROCS，这些处理器会绑定到不同的内核线程上。

```go
type p struct {
 // 指向绑定的 m，如果 p 是 idle 的话，那这个指针是 nil
	m           muintptr
  // 本地可运行的队列，不用通过锁即可访问
	runqhead uint32 // 队列头
	runqtail uint32 // 队列尾
	// 使用数组实现的循环队列
	runq     [256]guintptr
	runnext guintptr
}
```

#### 状态
---
+ _Pidle: 处理器没有运行用户代码或者调度器，被空闲队列或者改变其状态的结构持有，运行队列为空
+ _Prunning: 被线程 M 持有，并且正在执行用户代码或者调度器
+ _Psyscall: 没有执行用户代码，当前线程陷入系统调用
+ _Pgcstop:	被线程 M 持有，当前处理器由于垃圾回收被停止
+ _Pdead: 当前处理器已经不被使用


### Scheduler
---
调度器: 所有 Goroutine 被调度的核心，存放了调度器持有的全局资源

1. 管理了能够将 G 和 M 进行绑定的 M 队列
2. 管理了空闲的 P 链表（队列）
3. 管理了 G 的全局队列
4. 管理了可被复用的 G 的全局缓存
5. 管理了 defer 池

```go
type schedt struct {
	lock mutex

	pidle      puintptr // 空闲 p 链表
	npidle     uint32   // 空闲 p 数量
	nmspinning uint32   // 自旋状态的 M 的数量
	runq       gQueue   // 全局 runnable G 队列
	runqsize   int32
	gFree      struct { // 有效 dead G 的全局缓存.
		lock    mutex
		stack   gList // 包含栈的 Gs
		noStack gList // 没有栈的 Gs
		n       int32
	}
	sudoglock  mutex // sudog 结构的集中缓存
	sudogcache *sudog
	deferlock  mutex // 不同大小的有效的 defer 结构的池
	deferpool  [5]*_defer
}
```

```go
// src/runtime/proc.go#534
func schedinit() {
	_g_ := getg()
	...

	sched.maxmcount = 10000

	...
	sched.lastpoll = uint64(nanotime())
	procs := ncpu
	if n, ok := atoi32(gogetenv("GOMAXPROCS")); ok && n > 0 {
		procs = n
	}
	if procresize(procs) != nil {
		throw("unknown runnable goroutine during bootstrap")
	}
}
```

1. 使用 make([]p, nprocs) 初始化全局变量 allp，即 allp = make([]p, nprocs)
2. 循环创建并初始化 nprocs 个 p 结构体对象并依次保存在 allp 切片之中
3. 把 m0 和 allp[0] 绑定在一起，即 m0.p = allp[0]，allp[0].m = m0
4. 把除了 allp[0] 之外的所有 p 放入到全局变量 sched 的 pidle 空闲队列之中


### 创建Goroutine
```go
// src/runtime/proc.go#3550
func newproc(siz int32, fn *funcval) {
	argp := add(unsafe.Pointer(&fn), sys.PtrSize)
	gp := getg()
	pc := getcallerpc()
	systemstack(func() {
		newg := newproc1(fn, argp, siz, gp, pc)

		_p_ := getg().m.p.ptr()
		runqput(_p_, newg, true)

		if mainStarted {
			wakep()
		}
	})
}
```
M/P/G 彼此的初始化顺序遵循: mcommoninit procresize newproc，他们分别负责初始化 M 资源池（allm）、P 资源池（allp）、G 的运行现场（g.sched）以及调度队列(p.runq)
![IMAGE](resources/B40C8736CA4B6FF617AAEAF21CB5413C.jpg =569x256)

![IMAGE](resources/538AE86049CC4175631EB5EF48D50A17.jpg =743x726)
![IMAGE](resources/0F80B5BAB5A978D456634D30C43DD754.jpg =767x483)


### 从两级线程模型来看, 似乎并不需要P的参与, 有G和M就可以了, 那为什么要加入P呢？
---
如果没有P会带来一些问题:
1. 不同的G在不同的M上并发运行时可能都需向系统申请资源(如堆内存), 由于资源是全局的, 将会由于资源竞争造成很多系统性能损耗, 让P去管理G对象, M要想运行G必须先与一个P绑定, 然后才能运行该P管理的G。这样带来的好处是，我们可以在P对象中预先申请一些系统资源, G需要的时候先向自己的本地P申请, 如果不够用或没有再向全局申请, 而且从全局拿的时候会多拿一部分, 以供后面高效的使用
2. P解耦了G和M对象, 这样即使M由于被其上正在运行的G阻塞住, 其余与该M关联的G也可以随着P一起迁移到别的活跃的M上继续运行, 从而让G总能及时找到M并运行自己, 从而提高系统的并发能力。


### 如何在一个多核心系统上尽量合理分配G到多个M上运行, 充分利用多核, 提高并发能力呢？
---
如果我们在一个Goroutine中通过go关键字创建了大量G, 这些G虽然暂时会被放在同一个队列, 如果这时还有空闲P(系统内P的数量默认等于系统cpu核心数), Go运行时系统始终能保证至少有一个(通常也只有一个)活跃的M与空闲P绑定去各种G队列去寻找可运行的G任务, 该种M称为自旋的M。一般寻找顺序为: 自己绑定的P的队列 -> 全局队列 -> 其他P队列。如果自己P队列找到就拿出来开始运行, 否则去全局队列看看, 由于全局队列需要锁保护, 如果里面有很多任务, 会转移一批到本地P队列中, 避免每次都去竞争锁。如果全局队列还是没有，直接从其他P队列拿一半任务回来。

###  如果某个M在执行G的过程中被G中的系统调用阻塞了, 发生什么？
---
这个M将会被内核调度器调度出CPU并处于阻塞状态, 与该M关联的其他G就没有办法继续执行了, 但Go运行时系统的一个监控线程(sysmon线程)能探测到这样的M, 并把与该M绑定的P剥离, 寻找其他空闲或新建M接管该P, 然后继续运行其中的G, 然后等到该M从阻塞状态恢复, 需要重新找一个空闲P来继续执行原来的G, 如果这时系统正好没有空闲的P, 就把原来的G放到全局队列当中, 等待其他M+P组合发掘并执行。

### 如果某一个G在M运行时间过长, 有没有办法做抢占式调度, 让该M上的其他G获得一定的运行时间, 以保证调度系统的公平性?(WorkingStealing)
---
Go的运行时调度器中有类似的抢占机制, 但并不能保证抢占能成功。因为Go运行时系统并没有内核调度器的中断能力, 它只能通过向运行时间过长的G中设置抢占flag的方法温柔的让运行的G自己主动让出M的执行权。Goroutine在运行过程中可以动态扩展自己线程栈的能力, 可以从初始的2KB大小扩展到最大1G, 在每次调用函数之前需要先计算该函数调用需要的栈空间大小, 然后按需扩展。Go抢占式调度的机制就是利用在判断要不要扩栈的时候顺便查看以下自己的抢占flag, 决定是否继续执行, 还是让出自己。
运行时系统的监控线程会计时并设置抢占flag到运行时间过长的G，然后G在有函数调用的时候会检查该抢占flag, 如果已设置就将自己放入全局队列, 这样该M上关联的其他G就有机会执行了。但如果正在执行的G是个很耗时的操作且没有任何函数调用(如只是for循环中的计算操作), 即使抢占flag已经被设置, 该G还是将一直霸占着当前M直到执行完自己的任务。

### 可视化 GMP 
---
方式1 :go tool trace
```go
package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {

	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//main
	fmt.Println("Hello World")
}
// go run trace.go 
// go tool trace trace.out 
```

方式2: debug trace
```shell
go build trace2.go
GODEBUG=schedtrace=1000 ./trace2 
```

## sysmon
sysmon 中会进行 netpool（获取 fd 事件）、retake（抢占）、forcegc（按时间强制执行 gc），scavenge heap（释放自由列表中多余的项减少内存占用）等处理。