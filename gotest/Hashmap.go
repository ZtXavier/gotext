package main

import (
	"fmt"
)

type HashMap struct {
	key string
	value string
	hashcode int
	next *HashMap
}

var table [16](*HashMap)

func initTable() {
	for i := range table {
		table[i] = &HashMap{"","",i,nil}
	}
}

func getInstance()[16] (*HashMap){
	if(table[0] == nil){
		initTable()
	}
	return table
}

func getHashCode(k string) int{
	if len(k) == 0{
		return 0
	}
	var hashCode int = 0
	var lastIndex int = len(k) - 1
	for i := range k {
		if i == lastIndex {
			hashCode += int(k[i])
			break
		}
		hashCode += (hashCode + int(k[i]))*31
	}
	return hashCode
}

func indexTable(hashCode int) int{
	return hashCode % 16
}

func indexNode(hashCode int) int{
	return hashCode >> 4
}


func put(k string,v string) string{
	var hashCode = getHashCode(k)
	var thisNode = HashMap{k,v,hashCode,nil}
	var tableIndex = indexTable(hashCode)
	var nodeIndex = indexNode(hashCode)
	var headPtr [16](*HashMap) = getInstance()
	var headNode = headPtr[tableIndex]

	if(*headNode).key == ""{
		*headNode = thisNode;
		return ""
	}

	var lastNode *HashMap = headNode
	var nextNode *HashMap = (*headNode).next

	for nextNode != nil && (indexNode((*nextNode).hashcode) < nodeIndex){
		lastNode = nextNode
		nextNode = (*nextNode).next
	}
	if(*lastNode).hashcode == thisNode.hashcode {
		var oldValue string = lastNode.value
		lastNode.value = thisNode.value
		return oldValue;
	}
	if lastNode.hashcode < thisNode.hashcode {
		lastNode.next = &thisNode
	}
	if nextNode != nil {
		thisNode.next = nextNode
	}
	return ""
}

func get(k string) string {
	var hashCode = getHashCode(k)
	var tableIndex = indexTable(hashCode)

	var headPtr [16](*HashMap) = getInstance()
	var node *HashMap = headPtr[tableIndex]
	
	if(*node).next != nil {
		return (*node).value 
	}
	for (*node).next != nil {
		if k == (*node).key {
			return (*node).value
		} 
		node = (*node).next
	}
	return ""
}


func main() {
	getInstance()
	put("a","a-v")
	put("b","b_v")
	fmt.Println(get("a"))
	fmt.Println(get("b"))
	put("p","p-v")
	fmt.Println(get("p"))
}


