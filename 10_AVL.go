package avl

type node struct {
    key   int
    h     int8
    left  *node
    right *node
}

func Node(key int) *node {
    var output node
    output.h = 1
    output.right = nil
    output.left = nil
    output.key = key
    return &output
}

func (n *node) height() int8 {
    if n == nil {
        return 0
    } else {
        return n.h
    }
}

func (n *node) bFactor() int8 {
    return n.right.height() - n.left.height()
}

func (n *node) fixHeight() {
    heightLeft := n.left.height()
    heightRight := n.right.height()
    if heightLeft > heightRight {
        n.h = heightLeft + 1
    } else {
        n.h = heightRight + 1
    }
}

func rotateLeft(n *node) *node {
    p := n.right
    n.right = p.left
    p.left = n
    n.fixHeight()
    p.fixHeight()
    return p
}

func rotateRight(n *node) *node {
    p := n.left
    n.left = p.right
    p.right = n
    n.fixHeight()
    p.fixHeight()
    return p
}

func balance(n *node) *node {
    n.fixHeight()
    if n.bFactor() == 2 {
        if n.right.bFactor() < 0 {
            n.right = rotateRight(n.right)
        }
        return rotateLeft(n)
    } else if n.bFactor() == -2 {
        if n.left.bFactor() > 0 {
            n.left = rotateLeft(n.left)
        }
        return rotateRight(n)
    }
    return n
}

func Insert(n *node, key int) *node {
    if n == nil {
        return Node(key)
    } else if key < n.key {
        n.left = Insert(n.left, key)
    } else if key > n.key {
        n.right = Insert(n.right, key)
    } else {
        return n
    }
    return balance(n)
}

func (n *node) findMin() *node {
    if n.left == nil {
        return n
    } else {
        return n.left.findMin()
    }
}

func (n *node) removeMin() *node {
    if n.left == nil {
        return n.right
    }
    n.left = n.left.removeMin()
    return balance(n)
}

func Delete(n *node, key int) *node {
    if n == nil {
        return nil
    }
    if key < n.key {
        n.left = Delete(n.left, key)
    } else if key > n.key {
        n.right = Delete(n.right, key)
    } else {
        q := n.left
        r := n.right
        if r == nil {
            return q
        }
        min := r.findMin()
        min.right = r.removeMin()
        min.left = q
        return balance(min)
    }
    return balance(n)
}

func main() {
    root := avl.Node(1)
    root = avl.Insert(root, 2)
    root = avl.Insert(root, 3)
    root = avl.Insert(root, 4)
    root = avl.Insert(root, 5)
    root = avl.Insert(root, 6)
    root = avl.Insert(root, 7)
    fmt.Println(root)
    root = avl.Delete(root, 4)
    root = avl.Delete(root, 5)
    fmt.Println(root)
}
