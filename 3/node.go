package node

import "math"

type Node struct {
	value     string `json:"value"`
	leftNode  *Node  `json:"leftNode"`
	rightNode *Node  `json:"rightNode"`
}

func New(value string, leftNode *Node, rightNode *Node) *Node {
	return &Node{
		value:     value,
		leftNode:  leftNode,
		rightNode: rightNode,
	}
}

func helperSerialize(node *Node) []string {
	if node == nil {
		return []string{}
	}

	var sequences []string
	sequences = append(sequences, helperSerialize(node.leftNode)...)
	sequences = append(sequences, node.value)
	sequences = append(sequences, helperSerialize(node.rightNode)...)
	return sequences
}

func Serialize(node *Node) []string {
	return helperSerialize(node)
}

func helperDeserialize(strs []string) *Node {
	switch {
	case len(strs) == 0:
		return nil
	case len(strs) == 1:
		return &Node{value: strs[0]}
	}

	var n Node
	middleIndex := int(math.Floor(float64(len(strs)) / 2))
	n.value = strs[middleIndex]
	n.leftNode = helperDeserialize(strs[0:middleIndex])
	if middleIndex+1 < len(strs) {
		n.rightNode = helperDeserialize(strs[middleIndex+1:])
	}

	return &n
}

func Deserialize(strs []string) *Node {
	return helperDeserialize(strs)
}
