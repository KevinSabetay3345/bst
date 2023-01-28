package bst

import (
	"testing"
)

func TestEmptyBST(t *testing.T) {
	var bst BST

	checkNoElems(bst, t)

	min, err := bst.Minimum()
	if min != 0 {
		t.Errorf("Expected %d, Got %d", 0, min)
	}
	if err == nil {
		t.Fatalf("Expected %s, Got %v", "there are no elements", err)
	}
	if err.Error() != "there are no elements" {
		t.Errorf("Expected %s, Got %s", "there are no elements", err.Error())
	}

	max, err := bst.Maximum()
	if max != 0 {
		t.Errorf("Expected %d, Got %d", 0, max)
	}
	if err == nil {
		t.Fatalf("Expected %s, Got %v", "there are no elements", err)
	}
	if err.Error() != "there are no elements" {
		t.Errorf("Expected %s, Got %s", "there are no elements", err.Error())
	}

	sorted := bst.Inorder()
	if len(sorted) > 0 {
		t.Errorf("Expected %d, Got %d", 0, len(sorted))
	}

	for i := 0; i < 10; i++ {
		if bst.Exists(i) {
			t.Errorf("Expected %t, Got %t", false, true)
		}

		err = bst.Delete(i)
		if err == nil {
			t.Errorf("Expected %s, Got %v", "value not found", err)
		}

		next, err := bst.Next(i)
		if next != 0 {
			t.Errorf("Expected %d, Got %d", 0, next)
		}
		if err == nil {
			t.Fatalf("Expected %s, Got %v", "element not found", err)
		}
		if err.Error() != "element not found" {
			t.Errorf("Expected %s, Got %s", "element not found", err.Error())
		}
	}
}

func TestOneElemBST(t *testing.T) {
	//insert
	var bst BST

	err := bst.Insert(1)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst.Insert(1)
	if err == nil {
		t.Fatalf("Expected %s, Got %v", "value already exists", err)
	}
	if err.Error() != "value already exists" {
		t.Errorf("Expected %s, Got %s", "value already exists", err.Error())
	}

	checkOneElem(bst, 1, t)

	//delete
	err = bst.Delete(0)
	if err == nil {
		t.Errorf("Expected %s, Got %v", "value not found", err)
	}
	err = bst.Delete(1)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkNoElems(bst, t)

	err = bst.Insert(3)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkOneElem(bst, 3, t)

	//get
	if bst.Exists(4) {
		t.Errorf("Expected %t, Got %t", false, true)
	}
	if !bst.Exists(3) {
		t.Errorf("Expected %t, Got %t", true, false)
	}

	min, err := bst.Minimum()
	if err != nil {
		t.Fatalf("Expected %v, Got %s", nil, err.Error())
	}
	if min != 3 {
		t.Errorf("Expected %d, Got %d", 3, min)
	}

	max, err := bst.Maximum()
	if err != nil {
		t.Fatalf("Expected %v, Got %s", nil, err.Error())
	}
	if max != 3 {
		t.Errorf("Expected %d, Got %d", 3, max)
	}

	next, err := bst.Next(0)
	if next != 0 {
		t.Errorf("Expected %d, Got %d", 0, next)
	}
	if err == nil {
		t.Fatalf("Expected %s, Got %v", "element not found", err)
	}
	if err.Error() != "element not found" {
		t.Errorf("Expected %s, Got %s", "element not found", err.Error())
	}

	next, err = bst.Next(3)
	if next != 0 {
		t.Errorf("Expected %d, Got %d", 0, next)
	}
	if err == nil {
		t.Fatalf("Expected %s, Got %v", "there is no next element", err)
	}
	if err.Error() != "there is no next element" {
		t.Errorf("Expected %s, Got %s", "there is no next element", err.Error())
	}

	sorted := bst.Inorder()
	if len(sorted) != 1 {
		t.Fatalf("Expected %d, Got %d", 1, len(sorted))
	}
	if sorted[0] != 3 {
		t.Fatalf("Expected %d, Got %d", 3, sorted[0])
	}
}

