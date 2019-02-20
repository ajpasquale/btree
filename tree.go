package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// LevelOne is a dick
const (
	LevelTwo = 3*1<<iota - 1
	LevelThree
	LevelFour
	LevelFive
	LevelSix
	LevelSeven
	LevelOne = 0
)

//Node is a structure to hold a binary tree
type Node struct {
	value int
	left  *Node
	right *Node
}

// Insert into the binary tree
func Insert(root *Node, v int) *Node {
	if root == nil {
		root = &Node{
			value: v,
			left:  nil,
			right: nil,
		}
	} else if root.value > v {
		root.left = Insert(root.left, v)
	} else {
		root.right = Insert(root.right, v)
	}
	return root
}

//Traverse the tree
func Traverse(root *Node) {
	if root == nil {
		return
	}
	Traverse(root.left)
	Traverse(root.right)
}

//Size returns (number of nodes) in the tree
func Size(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + Size(root.left) + Size(root.right)
}

//Height returns the height of the tree
func Height(root *Node) int {
	if root == nil {
		return 0
	}
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	return max(Height(root.left), Height(root.right)) + 1
}

//Find will return a value in the tree
func Find(root *Node, v int) int {
	for {
		if root == nil {
			break
		}
		if root.value > v {
			fmt.Printf("Value: %d is less than root value: %d. Going left.\n", v, root.value)
			root = root.left
		} else if root.value < v {
			fmt.Printf("Value: %d is greater than root value: %d. Going right.\n", v, root.value)
			root = root.right
		} else {
			return root.value
		}
	}
	return 0
}

//Print prints a graphical representation of the tree :)
func Print(root *Node) {
	ts := Height(root)
	var levels []int
	var spaces []string
	var branches []string

	for i := 0; i < ts; i++ {
		j := float64(i)
		levels = append(levels, int(math.Pow(2, j)))
	}
	for i := ts - 1; i >= 0; i-- {
		j := float64(i)
		//spaces = append(spaces, 3*int(math.Pow(2, j)))
		spaces = append(spaces, strings.Repeat(" ", 3*int(math.Pow(2, j))-1))
		branches = append(branches, strings.Repeat("-", 3*int(math.Pow(2, j))-1))
	}
}

func toDash(h int) []string {
	var dashes []string
	//var prevPatternLen int
	for i := h; i >= 0; i-- {
		var tDash string
		var tSpace string
		j := float64(i)
		patternLen := 3*int(math.Pow(2, j)) - 1
		if patternLen < 3 {
			patternLen++
		}
		tDash = strings.Repeat("-", patternLen)
		tSpace = strings.Repeat(" ", patternLen)

		tDash = replaceAtIndex(tDash, '|', len(tDash)/2)
		tDash = tSpace + tDash
		dashes = append(dashes, tDash)
		//prevPatternLen = patternLen
	}
	return dashes
}
func replaceAtIndex(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

// when looking at the output visually it seems to produce an inverse
// of x squared? maybe useful for optimization?
func spacingFactor(h int) []int {
	var sf []int
	for i := 0; i < h-1; i++ {
		var tf int
		if i == 0 {
			sf = append(sf, 0)
		}
		f := float64(i)
		tf = 3*int(math.Pow(2, f)) - 1
		sf = append(sf, tf)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sf)))
	return sf
}
func spacing() []int {
	var levels = []int{LevelOne, LevelTwo, LevelThree,
		LevelFour, LevelFive, LevelSix, LevelSeven}
	return levels
}
func main() {

	sf := spacingFactor(7)
	sp := spacing()
	fmt.Println(sf)
	fmt.Println(sp)

}
