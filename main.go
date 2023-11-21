package main

import "fmt"

const size = 5

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int // Changed the type to int
}

type Hash map[string]*Node

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head
	return Queue{Head: head, Tail: tail, Length: 0}
}

func (c *Cache) Check(str string) {
	node := &Node{}
	val, ok := c.Hash[str]
	if ok {
		node = c.Remove(val)
	} else {
		node = &Node{Value: str}
	}
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("Remove %s\n", n.Value)
	left := n.Left
	right := n.Right
	left.Right = right
	right.Left = left
	delete(c.Hash, n.Value)
	c.Queue.Length--
	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("add:%s\n", n.Value)
	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n
	c.Queue.Length++
	if c.Queue.Length > size {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d-[", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Value)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("start cache")
	cache := NewCache()
	for _, value := range []string{"cow", "dog", "cat", "rat", "dog"} {
		cache.Check(value)
		cache.Display()
	}
}