func TestTwoElemsBST(t *testing.T) {
	//insert
	var bstRight BST
	err := bstRight.Insert(1)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bstRight.Insert(5)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}

	checkTwoElemsRight(bstRight, 1, 5, t)

	var bstLeft BST
	err = bstLeft.Insert(5)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bstLeft.Insert(1)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}

	checkTwoElemsLeft(bstLeft, 5, 1, t)

	//delete
	err = bstRight.Delete(1)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkOneElem(bstRight, 5, t)

	err = bstRight.Insert(2)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkTwoElemsLeft(bstRight, 5, 2, t)

	err = bstLeft.Delete(5)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkOneElem(bstLeft, 1, t)

	err = bstLeft.Insert(8)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkTwoElemsRight(bstLeft, 1, 8, t)

	//get
	var bst BST
	bst.Insert(2)
	bst.Insert(8)

	if !bst.Exists(2) {
		t.Errorf("Expected %t, Got %t", true, false)
	}
	if !bst.Exists(8) {
		t.Errorf("Expected %t, Got %t", true, false)
	}
	if bst.Exists(4) {
		t.Errorf("Expected %t, Got %t", false, true)
	}

	min, err := bst.Minimum()
	if err != nil {
		t.Fatalf("Expected %v, Got %s", nil, err.Error())
	}
	if min != 2 {
		t.Errorf("Expected %d, Got %d", 2, min)
	}

	max, err := bst.Maximum()
	if err != nil {
		t.Fatalf("Expected %v, Got %s", nil, err.Error())
	}
	if max != 8 {
		t.Errorf("Expected %d, Got %d", 8, max)
	}

	next, err := bst.Next(8)
	if next != 0 {
		t.Errorf("Expected %d, Got %d", 0, next)
	}
	if err == nil {
		t.Fatalf("Expected %s, Got %v", "there is no next element", nil)
	}
	if err.Error() != "there is no next element" {
		t.Errorf("Expected %s, Got %s", "there is no next element", err.Error())
	}

	next, err = bst.Next(2)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if next != 8 {
		t.Errorf("Expected %d, Got %d", 8, next)
	}

	sorted := bst.Inorder()
	if len(sorted) != 2 {
		t.Fatalf("Expected %d, Got %d", 2, len(sorted))
	}
	if sorted[0] != 2 {
		t.Fatalf("Expected %d, Got %d", 2, sorted[0])
	}
	if sorted[1] != 8 {
		t.Fatalf("Expected %d, Got %d", 8, sorted[1])
	}
}

func TestThreeElemsBST(t *testing.T) {
	//insert
	var bst BST
	err := bst.Insert(10)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst.Insert(80)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst.Insert(5)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkThreeElemsBalanced(bst, 10, 5, 80, t)

	var bst2 BST
	err = bst2.Insert(10)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst2.Insert(30)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst2.Insert(50)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkThreeElemsRigth(bst2, 10, 30, 50, t)

	var bst3 BST
	err = bst3.Insert(10)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst3.Insert(5)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst3.Insert(3)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkThreeElemsLeft(bst3, 10, 5, 3, t)

	var bst4 BST
	err = bst4.Insert(10)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst4.Insert(5)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst4.Insert(8)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkThreeElemsLeftRight(bst4, 10, 5, 8, t)

	var bst5 BST
	err = bst5.Insert(10)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst5.Insert(15)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst5.Insert(13)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkThreeElemsRightLeft(bst5, 10, 15, 13, t)

	//delete
	err = bst.Delete(10)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkTwoElemsRight(bst, 5, 80, t)

	err = bst2.Delete(30)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkTwoElemsRight(bst2, 10, 50, t)

	err = bst3.Delete(5)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkTwoElemsLeft(bst3, 10, 3, t)

	err = bst3.Insert(2)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst3.Delete(10)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkTwoElemsLeft(bst3, 3, 2, t)

	err = bst4.Delete(5)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkTwoElemsLeft(bst4, 10, 8, t)

	err = bst5.Delete(13)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	checkTwoElemsRight(bst5, 10, 15, t)

	//get
	bst.Delete(80)
	bst.Delete(5)
	checkNoElems(bst, t)

	bst.Insert(5)
	bst.Insert(80)
	bst.Insert(1)
	checkThreeElemsBalanced(bst, 5, 1, 80, t)

	min, err := bst.Minimum()
	if err != nil {
		t.Fatalf("Expected %v, Got %s", nil, err.Error())
	}
	if min != 1 {
		t.Errorf("Expected %d, Got %d", 1, min)
	}

	max, err := bst.Maximum()
	if err != nil {
		t.Fatalf("Expected %v, Got %s", nil, err.Error())
	}
	if max != 80 {
		t.Errorf("Expected %d, Got %d", 80, max)
	}

	next, err := bst.Next(5)
	if err != nil {
		t.Fatalf("Expected %v, Got %s", nil, err.Error())
	}
	if next != 80 {
		t.Errorf("Expected %d, Got %d", 80, next)
	}

	sorted := bst.Inorder()
	if len(sorted) != 3 {
		t.Fatalf("Expected %d, Got %d", 3, len(sorted))
	}
	if sorted[0] != 1 {
		t.Fatalf("Expected %d, Got %d", 1, sorted[0])
	}
	if sorted[1] != 5 {
		t.Fatalf("Expected %d, Got %d", 5, sorted[1])
	}
	if sorted[2] != 80 {
		t.Fatalf("Expected %d, Got %d", 80, sorted[2])
	}

	if !bst.Exists(80) {
		t.Errorf("Expected %t, Got %t", true, false)
	}
	if bst.Exists(10) {
		t.Errorf("Expected %t, Got %t", false, true)
	}
}

