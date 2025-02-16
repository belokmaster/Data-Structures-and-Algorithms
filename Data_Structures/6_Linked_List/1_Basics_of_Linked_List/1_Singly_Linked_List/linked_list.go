package linkedList

import (
	"fmt"
)

// Односвязный список - это фундаментальная структура данных, состоящая из узлов,
// где каждый узел содержит поле данных и ссылку на следующий узел в связанном списке.
type Node struct {
	Val  int
	Next *Node
}

// Обход включает посещение каждого узла в связанном списке
func TraverseLinkedList(head *Node) {
	current := head

	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}

	fmt.Println()
}

func SearchLinkedList(head *Node, target int) bool {
	current := head

	for current != nil {
		if current.Val == target {
			return true
		}
		current = current.Next
	}

	return false
}

func LengthLinkedList(head *Node) int {
	length := 0
	current := head

	for current != nil {
		length++
		current = current.Next
	}

	return length
}

func InsertAtBeginning(head *Node, value int) *Node {
	newNode := &Node{Val: value}
	newNode.Next = head
	head = newNode
	return head
}

func InsertAtEnd(head *Node, value int) *Node {
	newNode := &Node{Val: value}
	// Если список пустой, новый узел становится головой
	if head == nil {
		return newNode
	}

	current := head
	// Проходим по всем узлам, пока не найдем последний
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
	return head
}

func InsertAtIndexPosition(head *Node, index, value int) *Node {
	if index < 0 {
		fmt.Println("Некорректный индекс.")
		return head
	}

	newNode := &Node{Val: value}
	if index == 0 {
		newNode.Next = head
		head = newNode
		return head
	}

	current := head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Next
	}

	if current == nil {
		fmt.Println("Индекс вышел за рамки.")
		return head
	}

	newNode.Next = current.Next
	current.Next = newNode

	return head
}

func DeleteAtBeginning(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Переносим указатель головы на следующий элемент
	head = head.Next

	// Возвращаем новый указатель на голову списка
	return head
}

func DeleteAtEnd(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Если список состоит только из одного элемента
	if head.Next == nil {
		return nil
	}

	current := head
	for current.Next != nil && current.Next.Next != nil {
		current = current.Next
	}

	current.Next = nil
	return head
}

func DeleteAtIndexPosition(head *Node, index int) *Node {
	if head == nil || index < 0 {
		return nil
	}

	if index == 0 {
		head = head.Next
		return head
	}

	current := head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Next
	}

	if current == nil || current.Next == nil {
		fmt.Println("Индекс вышел за рамки.")
		return head
	}

	// Удаление узла на индексе
	current.Next = current.Next.Next

	return head
}

func DeleteAllLinkedList(head *Node) *Node {
	current := head
	// Проходим по всем узлам и удаляем их
	for current != nil {
		nextNode := current.Next
		current = nil      // Убираем ссылку на текущий узел
		current = nextNode // Переходим к следующему узлу
	}
	return nil // Возвращаем nil, так как список пуст
}
