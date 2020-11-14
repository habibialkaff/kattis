package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Item struct {
	value int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value > pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

type Queue []int

func (s *Queue) Enqueue(val int) {
	*s = append(*s, val)
}

func (s *Queue) Dequeue() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	} else {
		element := (*s)[0]
		*s = (*s)[1:]
		return element, true
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var a int
		fmt.Sscanf(scanner.Text(), "%d\n", &a)

		var stack Stack
		var queue Queue
		pqueue := make(PriorityQueue, 0)

		heap.Init(&pqueue)

		isStack, isQueue, isPqueue := true, true, true

		for i := 0; i < a; i++ {
			scanner.Scan()
			var cmd, val int
			fmt.Sscanf(scanner.Text(), "%d%d\n", &cmd, &val)

			if cmd == 1 {
				if isStack {
					stack.Push(val)
				}

				if isQueue {
					queue.Enqueue(val)
				}

				if isPqueue {
					item := &Item{
						value:  val,
					}

					heap.Push(&pqueue, item)
				}
			} else {
				if isStack {
					x, y := stack.Pop()
					if !y || x != val {
						isStack = false
					}
				}

				if isQueue {
					x, y := queue.Dequeue()
					if !y || x != val {
						isQueue = false
					}
				}

				if isPqueue {
					if pqueue.Len() == 0 {
						isPqueue = false
					} else {
						item := heap.Pop(&pqueue).(*Item)
						if item.value != val {
							isPqueue = false
						}
					}
				}
			}
		}

		result := "impossible"

		if isStack {
			result = "stack"

			if isQueue || isPqueue {
				result = "not sure"
			}
		} else if isQueue {
			result = "queue"
			if isPqueue {
				result = "not sure"
			}
		} else if isPqueue {
			result = "priority queue"
		}

		fmt.Println(result)
	}
}
