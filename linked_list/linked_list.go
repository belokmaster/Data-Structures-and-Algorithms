package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Node представляет собой узел связанного списка
type Node struct {
	value int   // значение узла
	next  *Node // указатель на следующий узел
}

// LinkedList представляет собой односвязный список
type LinkedList struct {
	head *Node // указатель на первый узел списка
}

// pushFront добавляет элемент в начало списка
func (ll *LinkedList) PushFront(value int) {
	newNode := &Node{value: value, next: ll.head} // создаем новый узел
	ll.head = newNode                             // новый узел становится головным узлом
}

// popFront удаляет элемент из начала списка
func (ll *LinkedList) PopFront() {
	if ll.head != nil { // если список не пуст
		ll.head = ll.head.next // головной узел становится следующим
	}
}

// pushBack добавляет элемент в конец списка
func (ll *LinkedList) PushBack(value int) {
	newNode := &Node{value: value} // создаем новый узел
	if ll.head == nil {            // если список пуст
		ll.head = newNode // новый узел становится головным
	} else {
		current := ll.head        // начинаем с головного узла
		for current.next != nil { // ищем последний узел
			current = current.next
		}
		current.next = newNode // добавляем новый узел в конец
	}
}

// popBack удаляет элемент из конца списка
func (ll *LinkedList) PopBack() {
	if ll.head == nil { // если список пуст
		return
	}
	if ll.head.next == nil { // если в списке только один элемент
		ll.head = nil // список становится пустым
		return
	}

	current := ll.head             // начинаем с головного узла
	for current.next.next != nil { // ищем предпоследний узел
		current = current.next
	}
	current.next = nil // удаляем последний узел
}

// add добавляет элемент в список по указанному индексу
func (ll *LinkedList) Add(index int, value int) {
	newNode := &Node{value: value} // создаем новый узел
	if index == 0 {                // если элемент добавляется в начало списка
		newNode.next = ll.head
		ll.head = newNode
		return
	}

	current := ll.head
	for i := 0; i < (index-1) && (current != nil); i++ { // ищем узел перед указанным индексом
		current = current.next
	}

	if current == nil { // если индекс вне диапазона
		fmt.Println("Индекс вне диапазона")
		return
	}

	newNode.next = current.next // вставляем новый узел в нужное место
	current.next = newNode
}

// remove удаляет элемент из список по указанному индексу
func (ll *LinkedList) Remove(index int) {
	if ll.head == nil { // если список пуст
		fmt.Println("Список пуст")
		return
	}

	if index == 0 { // если удаляется головной узел
		ll.head = ll.head.next
		return
	}

	current := ll.head
	for i := 0; i < (index-1) && (current != nil); i++ { // ищем узел перед указанным индексом
		current = current.next
	}

	if current == nil || current.next == nil { // если индекс вне диапазона
		fmt.Println("Индекс вне диапазона")
		return
	}

	current.next = current.next.next // удаляем узел
}

// deleteAll удаляет все элементы из списка
func (ll *LinkedList) DeleteAll() {
	ll.head = nil // просто обнуляем головной узел
}

// size возвращает количество элементов в списке
func (ll *LinkedList) Size() int {
	size := 0
	current := ll.head
	for current != nil { // проходим по всему списку, увеличивая счетчик
		size++
		current = current.next
	}
	return size
}

// display выводит все элементы списка
func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil { // проходим по всему списку, выводя значения узлов
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil") // конец списка
}

// generateRandomList создает связанный список с случайным количеством элементов и случайными значениями
func GenerateRandomList(maxElements, maxValue int) *LinkedList {
	ll := &LinkedList{}

	// Инициализируем новый генератор случайных чисел
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Определяем случайное количество элементов
	numElements := r.Intn(maxElements) + 1

	// Заполняем список случайными значениями
	for i := 0; i < numElements; i++ {
		value := r.Intn(maxValue)
		ll.PushBack(value)
	}

	return ll
}
