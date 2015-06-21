package merkle

import (
	"crypto/md5"
	"math/rand"
	"testing"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

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
			t.Errorf("tree wasn't built correctly")
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
