package main

import (
	"errors"
	"fmt"
	"log"
)

type Node struct {
	Data interface{}
	Next *Node
}

// 创建头结点
// @return data 被预设为 "head" 的头结点
func NewLink() *Node {
	return &Node{
		Data: "head",
		Next: nil,
	}
}

// 追加节点，该方法传参较为繁琐：link.Append(&Node{Data: 123,})
// @param  node struct point
// @return 返回新的尾部 node（即参数）
func (n *Node) append(node *Node) *Node {
	n.Next = node
	return node
}

// 更简洁的追加方法，只需传入值即可
// @param v 任意值
// @return 返回新的尾部 node
func (n *Node) Append(v interface{}) *Node {
	next := new(Node)
	next.Data = v
	n.Next = next
	return next
}

// 在链表的指定位置插入
// @param index 索引
// @param v 插入值
// @return error
func (n *Node) Insert(index int, v interface{}) error{
	N := n

	if index > n.Len() {
		return errors.New("index is wrong")
	}
	// 移动到插入位置的前一个节点
	for i := 0; i < index-1; i++ {
		// fmt.Println(N.Data)
		N = N.Next
	}
	// 保存后一个节点
	nextNode := N.Next

	// 插入的节点
	node := new(Node)
	node.Data = v
	// 指向之前保存的后一个节点
	node.Next = nextNode

	N.Next = node
	return nil
}

// 获取指定索引的 data ，从 1 开始，0 是头结点
// @param index 索引
// @return 值
func (n *Node) Get(index int) (interface{}, error) {
	var value interface{}
	N := n

	if index > n.Len() {
		return nil, errors.New("index is wrong")
	}
	for i := 0; i < index+1; i++ {
		value = N.Data
		N = N.Next
	}
	return value, nil
}

// 移除节点
// @param index 索引
// @return error
func (n *Node) Remove(index int) error {
	N := n
	if index > n.Len() {
		return errors.New("index is wrong")
	}
	if index == 0 {
		return errors.New("you can't remove head node")
	}

	for i := 0; i < index-1; i++ {
		N = N.Next
	}
	next := N.Next.Next
	N.Next = next
	return nil
}

// 获取链表长度，不包括头结点
// @return 长度
func (n *Node) Len() int {
	l := 0
	N := n
	for N.Next != nil {
		l++
		N = N.Next
	}
	return l
}

// 遍历链表
func (n *Node) Range() {
	N := n
	for N.Next != nil {
		fmt.Printf("%v -> ", N.Data)
		N = N.Next
	}
	// 输出最后一个节点的值
	fmt.Println(N.Data)
}

func main() {
	link := NewLink()

	link.Append(123).Append(456).Append("aabfa")

	//v, _ := link.Get(3)
	//fmt.Println(v)
	//
	//l := link.Len()
	//fmt.Println(l)
	link.Insert(3, "fawwrw")
	link.Range()
	fmt.Println(link.Len())
	err := link.Remove(3)
	if err != nil {
		log.Fatal(err)
	}
	link.Range()

}
