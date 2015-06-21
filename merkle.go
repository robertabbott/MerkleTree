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
	if mt1.TreeHeight != mt2.TreeHeight {
		return false
	}
	q1 := []*Node{mt1.Root}
	q2 := []*Node{mt2.Root}

	for len(q1) > 0 {
		if q1[0] != q2[0] {
			return false
		}
		q1 = updateQueue(q1)
		q2 = updateQueue(q2)
	}

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

func updateQueue(q []*Node) []*Node {
	if q[0].Left != nil {
		q = append(q, q[0].Left)
	}
	if q[0].Right != nil {
		q = append(q, q[0].Right)
	}
	if len(q) > 1 {
		return q[1:]
	} else {
		return []*Node{}
	}
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
