package tree

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	Key   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func NewTree() *Tree {
	return &Tree{root: nil}
}

func (self *Tree) Input() {
	fmt.Print("Please input the root")
	self.input(&self.root)
}

func (self *Tree) TraverseNLR(visit func(*TreeNode)) {
	self.traverseNLR(self.root, visit)
}

func (self *Tree) TraverseLNR(visit func(*TreeNode)) {
	self.traverseLNR(self.root, visit)
}

func (self *Tree) TraverseLRN(visit func(*TreeNode)) {
	self.traverseLRN(self.root, visit)
}

// Given a binary tree, find the vertical sum of the binary tree
func (self *Tree) VerticalSum(m *map[int]int) {
	self.verticalSum(self.root, 0, m)
}

func (self *Tree) input(p **TreeNode) {
	fmt.Println(" (key 0 is nil):")
	var k int
	fmt.Scanf("%d", &k)
	if k == 0 {
		*p = nil
	} else {
		*p = new(TreeNode)
		(*p).Key = k
		fmt.Print("Please input left child of ", k)
		self.input(&((*p).Left))
		fmt.Print("Please input right child of ", k)
		self.input(&((*p).Right))
	}
}

// Create a BST tree manually for testing
// Hint: Use three variables to keep tree nodes: left, right, and parent.
// Then build the tree gradually from the highest level to the root.
func (self *Tree) CreateTestingTree() {
	l := new(TreeNode)
	l.Key = 1
	l.Left = nil
	l.Right = nil
	r := new(TreeNode)
	r.Key = 3
	r.Left = nil
	r.Right = nil

	p := new(TreeNode)
	p.Key = 2
	p.Left = l
	p.Right = r

	l = p
	r = new(TreeNode)
	r.Key = 6
	r.Left = nil
	r.Right = nil

	p = new(TreeNode)
	p.Key = 5
	p.Left = l
	p.Right = r

	self.root = p
}

func (self *Tree) traverseNLR(p *TreeNode, visit func(*TreeNode)) {
	if p != nil {
		visit(p)
		self.traverseNLR(p.Left, visit)
		self.traverseNLR(p.Right, visit)
	}
}

func (self *Tree) traverseLNR(p *TreeNode, visit func(*TreeNode)) {
	if p != nil {
		self.traverseLNR(p.Left, visit)
		visit(p)
		self.traverseLNR(p.Right, visit)
	}
}

func (self *Tree) traverseLRN(p *TreeNode, visit func(*TreeNode)) {
	if p != nil {
		self.traverseLNR(p.Left, visit)
		self.traverseLNR(p.Right, visit)
		visit(p)
	}
}

func (self *Tree) verticalSum(p *TreeNode, index int, m *map[int]int) {
	if p != nil {
		self.processNode(p, index, m)
		self.verticalSum(p.Left, index-1, m)
		self.verticalSum(p.Right, index+1, m)
	}
}

func (self *Tree) processNode(p *TreeNode, index int, m *map[int]int) {
	if p != nil {
		value, ok := (*m)[index]
		if ok {
			value += p.Key
		} else {
			value = p.Key
		}
		(*m)[index] = value
	}
}

type QueueNode struct {
	Data *TreeNode
	Next *QueueNode
}

func NewQueueNode() *QueueNode {
	return &QueueNode{Data: nil, Next: nil}
}

type Queue struct {
	Front           *QueueNode
	Rear            *QueueNode
	Count           int
	IsPriorityQueue bool
}

func NewQueue(isPriorityQueue bool) *Queue {
	return &Queue{Front: nil, Rear: nil, Count: 0, IsPriorityQueue: isPriorityQueue}
}

func (queue *Queue) Enqueue(data *TreeNode) {
	var newPtr *QueueNode = NewQueueNode()
	newPtr.Data = data
	newPtr.Next = nil
	if queue.Count == 0 {
		queue.Front = newPtr
		queue.Rear = newPtr
	} else {
		if !queue.IsPriorityQueue {
			queue.Rear.Next = newPtr
			queue.Rear = newPtr
		} else {
			if data == nil { // do not accept nil data in priority queue
				return
			}
			ptr := queue.Front
			// search to find a correct place to put newPtr
			var prePtr *QueueNode
			prePtr = nil
			for ptr != nil && data.Key > ptr.Data.Key {
				prePtr = ptr
				ptr = ptr.Next
			}
			if ptr == nil {
				queue.Rear = newPtr
			}
			if prePtr == nil {
				queue.Front = newPtr
			} else {
				prePtr.Next = newPtr
			}
			newPtr.Next = ptr
		}
	}
	queue.Count++
}

func (queue *Queue) Dequeue(dataOut **TreeNode) error {
	if queue.Count == 0 {
		return errors.New("Queue is empty")
	}
	*dataOut = queue.Front.Data
	if queue.Count == 1 {
		queue.Rear = nil
	}
	queue.Front = queue.Front.Next
	queue.Count--
	return nil
}

func (queue *Queue) IsEmpty() bool {
	return queue.Count == 0
}

func (queue *Queue) GetFront() *TreeNode {
	if queue.Count == 0 {
		return nil
	}
	return queue.Front.Data
}

func (queue *Queue) GetRear() *TreeNode {
	if queue.Count == 0 {
		return nil
	}
	return queue.Rear.Data
}
