package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Node - my tree node structure
type Node struct {
	nodes    []Node
	length   int
	metadata []int
}

func day8() {
	input := strings.Split(readFile("../input/day8.txt")[0], " ")
	data := make([]int, len(input))
	for i, datum := range input {
		val, _ := strconv.Atoi(datum)
		data[i] = val
	}

	// Assuming there is one tree root (I believe this problem makes that assumption)
	tree := getNodes(data, 1, 1)
	fmt.Println(fmt.Sprintf("Metadata sum = %d", getTreeMetadataSum(tree)))
	fmt.Println(fmt.Sprintf("Root value = %d", getTreeValue(tree[0])))
}

func getNodes(data []int, numNodes int, level int) []Node {
	var incrementor int
	nodes := make([]Node, numNodes)
	for i := 0; i < numNodes; i++ {
		numChildren := data[incrementor]
		numMetadata := data[incrementor+1]
		var childrenLength int

		if numChildren == 0 {
			nodes[i] = Node{make([]Node, 0), 2 + numMetadata, data[incrementor+2 : incrementor+2+numMetadata]}
		} else {
			childNodes := getNodes(data[incrementor+2:], numChildren, level+1)
			for _, child := range childNodes {
				childrenLength += child.length
			}
			nodes[i] = Node{childNodes, 2 + childrenLength + numMetadata, data[incrementor+childrenLength+2 : incrementor+childrenLength+2+numMetadata]}
		}
		incrementor += 2 + childrenLength + numMetadata
	}

	return nodes
}

func getTreeMetadataSum(tree []Node) int {
	if len(tree) == 0 {
		return 0
	}
	var sum int
	for _, node := range tree {
		sum += getTreeMetadataSum(node.nodes)
		for _, metadatum := range node.metadata {
			sum += metadatum
		}
	}
	return sum
}

func getTreeValue(root Node) int {
	var value int
	if len(root.nodes) == 0 {
		for _, metadatum := range root.metadata {
			value += metadatum
		}
	} else {
		for _, metadatum := range root.metadata {
			if metadatum > 0 && metadatum <= len(root.nodes) {
				value += getTreeValue(root.nodes[metadatum-1])
			}
		}
	}

	return value
}
