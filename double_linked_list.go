package main

import (
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
func (list *List[T]) Insert(info T, index int) any {

    if index > list.size {
        err := any("List overflow")
        return err
    }

    var newNode node[T]
    newNode.info = info

    if list.size == 0 {
        list.head = &newNode
        list.tail = &newNode
        list.size = 1
        return nil
    }

    //if !list.checkCompatibility(info) {
    //   err := any("Item must be the same type as other elements in list")
    //    return err
    //}

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
func (list List[T]) Access(index int) (any, T) {

    if index >= list.size {
        err := any("List overflow")
        var result T
        return err, result
    }
    if index < list.size/2 {
        curNode := list.head
        for i := 0; i < index; i++ {
            curNode = curNode.next
        }
        return nil, curNode.info
    } else {
        curNode := list.tail
        index = list.size - index - 1
        for i := 0; i < index; i++ {
            curNode = curNode.prev
        }
        return nil, curNode.info
    }
}

func (list List[T]) Size() int {
    return list.size
}

//Erase returns err
func (list *List[T]) Erase(start, end int) any { //Removes range of elements ([first,last)).
    if start >= list.size || end >= list.size {
        err := any("List overflow")
        return err
    }
    var startNode *node[T]
    if start == 0 {
        if end == 0 {
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
    //start := time.Now()
    for i := 0; i < 10; i++ {
        arr.Insert(i, i)
    }
    arr.Print()
    arr.Erase(7, 9)
    fmt.Println("")
    arr.Print()
    /*
       for i := 0; i < 10; i++ {
           _, a := arr.Access(i)
           fmt.Print(a)
           fmt.Print(" ")
       }*/
    //fmt.Println(time.Since(start))
}
