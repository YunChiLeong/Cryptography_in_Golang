/*
Yun Chi Leong
COSC 4010 03 Fall 2021
Assignment  1
September 10 2021
*/

package hashtable

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

const ArraySize int = 65536 //2^16
type Hashtable struct {
	array [ArraySize]*bucket
}

// Separate chaining to handle collision
type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key   string
	value int
	next  *bucketNode
}

//New creates a bucket in each slot of the hash table
func New() *Hashtable {
	result := &Hashtable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

/*Insert inserts a new key/value pair into the hashtable,
return an error if the key already exists.*/
func (h *Hashtable) Insert(key string, value int) error {
	index := hash(key)
	//if key does not exist, insert key
	if h.array[index].search(key) == false {
		h.array[index].insert(key, value)
	} else {
		return errors.New("The key already exist.")
	}
	return nil
}

/* Update updates an existing key to be associated with a different value,
returns an error if the key doesn't already exist.*/
func (h *Hashtable) Update(key string, value int) error {
	index := hash(key)
	if h.array[index].search(key) == true {
		h.array[index].update(key, value)
	} else {
		return errors.New("The key does not exist.")
	}
	return nil
}

/* Delete deletes a key/value pair from the hashtable,
returns an error if the given key doesn't exist.*/
func (h *Hashtable) Delete(key string) error {
	index := hash(key)
	if h.array[index].search(key) == true {
		h.array[index].delete(key)
	} else {
		return errors.New("The key could be deleted because it does not exist.")
	}
	return nil
}

// Exists returns true if the key exists in the hashtable, false otherwise.
func (h *Hashtable) Exists(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

/* Get returns the value associated with the given key,
returns an error if value doesn't exist.*/
func (h *Hashtable) Get(key string) (int, error) {
	index := hash(key)
	//Check if key exists
	if h.array[index].search(key) == true {
		return h.array[index].get(key), nil
	} else {
		return 0, errors.New("The key does not exist.")
	}
}

//hash takes a key, encrypt it with sha256 and turn first two bytes of the hash into index
func hash(key string) int {
	// hash: [32]uint8
	hash := sha256.Sum256([]byte(key))
	// truncHash: uint16
	truncHash := (uint16(hash[0]) << 8) + uint16(hash[1])
	index := int(truncHash)
	return index
}

//insert takes a key and value and insert them in the bucket
func (b *bucket) insert(k string, v int) {
	if !b.search(k) {
		newNode := &bucketNode{key: k, value: v}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists.")
	}
}

//seach takes a key and return true if the bucket has the key
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

//delete takes a key and unlink it from the linked list bucket
func (b *bucket) delete(k string) {
	//if the node to be deleted is the head
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
			return
		}
		previousNode = previousNode.next
	}
}

//get takes a string, find the node in the bucket and return its value
func (b *bucket) get(k string) int {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return currentNode.value
		}
		currentNode = currentNode.next
	}
	return 0
}

//update takes a string and a value, find the node of key and change the value to new value
func (b *bucket) update(k string, new_v int) {

	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			currentNode.value = new_v
		}
		currentNode = currentNode.next
	}
}
