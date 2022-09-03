### Collection
---
> List

1. ArrayList: Object[] 数组
2. Vector: Object[] 数组
3. LinkedList: 双向链表(JDK1.6 之前为循环链表，JDK1.7 取消了循环)

> Set

1. HashSet: 无序, 唯一, 基于 HashMap 实现的, 底层采用 HashMap 来保存元素
2. LinkedHashSet: LinkedHashSet 是 HashSet 的子类, 并且其内部是通过 LinkedHashMap 来实现的.
3. TreeSet: 有序, 唯一, 红黑树(自平衡的排序二叉树)

> Queue

1. PriorityQueue: Object[] 数组来实现二叉堆
2. ArrayQueue: Object[] 数组 + 双指针

> Map

1. HashMap: JDK1.8 之前 HashMap 由数组+链表组成的, 数组是 HashMap 的主体, 链表则是主要为了解决哈希冲突而存在的. JDK1.8 以后在解决哈希冲突时有了较大的变化, 当链表长度大于阈值(默认为 8)(将链表转换成红黑树前会判断, 如果当前数组的长度小于 64, 那么会选择先进行数组扩容, 而不是转换为红黑树)时, 将链表转化为红黑树, 以减少搜索时间
2. LinkedHashMap: LinkedHashMap 继承自 HashMap, 所以它的底层仍然是基于散列结构即由数组和链表或红黑树组成. 另外, LinkedHashMap 在上面结构的基础上, 增加了一条双向链表, 使得上面的结构可以保持键值对的插入顺序. 同时通过对链表进行相应的操作, 实现了访问顺序相关逻辑.
3. Hashtable: 数组+链表组成的, 数组是 Hashtable 的主体, 链表则是主要为了解决哈希冲突而存在的
4. TreeMap: 红黑树(自平衡的排序二叉树)


### List Set Queue Map 四者的区别?
---
1. List: 存储的元素是有序, 可重复的
2. Set: 存储的元素是无序的, 不可重复的
3. Queue: 按特定的排队规则来确定先后顺序, 存储的元素是有序的, 可重复的
4. Map: 使用键值对(key-value)存储, key 是无序的, 不可重复的, value 是无序的, 可重复的, 每个键最多映射到一个值


### ArrayList 和 Vector 的区别?
---
1. ArrayList 是 List 的主要实现类, 底层使用 Object[ ]存储, 适用于频繁的查找工作, 线程不安全.
2. Vector 是 List 的古老实现类, 底层使用Object[ ] 存储, 线程安全的


### ArrayList 与 LinkedList 区别?
---
1. 是否保证线程安全: ArrayList 和 LinkedList 都是不同步的, 也就是不保证线程安全
2. 底层数据结构: ArrayList 底层使用的是 Object 数组; LinkedList 底层使用的是 双向链表 数据结构

> 插入和删除是否受元素位置的影响:

1. ArrayList 采用数组存储, 所以插入和删除元素的时间复杂度受元素位置的影响.
2. LinkedList 采用链表存储, 所以, 如果是在头尾插入或者删除元素不受元素位置的影响.
3. 是否支持快速随机访问: LinkedList 不支持高效的随机元素访问, 而 ArrayList 支持. 
4. 内存空间占用: ArrayList 的空 间浪费主要体现在在 list 列表的结尾会预留一定的容量空间, 而 LinkedList 的空间花费则体现在它的每一个元素都需要消耗比 ArrayList 更多的空间(因为要存放直接后继和直接前驱以及数据)
5. 我们在项目中一般是不会使用到 LinkedList 的, 需要用到 LinkedList 的场景几乎都可以使用 ArrayList 来代替, 并且, 性能通常会更好! 就连 LinkedList 的作者约书亚 · 布洛克（Josh Bloch）自己都说从来不会使用 LinkedList

### 比较 HashSet LinkedHashSet 和 TreeSet 三者的异同
---
1. HashSet LinkedHashSet 和 TreeSet 都是 Set 接口的实现类, 都能保证元素唯一, 并且都不是线程安全的
2. HashSet LinkedHashSet 和 TreeSet 的主要区别在于底层数据结构不同. HashSet 的底层数据结构是哈希表(基于 HashMap 实现) LinkedHashSet 的底层数据结构是链表和哈希表, 元素的插入和取出顺序满足 FIFO. TreeSet 底层数据结构是红黑树, 元素是有序的，排序的方式有自然排序和定制排序
3. 底层数据结构不同又导致这三者的应用场景不同. HashSet 用于不需要保证元素插入和取出顺序的场景, LinkedHashSet 用于保证元素的插入和取出顺序满足 FIFO 的场景, TreeSet 用于支持对元素自定义排序规则的场景


### HashMap 和 Hashtable 的区别
---
1. 线程是否安全: HashMap 是非线程安全的, Hashtable 是线程安全的, 因为 Hashtable 内部的方法基本都经过synchronized 修饰. (如果你要保证线程安全的话就使用 ConcurrentHashMap)
2. 效率: 因为线程安全的问题, HashMap 要比 Hashtable 效率高一点. 另外，Hashtable 基本被淘汰, 不要在代码中使用它
3. 对 Null key 和 Null value 的支持: HashMap 可以存储 null 的 key 和 value, 但 null 作为键只能有一个, null 作为值可以有多个. Hashtable 不允许有 null 键和 null 值, 否则会抛出 NullPointerException。
4. 初始容量大小和每次扩充容量大小的不同: ① 创建时如果不指定容量初始值, Hashtable 默认的初始大小为 11, 之后每次扩充, 容量变为原来的 2n+1. HashMap 默认的初始化大小为 16. 之后每次扩充, 容量变为原来的 2 倍. ② 创建时如果给定了容量初始值, 那么 Hashtable 会直接使用你给定的大小, 而 HashMap 会将其扩充为 2 的幂次方大小 (HashMap 中的tableSizeFor()方法保证).
5. 底层数据结构: HashMap当链表长度大于阈值时, 将链表转化为红黑树以减少搜索时间. Hashtable 没有这样的机制