func TestManyElemsBST(t *testing.T) {
	//insert
	var bst BST
	err := bst.Insert(20)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst.Insert(25)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst.Insert(5)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst.Insert(10)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst.Insert(12)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst.Insert(14)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	err = bst.Insert(18)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if bst.Size != 7 {
		t.Errorf("Expected %d, Got %d", 7, bst.Size)
	}

	//delete
	err = bst.Delete(0)
	if err == nil {
		t.Errorf("Expected %s, Got %v", "value not found", err)
	}
	err = bst.Delete(70)
	if err == nil {
		t.Errorf("Expected %s, Got %v", "value not found", err)
	}

	err = bst.Delete(20)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if bst.Size != 6 {
		t.Errorf("Expected %d, Got %d", 6, bst.Size)
	}
	if bst.Root.Value != 18 {
		t.Errorf("Expected %d, Got %d", 18, bst.Root.Value)
	}
	if bst.Root.Left.Value != 5 {
		t.Errorf("Expected %d, Got %d", 5, bst.Root.Value)
	}

	err = bst.Delete(5)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if bst.Size != 5 {
		t.Errorf("Expected %d, Got %d", 5, bst.Size)
	}
	if bst.Root.Left.Value != 10 {
		t.Errorf("Expected %d, Got %d", 10, bst.Root.Value)
	}
	if bst.Root.Left.Left != nil {
		t.Errorf("Expected %v, Got %+v", nil, bst.Root.Left.Left)
	}

	err = bst.Insert(13)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}

	err = bst.Delete(18)
	if err != nil {
		t.Errorf("Expected %v, Got %s", nil, err.Error())
	}
	if bst.Root.Value != 14 {
		t.Errorf("Expected %d, Got %d", 14, bst.Root.Value)
	}
	if bst.Root.Left.Right.Right.Value != 13 {
		t.Errorf("Expected %d, Got %d", 13, bst.Root.Value)
	}

	var bst2 BST
	bst2.Insert(5)
	bst2.Insert(2)
	bst2.Insert(20)
	bst2.Insert(25)
	bst2.Insert(3)
	bst2.Insert(17)
	bst2.Insert(15)
	bst2.Insert(1)

	bst2.Delete(2)
	if bst2.Root.Value != 5 {
		t.Errorf("Expected %d, Got %d", 5, bst2.Root.Value)
	}
	if bst2.Root.Left.Value != 1 {
		t.Errorf("Expected %d, Got %d", 1, bst2.Root.Left.Value)
	}
	if bst2.Root.Left.Right.Value != 3 {
		t.Errorf("Expected %d, Got %d", 3, bst2.Root.Left.Right.Value)
	}

	bst2.Delete(20)
	if bst2.Root.Right.Value != 17 {
		t.Errorf("Expected %d, Got %d", 17, bst2.Root.Right.Value)
	}
	if bst2.Root.Right.Right.Value != 25 {
		t.Errorf("Expected %d, Got %d", 25, bst2.Root.Right.Right.Value)
	}
	if bst2.Root.Right.Left.Value != 15 {
		t.Errorf("Expected %d, Got %d", 15, bst2.Root.Right.Left.Value)
	}

	//get
	var bst3 BST
	bst3.Insert(5)
	bst3.Insert(2)
	bst3.Insert(20)
	bst3.Insert(3)
	bst3.Insert(17)
	bst3.Insert(1)

	min, err := bst3.Minimum()
	if err != nil {
		t.Fatalf("Expected %v, Got %s", nil, err.Error())
	}
	if min != 1 {
		t.Errorf("Expected %d, Got %d", 1, min)
	}

	max, err := bst3.Maximum()
	if err != nil {
		t.Fatalf("Expected %v, Got %s", nil, err.Error())
	}
	if max != 20 {
		t.Errorf("Expected %d, Got %d", 20, max)
	}

	next, err := bst3.Next(17)
	if err != nil {
		t.Fatalf("Expected %v, Got %s", nil, err.Error())
	}
	if next != 20 {
		t.Errorf("Expected %d, Got %d", 20, next)
	}

	sorted := bst3.Inorder()
	if len(sorted) != 6 {
		t.Fatalf("Expected %d, Got %d", 6, len(sorted))
	}
	if sorted[0] != 1 {
		t.Fatalf("Expected %d, Got %d", 1, sorted[0])
	}
	if sorted[1] != 2 {
		t.Fatalf("Expected %d, Got %d", 2, sorted[1])
	}
	if sorted[2] != 3 {
		t.Fatalf("Expected %d, Got %d", 3, sorted[2])
	}

	if !bst3.Exists(17) {
		t.Errorf("Expected %t, Got %t", true, false)
	}
}

