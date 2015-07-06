package merkle

import (
	"crypto/md5"
	"math/rand"
	"testing"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func TestFindDiff(t *testing.T) {

}

func TestTreeFromData(t *testing.T) {
	var data []byte
	var dataPieces [][]byte
	pieces := 4
	data = randByteArr(36)
	for i := 0; i < pieces; i++ {
		dataPieces = append(dataPieces, data[i*9:(1+i)*9])
	}
	t1 := TreeFromData(data, pieces)
	t2 := BuildTree(dataPieces)
	if CompareTrees(t1, t2) == false {
		t.Fatalf("TreeFromData did not build tree correctly")
	}
}

func TestCompareTrees(t *testing.T) {
	var data [][]byte
	var d2, d3 [][]byte
	for i := 0; i < 5; i++ {
		data = append(data, randByteArr(16))
		d2 = append(d2, randByteArr(16))
	}
	d3 = d2
	d3[0] = []byte("seamus")
	t1 := BuildTree(data)
	t2 := BuildTree(data)
	t3 := BuildTree(d2) // totally different from t1
	t4 := BuildTree(d3) // only slightly different from t1
	if CompareTrees(t1, t2) == false {
		t.Fatalf("Trees are identical but compare trees returned false")
	}
	if CompareTrees(t1, t3) == true {
		t.Fatalf("Trees are identical but compare trees returned false")
	}
	if CompareTrees(t1, t4) == true {
		t.Fatalf("Trees are identical but compare trees returned false")
	}
}

func TestBuildTreeEven(t *testing.T) {
	var data [][]byte
	for i := 0; i < 4; i++ {
		data = append(data, randByteArr(16))
	}
	t1 := BuildTree(data)
	nodes := getLeaves(t1)
	for index, val := range data {
		sum := md5.Sum(val)
		if CmpByteArr(sum[:], nodes[index].DataHash) == false {
			t.Fatalf("tree wasn't built correctly")
		}
	}
}

func TestBuildTreeOdd(t *testing.T) {
	var data [][]byte
	for i := 0; i < 5; i++ {
		data = append(data, randByteArr(16))
	}
	t1 := BuildTree(data)
	nodes := getLeaves(t1)
	for index, val := range data {
		sum := md5.Sum(val)
		if CmpByteArr(sum[:], nodes[index].DataHash) == false {
			t.Fatalf("tree wasn't built correctly")
		}
	}
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randByteArr(n int) []byte {
	return []byte(randSeq(n))
}
