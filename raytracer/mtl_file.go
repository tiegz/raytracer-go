package raytracer

// import (
// 	"bufio"
// 	"fmt"
// 	"strings"
// )

// // MtlFile represents a parsed .mtl file, which is a supplementary
// // file to .obj that represents materials it can reference for objects.
// // Support:
// //   newmtl NAME: define a material with name
// //   Ka R G B: define ambient color of the material, 0-1.0
// //   Kd R G B: define diffuse color of the material, 0-1.0
// //   Ks R G B: define specular color of the material, 0-1.0
// //   Ns R G B: define the weighted specular exponent, 0-1000
// // Unsupported lines are: d/tr, Tf, Ni,
// type MtlFile struct {
// 	IgnoredLineCount int
// 	Materials        map[string]Material
// 	// Vertices         []Tuple
// 	// Normals          []Tuple
// 	// DefaultGroup     *Shape
// 	// Groups           map[string]*Shape
// 	// CurrentGroupName string
// }

// // func (of *ObjFile) ToGroup() *Shape {
// // 	g := NewGroup()
// // 	for _, v := range of.Groups {
// // 		g.AddChildren(v)
// // 	}
// // 	return g
// // }

// func (m MtlFile) String() string {
// 	return fmt.Sprintf(
// 		"MtlFile(\n  IgnoredLineCount: %d,\n  Materials: %d\n)",
// 		m.IgnoredLineCount,
// 		m.Materials,
// 		// len(o.Vertices),
// 		// len(o.DefaultGroup.LocalShape.(Group).Children),
// 		// len(o.Groups),
// 		// o.CurrentGroupName,
// 	)
// }

// func ParseMtlFile(s string) MtlFile {
// 	mf := MtlFile{}

// 	scanner := bufio.NewScanner(strings.NewReader(s))
// 	// var fExtendedRegex = regexp.MustCompile(`(\d+)?/(\d+)?/(\d+)?`)
// 	// var ka, kb, c float64
// 	// var vna, vnb, vnc float64
// 	// var g string
// 	var currentMaterial Material

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		var name string
// 		var r, g, b float64

// 		if n, err := fmt.Sscanf(line, "newmtl %s", &name); err == nil && n == 1 {
// 			currentMaterial = Material{Label: name}
// 			continue
// 		}

// 		// Scan for materials
// 		if n, err := fmt.Sscanf(line, "Ka %f %f %f", &r, &g, &b); err == nil && n == 3 {
// 			currentMaterial.Ambient = NewColor(r, g, b)
// 			// of.Vertices = append(mf.Vertices, NewPoint(va, vb, vc))
// 			continue
// 		}

// 		of.IgnoredLineCount += 1
// 	}

// 	// of := ObjFile{}
// 	// of.Vertices = []Tuple{NewPoint(0, 0, 0)} // first element is dummy, this is a 1-indexed array
// 	// of.Normals = []Tuple{NewVector(0, 0, 0)} // first element is dummy, this is a 1-indexed array
// 	// defaultGroup := NewGroup()
// 	// of.Groups = map[string]*Shape{
// 	// 	"": defaultGroup,
// 	// }
// 	// of.CurrentGroupName = ""
// 	// of.DefaultGroup = of.Groups[""]

// 	// scanner := bufio.NewScanner(strings.NewReader(s))
// 	// var fExtendedRegex = regexp.MustCompile(`(\d+)?/(\d+)?/(\d+)?`)
// 	// var va, vb, vc float64
// 	// var vna, vnb, vnc float64
// 	// var g string
// 	// for scanner.Scan() {
// 	// 	line := scanner.Text()

// 	// 	// Scan for vertices
// 	// 	if n, err := fmt.Sscanf(line, "v %f %f %f", &va, &vb, &vc); err == nil && n == 3 {
// 	// 		of.Vertices = append(of.Vertices, NewPoint(va, vb, vc))
// 	// 		continue
// 	// 	}

// 	// 	// Scan for vertex normals
// 	// 	if n, err := fmt.Sscanf(line, "vn %f %f %f", &vna, &vnb, &vnc); err == nil && n == 3 {
// 	// 		of.Normals = append(of.Normals, NewVector(vna, vnb, vnc))
// 	// 		continue
// 	// 	}

// 	// 	// Scan for faces
// 	// 	if strings.HasPrefix(line, "f ") {
// 	// 		faceVertices := []Tuple{}
// 	// 		faceNormals := []Tuple{}

// 	// 		tokens := strings.Split(line, " ")
// 	// 		for _, token := range tokens {
// 	// 			if foo := fExtendedRegex.FindSubmatch([]byte(token)); len(foo) > 0 {
// 	// 				vertexIdx, _, normalIdx := foo[1], foo[2], foo[3] // TODO: 2 is the texture index

// 	// 				if vi, err := strconv.Atoi(string(vertexIdx)); err == nil {
// 	// 					faceVertices = append(faceVertices, of.Vertices[vi])
// 	// 				}
// 	// 				if ni, err := strconv.Atoi(string(normalIdx)); err == nil {
// 	// 					faceNormals = append(faceNormals, of.Normals[ni])
// 	// 				}
// 	// 			} else {
// 	// 				i, err := strconv.Atoi(token)
// 	// 				if err == nil {
// 	// 					faceVertices = append(faceVertices, of.Vertices[i])
// 	// 				}
// 	// 			}
// 	// 		}

// 	// 		for _, tri := range fanTriangulation(faceVertices, faceNormals) {
// 	// 			tri.Material.Color = Colors["Red"]
// 	// 			g := of.Groups[of.CurrentGroupName]
// 	// 			g.AddChildren(tri)
// 	// 		}
// 	// 		continue
// 	// 	}

// 	// 	// Scan for groups
// 	// 	if _, err := fmt.Sscanf(line, "g %s", &g); err == nil {
// 	// 		of.CurrentGroupName = g
// 	// 		if _, found := of.Groups[g]; !found {
// 	// 			group := NewGroup()
// 	// 			of.Groups[g] = group
// 	// 		}
// 	// 		continue
// 	// 	}

// 	// }

// 	// // fmt.Printf("The current group %s has %d children\n", of.CurrentGroupName, len(of.Groups[of.CurrentGroupName].LocalShape.(Group).Children))
// 	// return of
// }
