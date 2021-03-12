package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type node struct {
	name     string
	consumed bool
	next     []*node
	previous []*node
}

var start *node
var nodes map[string]*node

func resetNodes() {
	dat, err := ioutil.ReadFile("input/day07.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	ptStr := strings.Split(string(dat), "\n")
	nodes = make(map[string]*node)

	possbileStarts := make(map[string]bool)

	for _, pt := range ptStr {
		if pt != "" {
			var from, to string
			fmt.Sscanf(pt, "Step %s must be finished before step %s can begin.", &from, &to)
			nodeFrom, ok := nodes[from]
			if !ok {
				nodeFrom = &node{name: from, consumed: false}
				nodes[from] = nodeFrom

				if _, ok := possbileStarts[from]; !ok {
					possbileStarts[from] = true
				}
			}
			nodeTo, ok := nodes[to]
			if !ok {
				nodeTo = &node{name: to, consumed: false}
				nodes[to] = nodeTo
			}
			nodeFrom.next = append(nodeFrom.next, nodeTo)
			sort.Slice(nodeFrom.next, func(i, j int) bool {
				return nodeFrom.next[i].name < nodeFrom.next[j].name
			})

			nodeTo.previous = append(nodeTo.previous, nodeFrom)

			possbileStarts[to] = false
		}
	}

	start = &node{name: "", consumed: true}
	for s, possible := range possbileStarts {
		if possible {
			start.next = append(start.next, nodes[s])
		}
	}

	sort.Slice(start.next, func(i, j int) bool {
		return start.next[i].name < start.next[j].name
	})
}

func consumableNodes(node *node) []string {
	consumables := make([]string, 0)
	if !node.consumed {
		allPreviousNodesComsumed := true
		for _, p := range node.previous {
			if !p.consumed {
				allPreviousNodesComsumed = false
				break
			}
		}
		if allPreviousNodesComsumed {
			consumables = append(consumables, node.name)
		}
		return consumables
	}

	for _, next := range node.next {
		nc := consumableNodes(next)
		for _, newName := range nc {
			contains := false
			for _, currentName := range consumables {
				if currentName == newName {
					contains = true
				}
			}

			if !contains {
				consumables = append(consumables, newName)
			}
		}
	}

	sort.Slice(consumables, func(i, j int) bool {
		return consumables[i] < consumables[j]
	})

	return consumables
}

func Day07() (string, int) {
	fmt.Println("Day 07")

	var part1Order strings.Builder
	var part2Order strings.Builder

	resetNodes()

	for cn := consumableNodes(start); len(cn) != 0; cn = consumableNodes(start) {
		part1Order.WriteString(cn[0])
		nodes[cn[0]].consumed = true
	}

	resetNodes()
	var workers [5]struct {
		node    *node
		endTime int
	}

	totalTime := 0
	for totalTime = 0; ; totalTime++ {
		j := 0
		finishedWorker := 0

		// Finish workers
		for wrkIdx := range workers {
			worker := &workers[wrkIdx]

			if worker.node != nil && worker.endTime == totalTime {
				part2Order.WriteString(worker.node.name)
				worker.node.consumed = true
				worker.node = nil
				worker.endTime = 0
			}
		}

		// Get consumable nodes
		cn := consumableNodes(start)

		// Remove nodes that are currently processed
		var pn []string
		for _, n := range cn {
			isProcessed := false
			for wrkIdx := range workers {
				worker := &workers[wrkIdx]
				if worker.node != nil && worker.node.name == string(n) {
					isProcessed = true
					break
				}
			}

			if !isProcessed {
				pn = append(pn, n)
			}
		}

		// Start new workers
		for wrkIdx := range workers {
			worker := &workers[wrkIdx]

			if worker.endTime <= totalTime && j < len(pn) {
				worker.node = nodes[pn[j]]
				worker.endTime = totalTime + 60 + int(pn[j][0]) - 65 + 1
				j++
			} else if worker.endTime <= totalTime {
				finishedWorker++
			}
		}

		if 0 == len(cn) {
			break
		}
	}

	fmt.Println(workers)

	for _, worker := range workers {
		if worker.endTime > totalTime {
			totalTime = worker.endTime
		}
	}

	fmt.Println("Part 1:", part1Order.String())
	fmt.Println("Part 2:", totalTime, part2Order.String())

	return part1Order.String(), totalTime
}
