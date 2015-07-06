package merkle

import (
	"crypto/md5"
	"fmt"
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

// splits data in n pieces and builds tree from those pieces
func TreeFromData(data []byte, pieces int) *MerkleTree {
	size := len(data) / pieces
	dataPieces := [][]byte{}
	for i := 0; i < pieces; i++ {
		dataPieces = append(dataPieces, data[i*size:(i+1)*size])
	}
	return BuildTree(dataPieces)
}

func BuildTree(data [][]byte) *MerkleTree {
	var height uint

	height = 1
	nodes := GenerateLeaves(data)
	for len(nodes) > 1 {
		nodes = levelUp(nodes)
		height += 1
	}
	mt := &MerkleTree{
		Root:       nodes[0],
		TreeHeight: height,
	}
	return mt
}

func CompareTrees(mt1, mt2 *MerkleTree) bool {
	if mt1.TreeHeight != mt2.TreeHeight {
		return false
	}
	q1 := []*Node{mt1.Root}
	q2 := []*Node{mt2.Root}

	for len(q1) > 0 {
		if CmpByteArr(q1[0].DataHash, q2[0].DataHash) == false {
			return false
		}
		q1 = updateQueue(q1)
		q2 = updateQueue(q2)
	}

	return true
}

func FindDiff(mt1, mt2 *MerkleTree) [][]*Node {
	var nodes [][]*Node
	st1, st2 := TreeDiff(mt1, mt2) // return subtree where two trees differ
	if st1.Root.Right == nil && st1.Root.Left == nil {
		nodes = append(nodes, []*Node{st1.Root})
	} else {
		leftSt1 := &MerkleTree{Root: st1.Root.Left}
		leftSt2 := &MerkleTree{Root: st2.Root.Left}
		rightSt1 := &MerkleTree{Root: st1.Root.Left}
		rightSt2 := &MerkleTree{Root: st2.Root.Left}
		nodes = append(nodes, FindDiff(leftSt1, leftSt2)...)
		nodes = append(nodes, FindDiff(rightSt1, rightSt2)...)
	}
	return nodes
}

func TreeDiff(mt1, mt2 *MerkleTree) (st1, st2 *MerkleTree) {
	q1 := []*Node{mt1.Root}
	q2 := []*Node{mt2.Root}
	for len(q1) > 0 {
		if CmpByteArr(q1[0].DataHash, q2[0].DataHash) == false {
			return &MerkleTree{Root: q1[0]}, &MerkleTree{Root: q2[0]}
		}
		q1 = updateQueue(q1)
		q2 = updateQueue(q2)
	}
	return nil, nil
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

func getLeaves(t *MerkleTree) []*Node {
	var leaves []*Node
	q := []*Node{t.Root}

	for len(q) > 0 {
		if q[0].Left == nil && q[0].Right == nil {
			leaves = append(leaves, q[0])
		}
		if q[0].Left != nil {
			q = append(q, q[0].Left)
		}
		if q[0].Right != nil {
			q = append(q, q[0].Right)
		}
		if len(q) > 0 {
			q = q[1:]
		} else {
			q = []*Node{}
		}
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
		data := append(nodes[i*2].DataHash, nodes[i*2+1].DataHash...)
		hash := md5.Sum(data)
		node := &Node{
			DataHash: hash[:],
			Left:     nodes[i*2],
			Right:    nodes[i*2+1],
		}
		nextLevel = append(nextLevel, node)
	}
	if len(nodes)%2 == 1 {
		node := &Node{
			DataHash: nodes[len(nodes)-1].DataHash,
			Right:    nodes[len(nodes)-1],
		}
		nextLevel = append(nextLevel, node)
	}
	return nextLevel
}

func printNodeArr(nodes []*Node) {
	for _, node := range nodes {
		fmt.Println(node.DataHash)
	}
}

func CmpByteArr(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
