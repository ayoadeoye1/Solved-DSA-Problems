package golang

/**
 *
 * @name
 *
 * @description
 *
 *
 */
func canAliceWin(n int) bool {
	arem := 10
	brem := 9
	remain := n
	isAliceTurn := true

	var final bool

	for remain > 0 {
		if isAliceTurn {
			remain -= arem
			brem = arem - 1
			isAliceTurn = false
			if remain < 0 {
				final = false
				break
			}
			if remain == 0 {
				final = true
				break
			}
		} else {
			remain -= brem
			arem = brem - 1
			isAliceTurn = true
			if remain < 0 {
				final = true
				break
			}
			if remain == 0 {
				final = false
				break
			}
		}
	}
	return final
}

/**
 *
 * @name Has-Cycle
 *
 * @description Linked List Cycle
 *
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

/**
 *
 * @name Majority-Element
 *
 * @description Elements That Occured Most In An Array
 *
 * Optimal Solution -- Boyer Moore Voting Algorithm
 */
func majorityElement(nums []int) int {
	count := 0
	candidate := 0
	for i := 0; i <= len(nums)-1; i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}
