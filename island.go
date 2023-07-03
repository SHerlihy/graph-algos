package main

import (
    "fmt"
)

func BreadthTraversalCoords(graph *map[[2]int][][2]int, initNode [2]int, nodeCondition func([2]int)bool, addNeighboursToGraph func([2]int)){
    nodeQueue := [][2]int{initNode}

    for {
        queueLength := len(nodeQueue)

        if queueLength == 0{
            break
        }

        curNode := nodeQueue[0]
        nodeQueue = nodeQueue[1:]

        addNeighboursToGraph(curNode)

        graphVal := *graph

        curEdges, ok := graphVal[curNode]

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

func AddNeighboursToGraphCurry(graph *map[[2]int][][2]int, grid [][]string) func ([2]int) {
    return func (node [2]int) {
        addNeighboursToGraph(graph, grid, node)
    }
}

func addNeighboursToGraph (graph *map[[2]int][][2]int, grid [][]string, node [2]int) {
    graphVal := *graph
    _, ok := graphVal[node]

    if !ok {
        fmt.Errorf("\n AddNeighboursToGraph: Node %v not in graph", node)
    }

    xNode := node[0]
    yNode := node[1]

    tl := [2]int{xNode-1,yNode-1}
    tr := [2]int{xNode+1,yNode-1}
    br := [2]int{xNode+1, yNode+1}
    bl := [2]int{xNode-1, yNode+1}

    graphVal[node] = [][2]int{tl,tr,br,bl}
}

func FindCoordinatesCurry(matcher string, foundVals *[][2]int, coordVals[][]string, preventLoop func([2]int) bool) func([2]int)bool {
    return func (node [2]int) bool {
        return findCoordinates(matcher, foundVals, node, coordVals, preventLoop)
    }
}

func findCoordinates(matcher string, foundVals *[][2]int, node [2]int, coordVals[][]string, preventLoop func([2]int)bool)bool{
    noLoop := preventLoop(node)

    if !noLoop {
        return false
    }

    if node[0] < 0 || node[1] < 0 {
        return false
    }

    if node[0] > 2 || node[1] > 2 {
        return false
    }


    fVal := coordVals[node[0]][node[1]]

    if fVal == matcher{
        *foundVals = append(*foundVals, node)
    }

    return true
}

func PreventGraphLoopCoord() (*map[[2]int]bool, func(node [2]int) bool) {
    nodeSet := map[[2]int]bool{}

    return &nodeSet, func(node [2]int) bool {
        return preventLoopCoord(&nodeSet, node)
    }
}

func preventLoopCoord (nodeSet *map[[2]int]bool, node [2]int) bool {
    nSetVals := *nodeSet
    _, ok := nSetVals[node]

    if ok {
        return false
    }

    nSetVals[node] = true

    return true
}
