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


*/

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Log(twoSum([]int{1, 2}, 3))
}
