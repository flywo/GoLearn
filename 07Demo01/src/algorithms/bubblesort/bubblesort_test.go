package bubblesort

import "testing"

func TestBubbleSort1(t *testing.T) {
	values := []int{5,4,3,2,1}
	BubbleSort(values)
	if values[0]!=1 || values[1] != 2 || values[3] != 3 || values[4] != 4 || values[5] != 5 {
		t.Error("没有通过，正确顺序：12345，实际顺序：", values)
	}
}

func TestBubbleSort2(t *testing.T) {
	values := []int{5,5,3,2,1}
	BubbleSort(values)
	if values[0]!=1 || values[1] != 2 || values[3] != 3 || values[4] != 5 || values[5] != 5 {
		t.Error("没有通过，正确顺序：12355，实际顺序：", values)
	}
}

func TestBubbleSort3(t *testing.T) {
	values := []int{5}
	BubbleSort(values)
	if values[0] != 5 {
		t.Error("没有通过")
	}
}
