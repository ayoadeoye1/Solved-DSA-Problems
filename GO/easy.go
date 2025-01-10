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
