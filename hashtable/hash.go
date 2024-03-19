package main

import "fmt"

// we need a array size for creat a hash func , 
//because  of we are using the  division method, so that why.. 


const ArraySize = 10


// so we need to create a hash table, 
// the array name [array size] array ponit to a bucket .. 

type HashTable struct {
	Array [ArraySize]*Bucket
}


// here a bucket , the array pointed to these bucket, 
// when element add so it will came here .. 


type Bucket struct {
	Head *BucketNode
}


// bucket node 

type BucketNode struct {
	Key  string
	Next *BucketNode
}

// here we are create the hash table,
// add node to the all element of array 


func Init() *HashTable {

	result := &HashTable{}

	for i := range result.Array {

		result.Array[i] = &Bucket{}

	}
	return result

}

//here we need a value
// frist we need hash the value, that time we got an index
// then we need to search that index is avaible or not

func (h1 *HashTable) Insert(key string) {

	index := h1.hash(key)
	h1.Array[index].insert(key)

}


// here we are adding the value 

func (b *Bucket) insert(k string) {

	addNode := &BucketNode{Key: k}

	if !b.search(k) {

		addNode.Next = b.Head

		b.Head = addNode

	} else {

		fmt.Println("already exist ..")
	}
}

//hash the value
// here we got index from hashed 

func (h1 *HashTable) hash(k string) int {

	sum := 0

	for _, val := range k {

		sum += int(val)
	}

	return sum % len(h1.Array)
}

// search

/

func (h1 *HashTable) Serach(key string) bool {

	index := h1.hash(key)

	return h1.Array[index].search(key)

}

func (b *Bucket) search(k string) bool {

	currentNode := b.Head

	for currentNode != nil {

		if currentNode.Key == k {

			return true
		}
		currentNode = currentNode.Next
	}

	return false

}

func (h1 *HashTable) Delete(key string) {

	index := h1.hash(key)

	h1.Array[index].delete(key)

}

func (b *Bucket) delete(k string)  {

	current := b.Head

	if b.Head.Key == k{
		b.Head = b.Head.Next
		return 
	}

	for current != nil {

		if current.Next.Key == k{
			current.Next = current.Next.Next
			
		}

		current = current.Next

	}

	

}

func main() {

	h1 := Init()

	list := []string{
		"akhil",
		"arun",
		"amal",
		"aswan",
		"remesh",
		"ramu",
	}

	for _, val := range list {

		h1.Insert(val)
	}
	
	h1.Delete("ramu")
	fmt.Println(h1.Serach("ramu"))


	fmt.Println(h1.Serach("arun"))

}
