package sort

import (
	"sort"
)

type StringArray struct {
	Val []string
}

func NewStringArray(arr []string) *StringArray {
	return &StringArray{Val: arr}
}

func (sa *StringArray) Len() int {
	return len(sa.Val)
}

func (sa *StringArray) Less(i, j int) bool {
	// asc
	return sa.Val[i] < sa.Val[j]
}

func (sa *StringArray) Swap(i, j int) {
	sa.Val[i], sa.Val[j] = sa.Val[j], sa.Val[i]
}

func sortString(arr []string) []string {
	sArr := NewStringArray(arr)
	sort.Sort(sArr)
	return sArr.Val
}

type IntArray struct {
	Val []int
}

func NewIntArray(arr []int) *IntArray {
	return &IntArray{Val: arr}
}

func (ia *IntArray) Len() int {
	return len(ia.Val)
}

func (ia *IntArray) Less(i, j int) bool {
	// asc
	return ia.Val[i] < ia.Val[j]
}

func (ia *IntArray) Swap(i, j int) {
	ia.Val[i], ia.Val[j] = ia.Val[j], ia.Val[i]
}

func sortInt(arr []int) []int {
	iArr := NewIntArray(arr)
	sort.Sort(iArr)
	return iArr.Val
}
