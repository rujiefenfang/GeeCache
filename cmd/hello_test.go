package main

import "testing"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func TestName(t *testing.T) {
	level := 0
	num := []string{"0", "-3", "9", "-10", "null", "5"}
	sign := 0
	GetTree(level, num, sign)

}

func GetTree(level int, num []string, sign int) {

}
