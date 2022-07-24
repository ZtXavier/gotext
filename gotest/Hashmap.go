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

// 定义一个哈希表
var table [16](*HashMap)

// 初始化一个哈希表
func initTable() {
	for i := range table {
		table[i] = &HashMap{"","",i,nil}
	}
}

// 获得一个初始化好的哈希table
func getInstance() [16] (*HashMap){
	if(table[0] == nil){
		initTable()
	}
	return table
}

// 这里是获得一个哈希数
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

// 获得索引值
func indexTable(hashCode int) int{
	return hashCode % 16
}

// 获得哈希节点
func indexNode(hashCode int) int{
	return hashCode >> 4
}

// 向哈希表中添加数据
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

// 获得哈希表中的数据
func get(k string) string {
	var hashCode = getHashCode(k)
	var tableIndex = indexTable(hashCode)

	var headPtr [16](*HashMap) = getInstance()
	var node *HashMap = headPtr[tableIndex]
	
	if(*node).key == k {
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


