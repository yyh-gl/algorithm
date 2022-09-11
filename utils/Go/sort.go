package utils

import (
	"sort"
)

type StringArray struct {
	Val   []string
	Order string
}

func NewStringArray(arr []string, order string) *StringArray {
	return &StringArray{Val: arr, Order: order}
}

func (sa *StringArray) Len() int {
	return len(sa.Val)
}

func (sa *StringArray) Less(i, j int) bool {
	if sa.Order == "asc" {
		return sa.Val[i] < sa.Val[j]
	}
	return sa.Val[i] > sa.Val[j]
}

func (sa *StringArray) Swap(i, j int) {
	sa.Val[i], sa.Val[j] = sa.Val[j], sa.Val[i]
}

func sortStringAsc(arr []string) []string {
	sArr := NewStringArray(arr, "asc")
	sort.Sort(sArr)
	return sArr.Val
}

func sortStringDesc(arr []string) []string {
	sArr := NewStringArray(arr, "desc")
	sort.Sort(sArr)
	return sArr.Val
}

type IntArray struct {
	Val   []int
	Order string
}

func NewIntArray(arr []int, order string) *IntArray {
	return &IntArray{Val: arr, Order: order}
}

func (ia *IntArray) Len() int {
	return len(ia.Val)
}

func (ia *IntArray) Less(i, j int) bool {
	if ia.Order == "asc" {
		return ia.Val[i] < ia.Val[j]
	}
	return ia.Val[i] > ia.Val[j]
}

func (ia *IntArray) Swap(i, j int) {
	ia.Val[i], ia.Val[j] = ia.Val[j], ia.Val[i]
}

func sortIntAsc(arr []int) []int {
	iArr := NewIntArray(arr, "asc")
	sort.Sort(iArr)
	return iArr.Val
}

func sortIntDesc(arr []int) []int {
	iArr := NewIntArray(arr, "desc")
	sort.Sort(iArr)
	return iArr.Val
}
