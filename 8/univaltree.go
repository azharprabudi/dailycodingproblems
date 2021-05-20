package univaltree

type Node struct {
	value int
	left  *Node
	right *Node
}

func helper(rootVal int, node *Node) int {
	if node == nil {
		return 0
	}

	count := 0
	count = count + helper(node.value, node.left)
	count = count + helper(node.value, node.right)

	if node.value == rootVal {
		count = count + 1
	}

	return count
}

func CountNumberOfUnivalSubtrees(node *Node) int {
	count := 0
	if node == nil {
		return count
	}

	if node.left != nil {
		count = count + helper(node.value, node.left)
	}

	if node.right != nil {
		count = count + helper(node.value, node.right)
	}

	return count + 1
}

func New(value int, left *Node, right *Node) *Node {
	return &Node{value: value, left: left, right: right}
}
