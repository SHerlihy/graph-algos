package main

import (
    "testing"
)

func TestDepthTraversal(t *testing.T){
    graph := map[int][]int{
        1: []int{10,11,12,13, 2, 3},
        10: []int{},
        11: []int{},
        12: []int{},
        13: []int{},
        2: []int{4},
        3: []int{4},
        4: []int{5},
        5: []int{14},
        14: []int{15},
        15: []int{},
    }

    initNode := 1
    nodesAsTraversed := []int{}
    opOnNodes := appendNodesCurry(&nodesAsTraversed)

    DepthTraversal(graph, initNode, opOnNodes)
    DepthTraversal(graph, initNode, preventGraphLoop)
    BreadthTraversal(graph, initNode, opOnNodes)

    t.Logf("\n%v", nodesAsTraversed)
}

func preventGraphLoop (node int) bool {
    nodeSet := map[int]bool{}

    _, ok := nodeSet[node]

    if ok {
        return false
    }

    nodeSet[node] = true

    return true
}

func appendNodesCurry (nodesAsTraversed *[]int) func(int) bool {
    return func (node int) bool {
        return appendNode(node, nodesAsTraversed)
    }
}

func appendNode (node int, nodesAsTraversed *[]int) bool {
    updatedNodes := append(*nodesAsTraversed, node)
    *nodesAsTraversed = updatedNodes
    return true
}
