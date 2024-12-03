package main

// reverse переворачивает список
func (ll *LinkedList) Reverse() {
	var prev *Node
	current := ll.head
	for current != nil {
		next := current.next // сохраняем следующий узел
		current.next = prev  // изменяем указатель текущего узла на предыдущий
		prev = current       // перемещаем предыдущий узел на текущий
		current = next       // переходим к следующему узлу
	}
	ll.head = prev // изменяем головной узел на последний узел
}
