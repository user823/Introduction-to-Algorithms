package main

import (
	"fmt"
	"unsafe"
)

/*
*使用一个指针实现双向链表
*/


func main() {
	l := createList([]int{1, 2, 3, 4, 5})
	printList(l)
	tmp := node{6, 0}
	printList(insertList(l, &tmp))    //插入后
	printList(deleteList(l, &tmp))    //删除后
}

type node struct {
	data int
	np uintptr //np = pre XOR next 
}

func createList(nums []int) *node{
	length := len(nums)
	if length > 0{
		head := &node{nums[0], 0} //pre and next = 0
		pre := head         
		var l uintptr // l 表示 pre的pre
		for i:=1; i<length ;i++ {    
			now := &node{nums[i], addr(pre)} //从末尾插入当前节点, next = 0
			pre.np  = l ^ addr(now)       
			l = addr(pre)
			pre = now
		}
		return head
	}
	return nil
}	

func searchList(L *node, k int) *node{
	if L == nil {
		return nil
	}
	var l uintptr
	var pre = L 
	if pre.data == k {
		return pre
	}
	var now = getNode(l ^ pre.np)
	for ; now!=nil&&now.data!=k; now=getNode(l ^ now.np) {
		l = addr(pre)
		pre = now
	}
	return now
}

func insertList(L *node, n *node) *node{
	if L == nil {
		L = n
	}
	var l uintptr
	var pre = L 
	var now = getNode(l ^ pre.np)
	for ; now!=nil; now=getNode(l ^ pre.np) {
		l = addr(pre)
		pre = now
	}
	n.np = addr(pre)
	pre.np = l ^ addr(n)
	return L
}

func deleteList(L *node, n *node) *node{
	var l uintptr
	var pre = L 
	var now = getNode(l ^ pre.np)
	if pre == n {
		return now
	}
	for ; now!=nil&&now!=n; now=getNode(l ^ pre.np){ //定位到目标处
		l = addr(pre)
		pre = now
	}	
	if now != nil {         //找到目标
		pre.np = l ^ (now.np ^ addr(pre))
		next := getNode(now.np ^ addr(pre))
		if next != nil{
			next.np = addr(pre) ^ (addr(now) ^ next.np)
		}
	}
	return L
}

func printList(L *node) {
	if L == nil {
		return
	}
	fmt.Printf("[")
	var l uintptr
	var pre = L 
	fmt.Printf("%d, ", pre.data)
	for now:=getNode(pre.np ^ l); now!=nil; now=getNode(pre.np ^ l) {
		fmt.Printf("%d, ", now.data)
		l = addr(pre)
		pre = now
	}
	fmt.Printf("]\n")
}

func addr(n *node) uintptr {
	return uintptr(unsafe.Pointer(n))
}

func getNode(ad uintptr) *node {
	return (*node)(unsafe.Pointer(ad))
}
