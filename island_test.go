package main

import (
    "testing"
)

func TestTwoDSliceToGraph(t *testing.T){
    grid := [][]string{
        []string{"w","i","i"},
        []string{"w","i","i"},
        []string{"w","w","w"},
    }

    graph := map[[2]int][][2]int{}

    foundVals := [][2]int{}

    visitedNodes, preventLoop := PreventGraphLoopCoord()

    matcher := "i"

    addNeighboursToGraph := AddNeighboursToGraphCurry(&graph, grid)

    findLand := FindCoordinatesCurry(matcher, &foundVals, grid, preventLoop)

    for ri, row := range grid {
        for ci, _ := range row {
            node := [2]int{ri,ci}
            visitedNodesVals := *visitedNodes
            _, ok := visitedNodesVals[node]

            if ok {
                continue
            }

            if grid[ri][ci] == "i"{
                BreadthTraversalCoords(&graph, node, findLand, addNeighboursToGraph)
            }
        }
    }

    t.Logf("\nFound vals: %v", foundVals)
}

