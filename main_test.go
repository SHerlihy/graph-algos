package main

import (
    "testing"
    "strconv"
)

func TestTwoDSliceToGraph(t *testing.T){
    grid := [][]string{
        []string{"w","i","i"},
        []string{"w","i","i"},
        []string{"w","w","w"},
    }

    retGraph := TwoDSliceToGraph(grid)

    foundVals := [][2]int{}

    matcher := "i"

    initNode := [2]int{1,1}

    findLand := findCoordinatesCurry(matcher, &foundVals, grid)

    BreadthTraversalCoords(retGraph, initNode, findLand)

    t.Logf("\nFound vals: %v", foundVals)
}

func findCoordinatesCurry(matcher string, foundVals *[][2]int, coordVals[][]string) func([2]int)bool {
    preventLoop := preventGraphLoop()
    return func (node [2]int) bool {
        return findCoordinates(matcher, foundVals, node, coordVals, preventLoop)
    }
}

func findCoordinates(matcher string, foundVals *[][2]int, node [2]int, coordVals[][]string, preventLoop func(int)bool)bool{
    nodeString := ""
    nodeString = strconv.Itoa(node[0])+strconv.Itoa(node[1])

    if node[1] < 0 {
    nodeString = strconv.Itoa(node[1])+strconv.Itoa(node[0])
    }

    if node[0] < 0 && node[1] < 0 {
    nodeString = strconv.Itoa(node[0])+strconv.Itoa(-node[1])
    }

    intNode, err := strconv.Atoi(nodeString)

    if err != nil {
        panic(err)
    }

    noLoop := preventLoop(intNode)

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
    preventLoop := preventGraphLoop()

    DepthTraversal(graph, initNode, opOnNodes)
    DepthTraversal(graph, initNode, preventLoop)
    BreadthTraversal(graph, initNode, opOnNodes)

    t.Logf("\n%v", nodesAsTraversed)
}

func preventGraphLoop () func(node int) bool {
    nodeSet := map[int]bool{}

    return func(node int) bool {
        return preventLoop(nodeSet, node)
    }
}

func preventLoop (nodeSet map[int]bool, node int) bool {
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
