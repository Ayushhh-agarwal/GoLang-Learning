package exercises

import (
	"fmt"
)

// 	Question :
//	Create an expression tree using golang structs. For example string a + b - c on tree would look like the following,
//				+
//			  /   \
//		 	 a     -
//				 /   \
//				b     c
//
//	Building the tree can be done manually through code. Once the tree is built,
//	traverse the tree in preorder and postorder format and print the values.
//
//						Expectation: Use Recursion, Strings, Pointers, Struct

type Node struct {
	Value     string
	LeftNode  *Node
	RightNode *Node
}

func NewNode(value string) *Node {
	return &Node{
		Value:     value,
		LeftNode:  nil,
		RightNode: nil,
	}
}

func isOperator(op string) bool {
	return op == "+" || op == "-" || op == "*" || op == "/"
}

func buildExpressionTreeWithPostorder(expression string) *Node {
	var stack []*Node
	var temp, t1, t2 *Node

	for _, char := range expression {
		token := string(char)

		if !isOperator(token) {
			temp = NewNode(token)
			stack = append(stack, temp)
		} else {
			temp = NewNode(token)

			t1 = stack[len(stack)-1]
			t2 = stack[len(stack)-2]

			temp.LeftNode = t2
			temp.RightNode = t1

			stack = append(stack, temp)
		}
	}
	temp = stack[len(stack)-1]
	return temp
}

func (n *Node) PreorderTraversal() {
	if n != nil {
		fmt.Printf("%s ", n.Value)
		n.LeftNode.PreorderTraversal()
		n.RightNode.PreorderTraversal()
	}
}

func (n *Node) PostorderTraversal() {
	if n != nil {
		n.LeftNode.PostorderTraversal()
		n.RightNode.PostorderTraversal()
		fmt.Printf("%s ", n.Value)
	}
}

func (n *Node) InOrderTraversal() {
	if n != nil {
		n.LeftNode.InOrderTraversal()
		fmt.Printf("%s ", n.Value)
		n.RightNode.InOrderTraversal()
	}
}

func MakeExpressionTree() {
	rootNode := NewNode("+")
	rootNode.LeftNode = NewNode("1")
	rootNode.RightNode = NewNode("-")
	rootNode.RightNode.LeftNode = NewNode("3")
	rootNode.RightNode.RightNode = NewNode("2")

	fmt.Println("Preorder Traversal:")
	rootNode.PreorderTraversal()
	fmt.Println()

	fmt.Println("Postorder Traversal:")
	rootNode.PostorderTraversal()
	fmt.Println()

	fmt.Println("InOrder Traversal:")
	rootNode.InOrderTraversal()
	fmt.Println()

	expression := "132-+"
	rootNode1 := buildExpressionTreeWithPostorder(expression)

	if rootNode1 != nil {
		fmt.Println("InOrder Traversal:")
		rootNode.InOrderTraversal()
		fmt.Println()

	} else {
		fmt.Println("Invalid expression")
	}
}
