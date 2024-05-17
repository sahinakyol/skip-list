package main

import (
	"fmt"
	"math/rand"
)

const MaxLevel = 16

type Node struct {
	key   int         // The key used for ordering.
	value interface{} // The value associated with the key.
	level []*Node     // Slice of pointers to nodes at different levels.
}

type SkipList struct {
	head  *Node // head node of the skip list.
	level int   // Current maximum level of the skip list.
}

func NewNode(key int, value interface{}, level int) *Node {
	return &Node{
		key:   key,
		value: value,
		level: make([]*Node, level),
	}
}

func NewSkipList() *SkipList {
	// Initialize the head node with key -1 and maximum level.
	head := NewNode(-1, nil, MaxLevel)
	return &SkipList{
		head:  head,
		level: 1,
	}
}

// randomLevel generates a random level for a new node.
func randomLevel() int {
	level := 1
	// Randomly increase the level with a 50% chance.
	for rand.Float32() < 0.5 && level < MaxLevel {
		level++
	}
	return level
}

func (sl *SkipList) Insert(key int, value interface{}) {
	update := make([]*Node, MaxLevel)
	current := sl.head

	// Traverse the skip list to find the position to insert the new key-value pair.
	for i := sl.level - 1; i >= 0; i-- {
		for current.level[i] != nil && current.level[i].key < key {
			current = current.level[i]
		}
		update[i] = current
	}

	// Move to the next node at the lowest level.
	current = current.level[0]

	// If the current node is nil or its key is different, insert a new node.
	if current == nil || current.key != key {
		newLevel := randomLevel()

		// Update the list level if the new node's level is higher.
		if newLevel > sl.level {
			for i := sl.level; i < newLevel; i++ {
				update[i] = sl.head
			}
			sl.level = newLevel
		}

		// Create the new node.
		newNode := NewNode(key, value, newLevel)

		// Update the pointers to insert the new node.
		for i := 0; i < newLevel; i++ {
			newNode.level[i] = update[i].level[i]
			update[i].level[i] = newNode
		}
	}
}

func (sl *SkipList) Search(key int) *Node {
	current := sl.head

	// Traverse the skip list to find the key.
	for i := sl.level - 1; i >= 0; i-- {
		for current.level[i] != nil && current.level[i].key < key {
			current = current.level[i]
		}
	}

	// Move to the next node at the lowest level.
	current = current.level[0]

	// Check if the current node contains the key.
	if current != nil && current.key == key {
		return current
	}
	return nil
}

func (sl *SkipList) Delete(key int) {
	update := make([]*Node, MaxLevel)
	current := sl.head

	// Traverse the skip list to find the position of the node to delete.
	for i := sl.level - 1; i >= 0; i-- {
		for current.level[i] != nil && current.level[i].key < key {
			current = current.level[i]
		}
		update[i] = current
	}

	// Move to the next node at the lowest level.
	current = current.level[0]

	// If the current node contains the key, update pointers to remove it.
	if current != nil && current.key == key {
		for i := 0; i < sl.level; i++ {
			if update[i].level[i] != current {
				break
			}
			update[i].level[i] = current.level[i]
		}

		// Decrease the list level if necessary.
		for sl.level > 1 && sl.head.level[sl.level-1] == nil {
			sl.level--
		}
	}
}

func (sl *SkipList) Display() {
	for i := 0; i < sl.level; i++ {
		current := sl.head.level[i]
		fmt.Printf("Level %d: ", i)
		for current != nil {
			fmt.Printf("%d:%v ", current.key, current.value)
			current = current.level[i]
		}
		fmt.Println()
	}
}
