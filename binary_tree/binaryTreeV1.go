package main

import (
	"fmt"
)

func newNode() *Node {
	n := &Node{}
	return n
}

func (n *Node) Insert(newValue interface{}) error {
	switch {
	case n.Data == nil:
		{
			n.Data = newValue
		}

	case newValue.(int) < (n.Data).(int):
		{
			if n.Left == nil {
				n.Left = newNode()
				n.Left.Data = newValue
			} else {
				return n.Left.Insert(newValue)
			}
		}

	case newValue.(int) > (n.Data).(int):
		{
			if n.Right == nil {
				n.Right = newNode()
				n.Right.Data = newValue
			} else {
				return n.Right.Insert(newValue)
			}

		}
	}

	return nil
}

func (n *Node) Find(value interface{}) *Node {

	if n.Data == nil {
		return nil
	}

	switch {
	case n.Data == value:
		{
			return n
		}

	case value.(int) < n.Data.(int):
		{
			if n.Left == nil {
				break
			}
			return n.Left.Find(value)
		}

	case value.(int) > n.Data.(int):
		{
			if n.Right == nil {
				break
			}
			return n.Right.Find(value)

		}
	}
	return nil
}

func (n *Node) Print() {
	if n == nil {
		return
	}

	fmt.Printf("%v ", n.Data)

	(n.Left).Print()
	(n.Right).Print()
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}

func min(num1 int, num2 int) int {
	if num1 > num2 {
		return num2
	}

	return num1
}

func (n *Node) Height() int {
	var leftHeight, rightHeight int

	if n.Left != nil {
		leftHeight = n.Left.Height()
	}

	if n.Right != nil {
		rightHeight = n.Right.Height()
	}

	return 1 + max(leftHeight, rightHeight)
}

type Node struct {
	Left  *Node
	Right *Node
	Data  interface{}
}

func (n *Node) FindMax() int {
	if n.Right == nil {
		return n.Data.(int)
	}

	return max(n.Data.(int), n.Right.FindMax())
}

func (n *Node) HasTwoChildren() bool {
	if n.Left != nil {
		if n.Right != nil {
			return true
		}
	}

	return false
}

func (n *Node) HasNoChildren() bool {
	if n.Left == nil {
		if n.Right == nil {
			return true
		}
	}

	return false
}

func (n *Node) HasLeftChild() bool {
	if n.Left != nil {
		return true
	}

	return false
}

func (n *Node) HasRightChild() bool {
	if n.Right != nil {
		return true
	}

	return false
}

func (n *Node) FindMin() int {
	if n.Left == nil {
		return n.Data.(int)
	}

	return min(n.Data.(int), n.Left.FindMin())
}

func (n *Node) IsEmpty() bool {
	if n.Data == nil {
		return true
	}

	return false
}

func (n *Node) Delete(value interface{}) *Node {

	if n.Data == nil {
		return nil
	}

	switch {
	case value.(int) < n.Data.(int):
		{
			if n.Left == nil {
				return nil
			}

			if n.Left.Data.(int) == value.(int) {
				n.Left = n.Left.Delete(value)
			} else {
				n.Left.Delete(value)
			}
		}

	case value.(int) > n.Data.(int):
		{
			if n.Right == nil {
				return nil
			}

			if n.Right.Data.(int) == value.(int) {
				n.Right = n.Right.Delete(value)
			} else {
				n.Right.Delete(value)
			}
		}

	case value.(int) == n.Data:
		{
			switch {

			case n.HasNoChildren():
				{
					n.Data = nil
					return nil
				}

			case n.HasLeftChild() || n.HasTwoChildren():
				{
					n.Data = n.Left.FindMax()
					n.Left = n.Left.Delete(n.Data)
				}

			case n.HasRightChild():
				{
					n.Data = n.Right.FindMin()
					n.Right = n.Right.Delete(n.Data)
				}

			}

		}

	}

	return n
}

func main() {
	var bTree Node

	bTree.Insert(10)
	bTree.Insert(2)
	bTree.Insert(18)
	bTree.Insert(39)
	bTree.Insert(41)
	bTree.Insert(4)
	bTree.Insert(1)
	bTree.Insert(40)

	fmt.Printf("\nPrinted Tree Values:\n")
	bTree.Print()


	fmt.Printf("\n\nTesting Find()")

	fmt.Printf("\nAttempting to Find 180: ")
	if bTree.Find(180) != nil {
		fmt.Printf("FOUND\n")
	} else {
		fmt.Printf("NOT FOUND\n")
	}

	fmt.Printf("Attempting to Find 4: ")
	if bTree.Find(4) != nil {
		fmt.Printf("FOUND\n")
	} else {
		fmt.Printf("NOT FOUND\n")
	}

	fmt.Printf("\nThe height of the current tree is %d.\n\n", bTree.Height())

	fmt.Printf("Testing Delete() on Node without children:\n")
	fmt.Printf("Removing 1...")
	bTree.Delete(1)
	fmt.Printf("\nPrinted Tree Values:\n")
	bTree.Print()

	fmt.Printf("\n\nTesting Delete() on Node with one right child:\n")
	fmt.Printf("Removing 4...")
	bTree.Delete(4)
	fmt.Printf("\nPrinted Tree Values:\n")
	bTree.Print()

	fmt.Printf("\n\nTesting Delete() on Node with one left child:\n")
	fmt.Printf("Removing 40...")
	bTree.Delete(40)
	fmt.Printf("\nPrinted Tree Values:\n")
	bTree.Print()

	fmt.Printf("\n\nTesting Delete() on Node with two children:\n")
	fmt.Printf("Removing 10...")
	bTree.Delete(10)
	fmt.Printf("\nPrinted Tree Values:\n")
	bTree.Print()

	fmt.Printf("\n\nTesting Delete() on a Node with a nonexisting value:\n")
	fmt.Printf("Removing 49...")
	bTree.Delete(49)
	fmt.Printf("\nPrinted Tree Values:\n")
	bTree.Print()

	fmt.Printf("\n\nRemoving 18...")
	bTree.Delete(18)
	fmt.Printf("\nPrinted Tree Values:\n")
	bTree.Print()

	fmt.Printf("\n\nRemoving 41...")
	bTree.Delete(41)
	fmt.Printf("\nPrinted Tree Values:\n")
	bTree.Print()

	fmt.Printf("\n\nRemoving 2...")
	bTree.Delete(2)
	fmt.Printf("\nPrinted Tree Values:\n")
	bTree.Print()

	if bTree.IsEmpty() {
		fmt.Printf("\n\nbTree is empty.")
	} else {
		fmt.Printf("\n\nbTree is not empty.")
	}

	fmt.Printf("\n\nTesting to Delete() on last node")
	fmt.Printf("\nRemoving 39...")
	bTree.Delete(39)
	fmt.Printf("\nPrinted Tree Values:\n")
	bTree.Print()

	if bTree.IsEmpty() {
		fmt.Printf("\n\nbTree is empty.")
	} else {
		fmt.Printf("\n\nbTree is not empty.")
	}

	fmt.Printf("\n\n")
}
