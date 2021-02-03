```ruby
# 版本1 数组
class TurboHash
  attr_reader :table

  def initialize
    @table = []
  end
  
  def [](key)
    find(key).last
  end

  def find(key)
    # 遍历数组查找对应的条目
    @table.find do |entry|
      key == entry.first
    end
  end

  def []=(key, value)
    if entry = find(key)
      entry[1] = value
    else
      @table << [key, value]
    end
  end
end

require "benchmark"

legacy = Hash.new
turbo  = TurboHash.new

n = 10_000

def set_and_find(target)
  key = rand
  target[key] = rand
  target[key]
end

Benchmark.bm do |x|
  x.report("Hash: ") { n.times { set_and_find(legacy) } }
  x.report("TurboHash: ") { n.times { set_and_find(turbo) } }
end


# 版本2
# 哈希值链表法
class Node
  attr_reader :object, :next

  def initialize(o, n)
    @object = o
    @next = n
  end
end

class TurboHash
  NUM_BINS = 11

  attr_reader :table

  def initialize
    @table = Array.new(NUM_BINS)
  end

  def [](key)
    if node = node_for(key)
      begin
        return node.object[1] if node.object[0] == key
      end while node = node.next
    end
  end

  def node_for(key)
    @table[index_of(key)]
  end

  def index_of(key)
    key.hash % NUM_BINS
  end

  def []=(key, value)
    @table[index_of(key)] = Node.new([key, value], node_for(key))
  end
end

# 版本3 哈希值链表扩容
class Node
  attr_reader :object, :next

  def initialize(o, n)
    @object = o
    @next = n
  end
end

class TurboHash
  STARTING_BINS = 16

  attr_reader :table

  def initialize
    @max_density = 5
    @entry_count = 0
    @bin_count   = STARTING_BINS
    @table = Array.new(@bin_count)
  end

  def grow
    @bin_count = @bin_count << 1
    new_table = Array.new(@bin_count)
    @table.each do |node|
      if node
        begin
          new_index = index_of(node.object[0])
          new_table[new_index] = Node.new(node.object, new_table[new_index])
        end while node = node.next
      end
    end
    @table = new_table
  end

  def full?
    @entry_count > @max_density * @bin_count
  end

  def [](key)
    if node = node_for(key)
      begin
        return node.object[1] if node.object[0] == key
      end while node = node.next
    end
  end

  def node_for(key)
    @table[index_of(key)]
  end

  def index_of(key)
    key.hash % @bin_count
  end

  def []=(key, value)
    grow if full?
    @table[index_of(key)] = Node.new([key, value], node_for(key))
    @entry_count += 1
  end
end
```

开放地址法:
不再额外维护一个链表结构，每次都是按 hash 索引在数组插入，如果要插入的位置已经有元素存在，那么就找到一个空的位置插入。
【线性探测 随机探测】



