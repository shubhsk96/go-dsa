package main

import "fmt"

type Stack struct {
	top  int
	list []*Node
}

type BinarySearchTree struct {
	root *Node
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	tree := &BinarySearchTree{root: nil}

	tree.createNode(6)
	tree.createNode(1)
	tree.createNode(3)
	tree.createNode(9)
	tree.createNode(7)
	tree.createNode(8)
	tree.createNode(5)

	fmt.Printf("BST Root: %v\n", tree.root.data)
	fmt.Printf("Inorder: %v", tree.printInorderTraversal(tree.root))
}

func (s *Stack) pop() *Node {
	if len(s.list) == 0 {
		return &Node{data: -1, left: nil, right: nil}
	}

	var item = s.list[s.top]
	s.list = s.list[:s.top]
	s.top = s.top - 1
	return item
}

func (s *Stack) push(node *Node) []*Node {
	s.top = s.top + 1
	s.list = append(s.list, node)
	return s.list
}

func (t *BinarySearchTree) printInorderTraversal(root *Node) []int {
	var result []int
	stack := &Stack{list: []*Node{}, top: -1}

	if root == nil {
		return []int{}
	}

	current := root
	for len(stack.list) != 0 || current != nil {
		if current != nil {
			stack.push(current)
			current = current.left
		} else {
			item := stack.list[stack.top].data
			poppedNode := stack.pop()
			result = append(result, item)
			current = poppedNode.right
		}
	}
	return result
}

func (t *BinarySearchTree) createNode(data int) {
	t.root = t.insertNode(t.root, data)
}

func (t *BinarySearchTree) insertNode(root *Node, data int) *Node {
	if root == nil {
		root = &Node{data: data, left: nil, right: nil}
		return root
	}

	if root.data >= data {
		root.left = t.insertNode(root.left, data)
	} else {
		root.right = t.insertNode(root.right, data)
	}
	return root
}
