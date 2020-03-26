```ruby
# 快速排序
# 通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列

array = (1..100).to_a.sample(10)

def quick_sort(array)
  return [] if array.empty?
  x, *m = array
  left, right = m.partition{|a| a < x}
  quick_sort(left) + [x] + quick_sort(right)
end


# 插入排序
# 将n按照大小插入到0~n-1中, 当左边都排序好了, 整个数组也就排序好了
array = [10,9,8,7,6,5,4,3,2,1]
def insert_sort(array)
  size = array.size
  (0...size).to_a.each do |i|
    key = array[i]
    j = i - 1
    while j >= 0 and array[j] > key
      array[j + 1] = array[j]
      j = j - 1
    end
    array[j + 1] = key
  end
  array
end


# 冒泡排序
# 依次比较两个相邻的元素,按大小交换位置
def bubble_sort(array)
  length = array.length
  (1...length).to_a.each do |i|
    (0...(length - i)).to_a.each do |j|
      array[j +1], array[j] = array[j], array[j + 1] if array[j + 1] < array[j]
    end
  end
  array
end


# 鸡尾酒排序
# 先找到最小的数字，把他放到第一位，然后找到最大的数字放到最后一位。然后再找到第二小的数字放到第二位，再找到第二大的数字放到倒数第二位
def cocktail_sort(array)
  size = array.length
  (1..size/2).to_a.each do
    (0...(size-1)).to_a.each do |i|
      array[i + 1], array[i] = array[i], array[i + 1] if array[i] > array[i + 1]
    end

    (0...(size-1)).to_a.each do |i|
      array[i], array[i + 1] = array[i + 1], array[i] if array[i] > array[i + 1]
    end
  end
  array
end


# 合并排序
# 即把待排序序列分为若干个子序列，每个子序列是有序的。然后再把有序子序列合并为整体有序序列。
def merge_sort(array)
  size = array.size
  return array if size <= 1
  
  left = array[0, size/2]
  right = array[size/2, size - size/2]
  merge(merge_sort(left), merge_sort(right))
end

def merge(left, right)
  sorted = []
  until left.empty? || right.empty?
    sorted << (left.first <= right.first ? left.shift : right.shift)
  end
  sorted + left + right
end
```