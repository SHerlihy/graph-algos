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


func BreadthTraversalCoords(graph map[[2]int][][2]int, initNode [2]int, nodeCondition func([2]int)bool){
    nodeQueue := [][2]int{initNode}

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

        fmt.Printf("\nPLEASEWORK: %v", len(curEdges))
        pass := nodeCondition(curNode)

        if len(curEdges) == 0 || !pass {
            continue
        }

        nodeQueue = append(nodeQueue, curEdges...)
    }
}
// couldnt care less about type of data in subject
func TwoDSliceToGraph (subject [][]string) map[[2]int][][2]int {
    graph := map[[2]int][][2]int{}

    yMax := len(subject)-1
    xMax := len(subject[0])-1

    if xMax < 2 || yMax < 2 {
        return graph
    }

    for i:= 0; i<=yMax; i++ {
        fmt.Printf("\nxMax %v, %v", i, i)
        verts := [2]int{i-1,i+1}

        for j:=0; j<=xMax; j++{
            horizi := [2]int{j-1,j+1}

            neighbours := [4][2]int{}

            for vi, vert := range verts {
                for hi, horiz := range horizi {
                    neighbours[(vi*2)+hi] = [2]int{horiz,vert}
                }
            }

                nKey := [2]int{j,i}
                graph[nKey] = neighbours[:]
        }
    }

    return graph
}

//func getCornerNeighbours(xMax int, yMax int) map[[2]int][][2]int {
//    cornerNeighbours := map[[2]int][2]int{}
//
//    tl := [][2]int{[2]int{0,1}, [2]int{1,1}, [2]int{1,0}}
//    cornerNeighbours[[2]int{0,0}] = tl
//
//    tr := [][2]int{[2]int{xMax,1}, [2]int{xMax-1,1}, [2]int{xMax-1,0}}
//    cornerNeighbours[[2]int{xMax,0}] = tr
//
//    br := [][2]int{[2]int{xMax-1,yMax}, [2]int{xMax-1,yMax-1}, [2]int{xMax,yMax-1}}
//    cornerNeighbours[[2]int{xMax,yMax}] = br
//
//    bl := [][2]int{[2]int{0,yMax-1}, [2]int{1,yMax-1}, [2]int{1, yMax}}
//    cornerNeighbours[[2]int{0,yMax}]
//
//    return cornerNeighbours
//}
