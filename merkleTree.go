package merkleTree

import (
	"crypto/md5"
)

type MerkleTree struct {
	TreeHeight uint
	Root       *Node
}

type Node struct {
	DataHash []byte
	Left     *Node
	Right    *Node
}

func BuildTree(mt *MerkleTree, data [][]byte) int {
	var height uint

	height = 1
	nodes := GenerateLeaves(data)
	for len(nodes) > 1 {
		nodes = levelUp(nodes)
		height += 1
	}
	mt.Root = nodes[0]
	mt.TreeHeight = height

	return 0
}

func CompareTrees(mt1, mt2 *MerkleTree) bool {
	n1 := mt1.Root
	n2 := mt2.Root

	// perform bfs, if nodes != return false
	return true
}

func FindDiff(mt1, mt2 *MerkleTree) *MerkleTree {
	// return subtree where two trees differ
}

func GenerateLeaves(data [][]byte) []*Node {
	var leaves []*Node
	for _, d := range data {
		hash := md5.Sum(d)
		node := &Node{
			DataHash: hash[:],
		}
		leaves = append(leaves, node)
	}
	return leaves
}

func levelUp(nodes []*Node) []*Node {
	var nextLevel []*Node
	for i := 0; i < len(nodes)/2; i++ {
		data := append(nodes[i].DataHash, nodes[i+1].DataHash...)
		hash := md5.Sum(data)
		node := &Node{
			DataHash: hash[:],
			Left:     nodes[i],
			Right:    nodes[i+1],
		}
		nextLevel = append(nextLevel, node)
	}
	if len(nodes)%2 == 1 {
		nextLevel = append(nextLevel, nodes[len(nodes)-1])
	}
	return nextLevel
}
