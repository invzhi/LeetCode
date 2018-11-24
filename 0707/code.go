// 707. Design Linked List

// Design your implementation of the linked list. You can choose to use the singly linked list or the doubly linked list. A node in a singly linked list should have two attributes: val and next. val is the value of the current node, and next is a pointer/reference to the next node. If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node in the linked list. Assume all nodes in the linked list are 0-indexed.

// Implement these functions in your linked list class:
// - get(index) : Get the value of the index-th node in the linked list. If the index is invalid, return -1.
// - addAtHead(val) : Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
// - addAtTail(val) : Append a node of value val to the last element of the linked list.
// - addAtIndex(index, val) : Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
// - deleteAtIndex(index) : Delete the index-th node in the linked list, if the index is valid.

// Example:
// MyLinkedList linkedList = new MyLinkedList();
// linkedList.addAtHead(1);
// linkedList.addAtTail(3);
// linkedList.addAtIndex(1, 2);  // linked list becomes 1->2->3
// linkedList.get(1);            // returns 2
// linkedList.deleteAtIndex(1);  // now the linked list is 1->3
// linkedList.get(1);            // returns 3

// Note:
// - All values will be in the range of [1, 1000].
// - The number of operations will be in the range of [1, 1000].
// - Please do not use the built-in LinkedList library.

package leetcode

type listNode struct {
	val  int
	next *listNode
}

// MyLinkedList is my implementation of the linked list.
type MyLinkedList struct {
	head   *listNode
	tail   *listNode
	length int
}

// Constructor initialize your data structure here.
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// Get get the value of the index-th node in the linked list. If the index is invalid, return -1.
func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.length {
		return -1
	}
	if index == l.length-1 {
		return l.tail.val
	}

	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.val
}

// AddAtHead add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
func (l *MyLinkedList) AddAtHead(val int) {
	head := &listNode{val: val}
	head.next = l.head
	l.head = head
	if l.tail == nil {
		l.tail = head
	}
	l.length++
}

// AddAtTail append a node of value val to the last element of the linked list.
func (l *MyLinkedList) AddAtTail(val int) {
	tail := &listNode{val: val}
	if l.tail != nil {
		l.tail.next = tail
	} else {
		l.head = tail
	}
	l.tail = tail
	l.length++
}

// AddAtIndex add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		l.AddAtHead(val)
		return
	}
	if index == l.length {
		l.AddAtTail(val)
		return
	}
	if index < 0 || index > l.length {
		return
	}

	var prev, node *listNode
	node = l.head
	for i := 0; i < index; i++ {
		prev = node
		node = node.next
	}
	prev.next = &listNode{val: val, next: node}
	l.length++
}

// DeleteAtIndex delete the index-th node in the linked list, if the index is valid.
func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.length {
		return
	}

	var prev, node *listNode
	node = l.head
	for i := 0; i < index; i++ {
		prev = node
		node = node.next
	}
	if prev != nil {
		prev.next = node.next
	} else {
		l.head = node.next
	}
	if node.next != nil {
		node.next = nil
	} else {
		l.tail = prev
	}
	l.length--
}
