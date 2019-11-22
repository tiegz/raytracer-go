package raytracer

import (
	"io/ioutil"
	"testing"
)

func TestIgnoringUnrecognizedLines(t *testing.T) {
	gibberish :=
		`There was a young lady named Bright
who traveled much faster than light.
She set out one day
in a relative way,
and came back the previous night.`
	parser := ParseObjFile(gibberish)
	assertEqualInt(t, 5, parser.IgnoredLineCount)
}

func TestVertexRecords(t *testing.T) {
	gibberish :=
		`v -1 1 0
v -1.0000 0.5000 0.0000
v 1 0 0
v 1 1 0`
	parser := ParseObjFile(gibberish)
	assertEqualTuple(t, NewPoint(-1, 1, 0), parser.Vertices[1])
	assertEqualTuple(t, NewPoint(-1, 0.5, 0), parser.Vertices[2])
	assertEqualTuple(t, NewPoint(1, 0, 0), parser.Vertices[3])
	assertEqualTuple(t, NewPoint(1, 1, 0), parser.Vertices[4])
}

func TestParsingTriangleFaces(t *testing.T) {
	file :=
		`v -1 1 0
v -1.0000 0 0
v 1 0 0
v 1 1 0

f 1 2 3
f 1 3 4`
	parser := ParseObjFile(file)
	group := parser.DefaultGroup
	g := group.LocalShape.(Group)

	shape1 := g.Children[0]
	t1 := shape1.LocalShape.(*Triangle)
	shape2 := g.Children[1]
	t2 := shape2.LocalShape.(*Triangle)

	assertEqualTuple(t, parser.Vertices[1], t1.P1)
	assertEqualTuple(t, parser.Vertices[2], t1.P2)
	assertEqualTuple(t, parser.Vertices[3], t1.P3)
	assertEqualTuple(t, parser.Vertices[1], t2.P1)
	assertEqualTuple(t, parser.Vertices[3], t2.P2)
	assertEqualTuple(t, parser.Vertices[4], t2.P3)
}

func TestTriangulatingPolygons(t *testing.T) {
	file := `v -1 1 0
v -1.0000 0 0
v 1 0 0
v 1 1 0
v 0 2 0

f 1 2 3 4 5`
	parser := ParseObjFile(file)
	group := parser.DefaultGroup
	g := group.LocalShape.(Group)

	shape1 := g.Children[0]
	t1 := shape1.LocalShape.(*Triangle)
	shape2 := g.Children[1]
	t2 := shape2.LocalShape.(*Triangle)
	shape3 := g.Children[2]
	t3 := shape3.LocalShape.(*Triangle)

	assertEqualTuple(t, parser.Vertices[1], t1.P1)
	assertEqualTuple(t, parser.Vertices[2], t1.P2)
	assertEqualTuple(t, parser.Vertices[3], t1.P3)
	assertEqualTuple(t, parser.Vertices[1], t2.P1)
	assertEqualTuple(t, parser.Vertices[3], t2.P2)
	assertEqualTuple(t, parser.Vertices[4], t2.P3)
	assertEqualTuple(t, parser.Vertices[1], t3.P1)
	assertEqualTuple(t, parser.Vertices[4], t3.P2)
	assertEqualTuple(t, parser.Vertices[5], t3.P3)
}

func TestTrianglesInGroups(t *testing.T) {
	dat, err := ioutil.ReadFile("files/triangles.obj")
	if err != nil {
		panic(err)
	}
	parser := ParseObjFile(string(dat))

	group1 := parser.Groups["FirstGroup"]
	g1 := group1.LocalShape.(Group)

	group2 := parser.Groups["SecondGroup"]
	g2 := group2.LocalShape.(Group)

	shape1 := g1.Children[0]
	t1 := shape1.LocalShape.(*Triangle)
	shape2 := g2.Children[0]
	t2 := shape2.LocalShape.(*Triangle)

	assertEqualTuple(t, parser.Vertices[1], t1.P1)
	assertEqualTuple(t, parser.Vertices[2], t1.P2)
	assertEqualTuple(t, parser.Vertices[3], t1.P3)
	assertEqualTuple(t, parser.Vertices[1], t2.P1)
	assertEqualTuple(t, parser.Vertices[3], t2.P2)
	assertEqualTuple(t, parser.Vertices[4], t2.P3)
}

// Scenario: Converting an OBJ file to a group
// Given file ← the file "triangles.obj"
// And parser ← parse_obj_file(file)
// When g ← obj_to_group(parser)
// Then g includes "FirstGroup" from parser
// And g includes "SecondGroup" from parser

func TestConvertingAnOBJFileToAGroup(t *testing.T) {
	dat, err := ioutil.ReadFile("files/triangles.obj")
	if err != nil {
		panic(err)
	}
	parser := ParseObjFile(string(dat))
	group := parser.ToGroup()
	g := group.LocalShape.(Group)

	// TODO: abstract into an assertContainsValue([]interface{}) method,
	// and maybe turn Group into a Contains interface?
	expectedGroup := parser.Groups["FirstGroup"]
	if !g.Contains(expectedGroup) {
		t.Errorf("\nExpected group to contain %s, but did not.\n", "FirstGroup")
	}

	expectedGroup = parser.Groups["SecondGroup"]
	if !g.Contains(expectedGroup) {
		t.Errorf("\nExpected group to contain %s, but did not.\n", "SecondGroup")
	}
}
