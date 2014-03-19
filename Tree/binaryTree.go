package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	left  *Tree
	right *Tree
	value int
}

//Traverse depth-first
func walk(t *Tree, channel chan int) {
	if t == nil {
		return
	}
	walk(t.left, channel)
	channel <- t.value // puts t.value on the channel.
	walk(t.right, channel)
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func walker(t *Tree) <-chan int {
	channel := make(chan int) //a channel that can send and receive an int
	go func() {
		walk(t, channel)
		close(channel)
	}()
	return channel
}

// Compare reads values from two Walkers
// that run simultaneously, and returns true
// if tree1 and tree2 have the same contents.
func compare(tree1, tree2 *Tree) bool {
	channel1, channel2 := walker(tree1), walker(tree2)
	for {
		value1, ok1 := <-channel1
		value2, ok2 := <-channel2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if value1 != value2 {
			break
		}
	}
	return false
}

func newRndTree(num, k int) *Tree {
	var tree *Tree
	for _, v := range rand.Perm(num) { // Perm return a slice of ints
		tree = insert(tree, (1+v)*k)
	}
	return tree
}

func insert(tree *Tree, value int) *Tree {
	if tree == nil {
		return &Tree{nil, nil, value}
	}
	if value < tree.value {
		tree.left = insert(tree.left, value)
		return tree
	}
	tree.right = insert(tree.right, value)
	return tree
}

func main() {
	tree := newRndTree(100, 1)
	fmt.Println(compare(tree, newRndTree(100, 1)), "same contents")
	fmt.Println(compare(tree, newRndTree(99, 1)), "Differing Sizes")
	fmt.Println(compare(tree, newRndTree(100, 2)), "Differing Values")
	fmt.Println(compare(tree, newRndTree(101, 2)), "Dissimilar")
}
