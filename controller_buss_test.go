package main

import (
	"fmt"
	"github.com/glebarez/sqlite"
	_ "github.com/glebarez/sqlite"
	"github.com/stretchr/testify/suite"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type BussSuite struct {
	suite.Suite
	db *gorm.DB
}

func (s *BussSuite) SetupTest() {
	s.db, _ = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	_ = s.db.AutoMigrate(&User{}, &Subject{}, &SubjectComment{}, &UserScore{})
}

func TestBussSuite(t *testing.T) {
	suite.Run(t, new(BussSuite))
}

func testPost(e *echo.Echo, body string) echo.Context {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec)
}

func auth(c echo.Context, userId int64) {
	c.Set("userId", userId)
}

func (s *BussSuite) Test() {
	e := echo.New()
	ctl := &controller{db: s.db}

	// 生成用户
	for i := 1; i <= 20; i++ {
		body := fmt.Sprintf(`{"account":"biubiu%s","password":"123456"}`, strconv.Itoa(i))
		c := testPost(e, body)

		err := ctl.signUp(c)
		if err != nil {
			assert.FailNow(s.T(), err.Error())
		}
	}

	// 用户发表subject
	for i := 1; i <= 10; i++ {
		body := fmt.Sprintf(`{"name":"iphone%s"}`, strconv.Itoa(i))
		c := testPost(e, body)

		auth(c, rand.Int63n(10))
		err := ctl.createSubject(c)
		if err != nil {
			assert.FailNow(s.T(), err.Error())
		}
	}

	// 用户吐槽
	for i := 1; i <= 100; i++ {
		body := fmt.Sprintf(`{"score":%d}`, rand.Intn(9)+1)
		c := testPost(e, body)
		c.SetPath("/subjects/:subjectId/comments")
		c.SetParamNames("subjectId")
		c.SetParamValues(strconv.Itoa(rand.Intn(10)))

		auth(c, rand.Int63n(100))
		err := ctl.createSubjectComment(c)
		if err != nil {
			assert.FailNow(s.T(), err.Error())
		}
	}

	var comments []SubjectComment
	s.db.Table("subject_comment").Find(&comments)
	fmt.Printf("%+v\n", comments)

	var subjects []Subject
	s.db.Table("subject").Find(&subjects)
	fmt.Printf("%+v\n", subjects)

	dg := simple.NewDirectedGraph()
	for _, subject := range subjects {
		node := graph.Node(simple.Node(subject.Id))
		dg.AddNode(node)
	}

	for _, comment := range comments {
		edge := simple.Edge{F: simple.Node(comment.UserId), T: simple.Node(comment.SubjectId)}
		dg.SetEdge(edge)
	}

	var d = toFullGraph(dg, 1)
	fmt.Printf("%+v\n", d)
}

type Asd struct {
	FromId int64
	ToId   int64
	Deep   int64
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
			*asdd = append(*asdd, Asd{
				FromId: currNode.ID(),
				ToId:   fromId,
				Deep:   deep,
			})
			nodes = dg.To(currNode.ID())
		} else {
			*asdd = append(*asdd, Asd{
				FromId: fromId,
				ToId:   currNode.ID(),
				Deep:   deep,
			})
			nodes = dg.From(currNode.ID())
		}
		iterator(to, nodes, dg, deep, currNode.ID(), asdd)
	}
}

func toFullGraph(
	dg *simple.DirectedGraph,
	id int64,
) []Asd {
	// 获取节点的所有邻居
	neighbors := dg.To(id)
	neighbors2 := dg.From(id)

	var asdd []Asd

	iterator(false, neighbors2, dg, 0, id, &asdd)

	deep2 := slices.MaxFunc(asdd, func(a, b Asd) int {
		if a.Deep > b.Deep {
			return 1
		} else if a.Deep < b.Deep {
			return -1
		}
		return 0
	}).Deep

	iterator(true, neighbors, dg, deep2, id, &asdd)

	deep3 := slices.MinFunc(asdd, func(a, b Asd) int {
		if a.Deep > b.Deep {
			return 1
		} else if a.Deep < b.Deep {
			return -1
		}
		return 0
	}).Deep

	// deep补正
	for i, _ := range asdd {
		asdd[i].Deep += -deep3 + 1
	}

	return asdd
}

func (s *BussSuite) TestBaseGraph() {
	dg := simple.NewDirectedGraph()
	// 假设我们有一些唯一标识符作为节点
	node1 := graph.Node(simple.Node(1))
	node2 := graph.Node(simple.Node(2))
	node3 := graph.Node(simple.Node(3))
	node4 := graph.Node(simple.Node(4))
	node5 := graph.Node(simple.Node(5))
	node6 := graph.Node(simple.Node(6))
	node7 := graph.Node(simple.Node(7))
	node8 := graph.Node(simple.Node(8))

	// 将节点添加到图中
	dg.AddNode(node1)
	dg.AddNode(node2)
	dg.AddNode(node3)
	dg.AddNode(node4)
	dg.AddNode(node5)
	dg.AddNode(node6)
	dg.AddNode(node7)
	dg.AddNode(node8)

	// 创建一条边，并指定两个端点
	edge1 := simple.Edge{F: node4, T: node2}
	edge2 := simple.Edge{F: node3, T: node2}
	edge3 := simple.Edge{F: node2, T: node1}
	edge4 := simple.Edge{F: node5, T: node1}
	edge5 := simple.Edge{F: node6, T: node5}
	edge6 := simple.Edge{F: node1, T: node7}
	edge7 := simple.Edge{F: node7, T: node8}

	// 将边添加到图中
	dg.SetEdge(edge1)
	dg.SetEdge(edge2)
	dg.SetEdge(edge3)
	dg.SetEdge(edge4)
	dg.SetEdge(edge5)
	dg.SetEdge(edge6)
	dg.SetEdge(edge7)

	var asdd = toFullGraph(dg, node1.ID())

	var expected = "[{FromId:1 ToId:7 Deep:1} {FromId:7 ToId:8 Deep:2} {FromId:1 ToId:2 Deep:3} {FromId:2 ToId:4 Deep:4} {FromId:2 ToId:3 Deep:4} {FromId:1 ToId:5 Deep:3} {FromId:5 ToId:6 Deep:4}]"
	assert.Equal(s.T(), expected, fmt.Sprintf("%+v", asdd))
}
