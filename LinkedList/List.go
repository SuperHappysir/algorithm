package LinkedList

type (
    Node struct {
        value int
        next  *Node
    }
    LinkedList struct {
        head *Node
    }
)

func (l *LinkedList) InsertFirst(value int) {
    node := &Node{value: value}
    if l.head != nil {
        node.next = l.head
    }
    l.head = node
}

func (l *LinkedList) InsertLast(value int) {
    node := &Node{value: value}
    if l.head == nil {
        l.head = node
        return
    }
    
    currentNode := l.head
    for currentNode.next != nil {
        currentNode = currentNode.next
    }
    currentNode.next = node
}

func (l *LinkedList) RemoveByValue(value int) *Node {
    if l.head == nil {
        return nil
    }
    currentNode := l.head
    if l.head.value == value {
        l.head = l.head.next
        return currentNode
    }
    for currentNode.next != nil {
        if currentNode.next.value == value {
            tempNode := currentNode.next
            currentNode.next = currentNode.next.next
            return tempNode
        }
        currentNode = currentNode.next
    }
    return nil
}

func (l *LinkedList) GetNode(value int) *Node {
    if l.head == nil {
        return nil
    }
    
    currentNode := l.head
    for currentNode != nil {
        if currentNode.value == value {
            return currentNode
        }
        currentNode = currentNode.next
    }
    return nil
}
