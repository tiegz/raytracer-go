package raytracer

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type ObjFile struct {
	IgnoredLineCount int
	Vertices         []Tuple
	Normals          []Tuple
	DefaultGroup     *Shape
	Groups           map[string]*Shape
	CurrentGroupName string
}

func (of *ObjFile) ToGroup() Shape {
	g := NewGroup()
	for _, v := range of.Groups {
		g.AddChildren(v)
	}
	return g
}

func (o ObjFile) String() string {
	return fmt.Sprintf(
		"ObjFile( IgnoredLineCount:%d Vertices:%d DefaultGroup:%d Groups:%d CurrentGroupName:%s )",
		o.IgnoredLineCount,
		len(o.Vertices),
		len(o.DefaultGroup.LocalShape.(Group).Children),
		len(o.Groups),
		o.CurrentGroupName,
	)
}

func ParseObjFile(s string) ObjFile {
	of := ObjFile{}
	of.Vertices = []Tuple{NewPoint(0, 0, 0)} // first element is dummy, this is a 1-indexed array
	of.Normals = []Tuple{NewVector(0, 0, 0)} // first element is dummy, this is a 1-indexed array
	defaultGroup := NewGroup()
	of.Groups = map[string]*Shape{
		"": &defaultGroup,
	}
	of.CurrentGroupName = ""
	of.DefaultGroup = of.Groups[""]

	scanner := bufio.NewScanner(strings.NewReader(s))
	var fExtendedRegex = regexp.MustCompile(`(\d+)?/(\d+)?/(\d+)?`)
	var va, vb, vc float64
	var vna, vnb, vnc float64
	var g string
	for scanner.Scan() {
		line := scanner.Text()

		// Scan for vertices
		if n, err := fmt.Sscanf(line, "v %f %f %f", &va, &vb, &vc); err == nil && n == 3 {
			of.Vertices = append(of.Vertices, NewPoint(va, vb, vc))
			continue
		}

		// Scan for vertex normals
		if n, err := fmt.Sscanf(line, "vn %f %f %f", &vna, &vnb, &vnc); err == nil && n == 3 {
			of.Normals = append(of.Normals, NewVector(vna, vnb, vnc))
			continue
		}

		// Scan for faces
		if strings.HasPrefix(line, "f ") {
			faceVertices := []Tuple{}
			faceNormals := []Tuple{}

			tokens := strings.Split(line, " ")
			for _, token := range tokens {
				if foo := fExtendedRegex.FindSubmatch([]byte(token)); len(foo) > 0 {
					vertexIdx, _, normalIdx := foo[1], foo[2], foo[3] // TODO: 2 is the texture index

					if vi, err := strconv.Atoi(string(vertexIdx)); err == nil {
						faceVertices = append(faceVertices, of.Vertices[vi])
					}
					if ni, err := strconv.Atoi(string(normalIdx)); err == nil {
						faceNormals = append(faceNormals, of.Normals[ni])
					}
				} else {
					i, err := strconv.Atoi(token)
					if err == nil {
						faceVertices = append(faceVertices, of.Vertices[i])
					}
				}
			}

			for _, tri := range fanTriangulation(faceVertices, faceNormals) {
				tri.Material.Color = Colors["Red"]
				g := of.Groups[of.CurrentGroupName]
				g.AddChildren(tri)
			}
			continue
		}

		// Scan for groups
		if _, err := fmt.Sscanf(line, "g %s", &g); err == nil {
			of.CurrentGroupName = g
			if _, found := of.Groups[g]; !found {
				group := NewGroup()
				of.Groups[g] = &group
			}
			continue
		}

		of.IgnoredLineCount += 1
	}

	// fmt.Printf("The current group %s has %d children\n", of.CurrentGroupName, len(of.Groups[of.CurrentGroupName].LocalShape.(Group).Children))
	return of
}

// Given a list of vertex indices, where the indices represent
// previously-parsed vertices, divide a polygon into triangles.
// Rotate around polygon face starting with first point:
// A->B->C, A->C->D, A->D->E
//              b
//           /     \
//          /       \
//        a/---------\c
//        \\         /
//         \  \     /
//          \   \  /
//           e----d
// And return the list of triangles
func fanTriangulation(faceVertices, faceNormals []Tuple) []*Shape {
	triangles := []*Shape{}
	for idx := 0; idx < len(faceVertices)-2; idx++ {
		var tri Shape
		if len(faceNormals) > 0 {
			tri = NewSmoothTriangle(
				faceVertices[0],
				faceVertices[idx+1],
				faceVertices[idx+2],
				faceNormals[0],
				faceNormals[idx+1],
				faceNormals[idx+2],
			)
		} else {
			tri = NewTriangle(faceVertices[0], faceVertices[idx+1], faceVertices[idx+2])
		}
		triangles = append(triangles, &tri)
	}
	return triangles
}
