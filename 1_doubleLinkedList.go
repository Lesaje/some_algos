package main

import (
    "errors"
    "fmt"
)

type node[T any] struct {
    info T
    next *node[T]
    prev *node[T]
}

type List[T any] struct {
    head *node[T]
    tail *node[T]
    size int
}

//Insert returns err if something wrong, nil otherwise
func (list *List[T]) Insert(info T, index int) error {

    if index > list.size {
        return errors.New("list overflow")
    }

    var newNode node[T]
    newNode.info = info

    if list.size == 0 {
        list.head = &newNode
        list.tail = &newNode
        list.size = 1
        return nil
    }

    if list.size == index {
        list.tail.next = &newNode
        newNode.prev = list.tail
        list.tail = &newNode
        list.size += 1
        return nil
    }
    if index == 0 {
        newNode.next = list.head
        list.head.prev = &newNode
        list.head = &newNode
        list.size += 1
        return nil
    }

    if index < list.size/2 {
        curNode := list.head
        for i := 0; i < index; i++ {
            curNode = curNode.next
        }
        newNode.prev = curNode.prev
        newNode.next = curNode
        curNode.prev.next = &newNode
        curNode.prev = &newNode
        list.size += 1
        return nil
    } else {
        curNode := list.tail
        index = list.size - index
        for i := 0; i < index; i++ {
            curNode = curNode.prev
        }
        newNode.prev = curNode
        newNode.next = curNode.next
        curNode.next.prev = &newNode
        curNode.next = &newNode
        list.size += 1
        return nil
    }
}

func (list List[T]) Print() { // because any cannot be converted to string, there is
    curNode := list.head // no String func for List
    for i := 0; i < list.size; i++ {
        fmt.Print(curNode.info)
        fmt.Print(" ")
        curNode = curNode.next
    }
}

//Access returns err, element
func (list List[T]) Access(index int) (T, error) {

    if index >= list.size {
        var result T
        return result, errors.New("list overflow")
    }
    if index < list.size/2 {
        curNode := list.head
        for i := 0; i < index; i++ {
            curNode = curNode.next
        }
        return curNode.info, nil
    } else {
        curNode := list.tail
        index = list.size - index - 1
        for i := 0; i < index; i++ {
            curNode = curNode.prev
        }
        return curNode.info, nil
    }
}

func (list List[T]) Size() int {
    return list.size
}

//Erase returns err
//Removes range of elements ([first,last)).
func (list *List[T]) Erase(start, end int) error {
    if start >= list.size || end >= list.size {
        return errors.New("list overflow")
    }
    var startNode *node[T]
    if start == 0 {
        if end == 0 {
            if list.size == 1 {
                list.size = 0
                return nil
            }
            list.head = list.head.next
            list.head.prev = nil
            list.size -= 1
            return nil
        }
        if end == list.size-1 {
            list.head = nil
            list.tail = nil
            list.size = 0
            return nil
        }

        if end < list.size/2 {
            startNode = list.head
            for i := 0; i < end; i++ {
                startNode = startNode.next
            }
        } else {
            startNode = list.tail
            startl := list.size - end - 1
            for i := 0; i < startl; i++ {
                startNode = startNode.prev
            }
        }
        list.head = startNode.next
        list.head.prev = nil
        list.size = list.size - end - 1
        return nil
    }

    if end == list.size-1 {
        if start == list.size-1 {
            list.tail = list.tail.prev
            list.tail.next = nil
            list.size -= 1
            return nil
        }
        if start == 0 {
            list.head = nil
            list.tail = nil
            list.size = 0
            return nil
        }

        if start < list.size/2 {
            startNode = list.head
            for i := 0; i < start; i++ {
                startNode = startNode.next
            }
        } else {
            startNode = list.tail
            startl := list.size - start - 1
            for i := 0; i < startl; i++ {
                startNode = startNode.prev
            }
        }
        list.tail = startNode.prev
        list.tail.next = nil
        list.size = list.size - (end - start + 1)
        return nil
    }

    if start < list.size/2 {
        startNode = list.head
        for i := 0; i < start; i++ {
            startNode = startNode.next
        }
    } else {
        startNode = list.tail
        startl := list.size - start - 1
        for i := 0; i < startl; i++ {
            startNode = startNode.prev
        }
    }

    var endNode *node[T]
    if end < list.size/2 {
        endNode = list.head
        for i := 0; i < end; i++ {
            endNode = endNode.next
        }
    } else {
        endNode = list.tail
        endl := list.size - end - 1
        for i := 0; i < endl; i++ {
            endNode = endNode.prev
        }
    }

    startNode.prev.next = endNode.next
    endNode.next.prev = startNode.prev
    list.size = list.size - (end - start + 1)
    return nil
}

func main() {
    var arr List[int]
    for i := 0; i < 10; i++ {
        arr.Insert(i, i)
    }
    arr.Print()
    arr.Erase(7, 9)
    fmt.Println("")
    arr.Print()
}