func checkNoElems(bst BST, t *testing.T) {
	if bst.Size != 0 {
		t.Errorf("Expected %d, Got %d", 0, bst.Size)
	}
	if bst.Root != nil {
		t.Errorf("Expected %v, Got %+v", nil, bst.Root)
	}
}

func checkOneElem(bst BST, val int, t *testing.T) {
	if bst.Size != 1 {
		t.Errorf("Expected %d, Got %d", 1, bst.Size)
	}
	if bst.Root == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root, nil)
	}
	if bst.Root.Value != val {
		t.Errorf("Expected %d, Got %d", val, bst.Root.Value)
	}
	if bst.Root.Left != nil {
		t.Errorf("Expected %v, Got %+v", nil, bst.Root.Left)
	}
	if bst.Root.Right != nil {
		t.Errorf("Expected %v, Got %+v", nil, bst.Root.Right)
	}
}

func checkTwoElemsRight(bst BST, root, value int, t *testing.T) {
	if bst.Size != 2 {
		t.Errorf("Expected %d, Got %d", 2, bst.Size)
	}
	if bst.Root == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root, nil)
	}
	if bst.Root.Value != root {
		t.Errorf("Expected %d, Got %d", root, bst.Root.Value)
	}

	if bst.Root.Left != nil {
		t.Errorf("Expected %v, Got %+v", nil, bst.Root.Left)
	}
	if bst.Root.Right == nil {
		t.Fatalf("Expected %s, Got %v", "not nil", nil)
	}
	if bst.Root.Right.Value != value {
		t.Errorf("Expected %d, Got %d", value, bst.Root.Right.Value)
	}
}

func checkTwoElemsLeft(bst BST, root, value int, t *testing.T) {
	if bst.Size != 2 {
		t.Errorf("Expected %d, Got %d", 2, bst.Size)
	}
	if bst.Root == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root, nil)
	}
	if bst.Root.Value != root {
		t.Errorf("Expected %d, Got %d", root, bst.Root.Value)
	}

	if bst.Root.Right != nil {
		t.Errorf("Expected %v, Got %+v", nil, bst.Root.Right)
	}
	if bst.Root.Left == nil {
		t.Fatalf("Expected %s, Got %v", "not nil", nil)
	}
	if bst.Root.Left.Value != value {
		t.Errorf("Expected %d, Got %d", value, bst.Root.Left.Value)
	}
}

