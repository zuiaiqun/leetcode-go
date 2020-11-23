package main

/*

Question:

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.


Answer:
寻找a+b=sum，等价于对给定的a，判断值=sum-b的元素是否存在，可用hashtable处理
时间复杂度：建立哈希表O(n) + 遍历数组O(n) =O(n)
空间复杂度：O(n)

统计结果:
运行时间基本在0ms，比所有ac的都低
空间高于66%的提交

cases:
1. 由于每个元素只能使用1次，查找leftValue时需要判断对应的下标是否满足条件：
   例如[3, 4], 6：遍历查找找leftValue=6-3=3时，找到的下标应大于3（顺序遍历查找，如果存在下标<3的一组满足，那应该这组就return了）




*/

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Log(twoSum([]int{1, 2}, 3))
}
