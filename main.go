package main

import (
    "fmt"
)

func DepthTraversal(graph map[int][]int, initNode int, nodeCondition func(int)bool) {
    nodeStack := []int{initNode}

    for {
        stackLen := len(nodeStack)

        if stackLen == 0{
            break
        }

        curNode := nodeStack[stackLen-1]
        nodeStack = nodeStack[:stackLen - 1]

        curEdges, ok := graph[curNode]

        if !ok {
            fmt.Errorf("Node %v not present in graph", curNode)
        }

        pass := nodeCondition(curNode)

        if len(curEdges) == 0 || !pass {
            continue
        }

        nodeStack = append(nodeStack, curEdges...)
    }
}

func BreadthTraversal(graph map[int][]int, initNode int, nodeCondition func(int)bool){
    nodeQueue := []int{initNode}

    for {
        queueLength := len(nodeQueue)

        if queueLength == 0{
            break
        }

        curNode := nodeQueue[0]
        nodeQueue = nodeQueue[1:]

        curEdges, ok := graph[curNode]

        if !ok {
            fmt.Errorf("Node %v not present in graph", curNode)
        }

        pass := nodeCondition(curNode)

        if len(curEdges) == 0 || !pass {
            continue
        }

        nodeQueue = append(nodeQueue, curEdges...)
    }
}

func PreventGraphLoopInt () (*map[int]bool, func(node int) bool) {
    nodeSet := map[int]bool{}

    return &nodeSet, func(node int) bool {
        return preventLoop(&nodeSet, node)
    }
}

func preventLoop (nodeSet *map[int]bool, node int) bool {
    nSetVals := *nodeSet
    _, ok := nSetVals[node]

    if ok {
        return false
    }

    nSetVals[node] = true

    return true
}

func AppendNodesCurry (nodesAsTraversed *[]int) func(int) bool {
    return func (node int) bool {
        return appendNode(node, nodesAsTraversed)
    }
}

func appendNode (node int, nodesAsTraversed *[]int) bool {
    updatedNodes := append(*nodesAsTraversed, node)
    *nodesAsTraversed = updatedNodes
    return true
}