func checkThreeElemsBalanced(bst BST, root, left, right int, t *testing.T) {
	if bst.Size != 3 {
		t.Errorf("Expected %d, Got %d", 3, bst.Size)
	}
	if bst.Root == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root, nil)
	}
	if bst.Root.Value != root {
		t.Errorf("Expected %d, Got %d", root, bst.Root.Value)
	}
	if bst.Root.Left == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root.Left, nil)
	}
	if bst.Root.Left.Value != left {
		t.Errorf("Expected %d, Got %d", left, bst.Root.Left.Value)
	}
	if bst.Root.Right == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root.Right, nil)
	}
	if bst.Root.Right.Value != right {
		t.Errorf("Expected %d, Got %d", right, bst.Root.Right.Value)
	}
}

func checkThreeElemsRigth(bst BST, root, val1, val2 int, t *testing.T) {
	if bst.Size != 3 {
		t.Errorf("Expected %d, Got %d", 3, bst.Size)
	}
	if bst.Root == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root, nil)
	}
	if bst.Root.Value != root {
		t.Errorf("Expected %d, Got %d", root, bst.Root.Value)
	}
	if bst.Root.Left != nil {
		t.Fatalf("Expected %v, Got %+v", nil, bst.Root.Left)
	}
	if bst.Root.Right == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root.Right, nil)
	}
	if bst.Root.Right.Value != val1 {
		t.Errorf("Expected %d, Got %d", val1, bst.Root.Right.Value)
	}
	if bst.Root.Right.Left != nil {
		t.Fatalf("Expected %v, Got %+v", nil, bst.Root.Right.Left)
	}
	if bst.Root.Right.Right == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root.Right.Right, nil)
	}
	if bst.Root.Right.Right.Value != val2 {
		t.Errorf("Expected %d, Got %d", val2, bst.Root.Right.Right.Value)
	}
}

func checkThreeElemsLeft(bst BST, root, val1, val2 int, t *testing.T) {
	if bst.Size != 3 {
		t.Errorf("Expected %d, Got %d", 3, bst.Size)
	}
	if bst.Root == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root, nil)
	}
	if bst.Root.Value != root {
		t.Errorf("Expected %d, Got %d", root, bst.Root.Value)
	}
	if bst.Root.Right != nil {
		t.Fatalf("Expected %v, Got %+v", nil, bst.Root.Right)
	}
	if bst.Root.Left == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root.Left, nil)
	}
	if bst.Root.Left.Value != val1 {
		t.Errorf("Expected %d, Got %d", val1, bst.Root.Left.Value)
	}
	if bst.Root.Left.Right != nil {
		t.Fatalf("Expected %v, Got %+v", nil, bst.Root.Left.Right)
	}
	if bst.Root.Left.Left == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root.Left.Left, nil)
	}
	if bst.Root.Left.Left.Value != val2 {
		t.Errorf("Expected %d, Got %d", val2, bst.Root.Left.Left.Value)
	}
}

func checkThreeElemsLeftRight(bst BST, root, val1, val2 int, t *testing.T) {
	if bst.Size != 3 {
		t.Errorf("Expected %d, Got %d", 3, bst.Size)
	}
	if bst.Root == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root, nil)
	}
	if bst.Root.Value != root {
		t.Errorf("Expected %d, Got %d", root, bst.Root.Value)
	}
	if bst.Root.Right != nil {
		t.Fatalf("Expected %v, Got %+v", nil, bst.Root.Right)
	}
	if bst.Root.Left == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root.Left, nil)
	}
	if bst.Root.Left.Value != val1 {
		t.Errorf("Expected %d, Got %d", val1, bst.Root.Left.Value)
	}
	if bst.Root.Left.Left != nil {
		t.Fatalf("Expected %v, Got %+v", nil, bst.Root.Left.Left)
	}
	if bst.Root.Left.Right == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root.Left.Right, nil)
	}
	if bst.Root.Left.Right.Value != val2 {
		t.Errorf("Expected %d, Got %d", val2, bst.Root.Left.Right.Value)
	}
}

