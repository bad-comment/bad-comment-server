package main

import (
	"fmt"
	tws "github.com/muyu66/two-way-score"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"slices"
)

type Comment struct {
	RaterId int64
	UserId  int64
	Score   int8
}

func calcScore(users *[]User, comments *[]Comment) map[int64]float64 {
	dg := simple.NewDirectedGraph()
	for _, user := range *users {
		node := graph.Node(simple.Node(user.Id))
		dg.AddNode(node)
	}

	for _, comment := range *comments {
		edge := ScoreEdge{F: simple.Node(comment.RaterId), T: simple.Node(comment.UserId), Score: comment.Score}
		dg.SetEdge(edge)
	}

	var d = toFullGraph(dg, 2)
	//fmt.Printf("%+v\n", d)

	var ddd []tws.Node
	for _, dd := range d {
		ddd = append(ddd, tws.Node{
			RaterId:  dd.FromId,
			TargetId: dd.ToId,
			Deep:     dd.Deep,
			Score:    int64(dd.Score),
		})
	}
	//fmt.Printf("%+v\n", ddd)
	ss, _ := tws.Calc(&ddd)
	fmt.Printf("%+v\n", ss)

	var a = make(map[int64]float64)
	for k, v := range ss {
		a[k.(int64)] = v
	}
	return a
}

func toFullGraph(
	dg *simple.DirectedGraph,
	id int64,
) []Asd {
	// 获取节点的所有邻居
	neighbors := dg.To(id)
	neighbors2 := dg.From(id)

	var asdd = make([]Asd, 0)

	iterator(false, neighbors2, dg, 0, id, &asdd)

	var deep2 int64 = 0
	if len(asdd) > 0 {
		deep2 = slices.MaxFunc(asdd, func(a, b Asd) int {
			if a.Deep > b.Deep {
				return 1
			} else if a.Deep < b.Deep {
				return -1
			}
			return 0
		}).Deep
	}

	iterator(true, neighbors, dg, deep2, id, &asdd)

	var deep3 int64 = 0
	if len(asdd) > 0 {
		deep3 = slices.MinFunc(asdd, func(a, b Asd) int {
			if a.Deep > b.Deep {
				return 1
			} else if a.Deep < b.Deep {
				return -1
			}
			return 0
		}).Deep
	}

	// deep补正
	for i, _ := range asdd {
		asdd[i].Deep += -deep3 + 1
	}

	return asdd
}

func iterator(
	to bool,
	neighbors graph.Nodes,
	dg *simple.DirectedGraph,
	deep int64,
	fromId int64,
	asdd *[]Asd,
) {
	if to {
		deep++
	} else {
		deep--
	}
	for neighbors.Next() {
		currNode := neighbors.Node()
		var nodes graph.Nodes
		if to {
			e := dg.Edge(currNode.ID(), fromId).(ScoreEdge)
			*asdd = append(*asdd, Asd{
				FromId: currNode.ID(),
				ToId:   fromId,
				Deep:   deep,
				Score:  e.Score,
			})
			nodes = dg.To(currNode.ID())
		} else {
			e := dg.Edge(fromId, currNode.ID()).(ScoreEdge)
			*asdd = append(*asdd, Asd{
				FromId: fromId,
				ToId:   currNode.ID(),
				Deep:   deep,
				Score:  e.Score,
			})
			nodes = dg.From(currNode.ID())
		}
		iterator(to, nodes, dg, deep, currNode.ID(), asdd)
	}
}

func (s ScoreEdge) From() graph.Node {
	return s.F
}

func (s ScoreEdge) To() graph.Node {
	return s.T
}

func (s ScoreEdge) ReversedEdge() graph.Edge {
	return nil
}

type ScoreEdge struct {
	F     graph.Node
	T     graph.Node
	Score int8
}

type Asd struct {
	FromId int64
	ToId   int64
	Deep   int64
	Score  int8
}
