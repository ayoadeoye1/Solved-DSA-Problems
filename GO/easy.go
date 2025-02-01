package golang

import (
	"fmt"
	"math"
)

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

/**
 *
 * @name Invert-Tree
 *
 * @description Invert-Binary-Tree
 *
 *
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

// BFS Graph Implementation
func bfs(start string, graph map[string][]string) {
	queue := []string{start}
	visted := make(map[string]bool)
	for len(queue) > 0 {
		// fmt.Println("Q deQ", queue)
		current := queue[0] //get index 1 of the queue
		queue = queue[1:]   //remove first index from queue
		if visted[current] {
			continue
		}
		visted[current] = true
		fmt.Println("Current City: ", current)
		for _, neigh := range graph[current] {
			if !visted[neigh] {
				queue = append(queue, neigh)
			}
		}
	}
}
func bfsMain() {
	graph := map[string][]string{
		"Abuja":      {"Minna", "Lokoja", "Jos", "Kaduna"},
		"Minna":      {"Abuja", "Kaduna", "Ilorin"},
		"Lokoja":     {"Abuja", "Benin City", "Makurdi"},
		"Jos":        {"Abuja", "Bauchi", "Kaduna", "Gombe"},
		"Kaduna":     {"Abuja", "Kano", "Zaria", "Jos"},
		"Ilorin":     {"Minna", "Ibadan", "Oshogbo"},
		"Benin City": {"Lokoja", "Asaba", "Warri"},
		"Makurdi":    {"Lokoja", "Enugu", "Jalingo"},
		"Bauchi":     {"Jos", "Gombe", "Maiduguri"},
		"Gombe":      {"Jos", "Bauchi", "Yola"},
		"Kano":       {"Kaduna", "Katsina", "Dutse"},
		"Zaria":      {"Kaduna", "Kano"},
		"Asaba":      {"Benin City", "Onitsha", "Awka"},
		"Onitsha":    {"Asaba", "Awka", "Enugu"},
		"Enugu":      {"Onitsha", "Awka", "Abakaliki"},
		"Ibadan":     {"Ilorin", "Oshogbo", "Abeokuta", "Lagos"},
		"Oshogbo":    {"Ibadan", "Ado-Ekiti", "Akure"},
		"Ado-Ekiti":  {"Oshogbo", "Akure"},
		"Akure":      {"Ado-Ekiti", "Benin City"},
		"Abeokuta":   {"Ibadan", "Lagos"},
		"Lagos":      {"Abeokuta", "Ibadan"},
	}
	fmt.Println("Traversal starts, Cities in Nigeria by distance from Abuja FCT") //Change the city name to find nearest to city you input
	bfs("Abuja", graph)
}

// Edge represents a weighted edge in the graph
type Edge struct {
	node   string
	weight int
}

// Dijkstra's algorithm implementation
func dijkstra(start string, graph map[string][]Edge) map[string]int {
	distances := make(map[string]int)
	for node := range graph {
		distances[node] = math.MaxInt32 // Initialize distances to infinity
	}
	distances[start] = 0 // Distance to the start node is 0

	// Simulate a priority queue using a slice
	queue := []string{start}

	return distances
}
