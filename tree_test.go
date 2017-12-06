package tree_test

import (
	"fmt"
	"github.com/audathuynh/tree"
	"reflect"
	"testing"
)

func TestTreeTraversal(t *testing.T) {
	aTree := tree.NewTree()
	aTree.CreateTestingTree()

	// LNR traversal
	fmt.Println("LNR traversal:")
	visit := func(p *tree.TreeNode) {
		if p != nil {
			fmt.Println(p.Key)
		}
	}
	aTree.TraverseLNR(visit)
}

func TestVerticalSum(t *testing.T) {
	aTree := tree.NewTree()
	aTree.CreateTestingTree()

	fmt.Println("Testing vertical sum algorithm")
	// Get vertical sum of the tree
	result := make(map[int]int)
	aTree.VerticalSum(&result)
	expected := make(map[int]int)
	expected[-1] = 2
	expected[-2] = 1
	expected[0] = 8
	expected[1] = 6
	eq := reflect.DeepEqual(result, expected)
	if !eq {
		fmt.Println("Vertical sum:", result)
		t.Errorf("Vertical sum is wrong")
	}
}