func checkThreeElemsRightLeft(bst BST, root, val1, val2 int, t *testing.T) {
	if bst.Size != 3 {
		t.Errorf("Expected %d, Got %d", 3, bst.Size)
	}
	if bst.Root == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root, nil)
	}
	if bst.Root.Value != root {
		t.Errorf("Expected %d, Got %d", root, bst.Root.Value)
	}
	if bst.Root.Left != nil {
		t.Fatalf("Expected %v, Got %+v", nil, bst.Root.Left)
	}
	if bst.Root.Right == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root.Right, nil)
	}
	if bst.Root.Right.Value != val1 {
		t.Errorf("Expected %d, Got %d", val1, bst.Root.Right.Value)
	}
	if bst.Root.Right.Right != nil {
		t.Fatalf("Expected %v, Got %+v", nil, bst.Root.Right.Right)
	}
	if bst.Root.Right.Left == nil {
		t.Fatalf("Expected %+v, Got %v", bst.Root.Right.Left, nil)
	}
	if bst.Root.Right.Left.Value != val2 {
		t.Errorf("Expected %d, Got %d", val2, bst.Root.Right.Left.Value)
	}
}

var NKEYS int = 1000

func key(i int) int {
	return NKEYS*((i*i-100*i)%NKEYS) + i
}

func TestStress(t *testing.T) {
	var bst BST
	var err error

	// Insert
	for i := 0; i < NKEYS; i++ {
		k := key(i)
		if bst.Size != i {
			t.Errorf("Expected %d, Got %d", i, bst.Size)
		}
		if bst.Exists(k) {
			t.Errorf("Expected %t, Got %t", false, true)
		}
		bst.Insert(k)
		if !bst.Exists(k) {
			t.Errorf("Expected %t, Got %t", true, false)
		}
	}

	if bst.Size != NKEYS {
		t.Errorf("Expected %d, Got %d", NKEYS, bst.Size)
	}

	// Insert again
	for i := 0; i < NKEYS; i++ {
		k := key(i)
		if bst.Size != NKEYS {
			t.Errorf("Expected %d, Got %d", i, bst.Size)
		}
		if !bst.Exists(k) {
			t.Errorf("Expected %t, Got %t", true, false)
		}

		err = bst.Insert(k)
		if err == nil {
			t.Fatalf("Expected %s, Got %v", "value already exists", err)
		}
		if err.Error() != "value already exists" {
			t.Fatalf("Expected %s, Got %v", "value already exists", err.Error())
		}
		if !bst.Exists(k) {
			t.Errorf("Expected %t, Got %t", true, false)
		}
		if bst.Size != NKEYS {
			t.Errorf("Expected %d, Got %d", NKEYS, bst.Size)
		}
	}

	// Delete even values
	for i := 0; i < NKEYS; i++ {
		k := key(i)
		if !bst.Exists(k) {
			t.Errorf("Expected %t, Got %t", true, false)
		}
		if i%2 == 0 {
			bst.Delete(k)
			if bst.Exists(k) {
				t.Errorf("Expected %t, Got %t", false, true)
			}
		}
	}
	if bst.Size != NKEYS/2 {
		t.Errorf("Expected %d, Got %d", NKEYS/2, bst.Size)
	}

	// Delete odd values
	for i := 0; i < NKEYS; i++ {
		k := key(i)
		if i%2 == 0 {
			if bst.Exists(k) {
				t.Errorf("Expected %t, Got %t", false, true)
			}
		} else {
			if !bst.Exists(k) {
				t.Errorf("Expected %t, Got %t", true, false)
			}
			bst.Delete(k)
			if bst.Exists(k) {
				t.Errorf("Expected %t, Got %t", false, true)
			}
		}
	}
	if bst.Size != 0 {
		t.Errorf("Expected %d, Got %d", 0, bst.Size)
	}

	// Verify no elems left
	for i := 0; i < NKEYS; i++ {
		k := key(i)
		if bst.Exists(k) {
			t.Errorf("Expected %t, Got %t", false, true)
		}
	}
	checkNoElems(bst, t)

}
