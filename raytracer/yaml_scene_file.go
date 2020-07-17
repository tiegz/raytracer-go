package raytracer

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

type YamlSceneFile struct {
	Camera             Camera
	World              World
	MaterialDefs       map[string]Material
	TransformationDefs map[string]Matrix
}

type YamlInstruction struct {
	// Type of instruction
	Define string
	Add    string

	// Extra fields
	Extend      string
	Width       int
	Height      int
	FieldOfView float64 `yaml:"field-of-view"`
	From        [3]float64
	To          [3]float64
	Up          [3]float64
	At          [3]float64
	Intensity   [3]float64
	Transform   yaml.Node
	Material    yaml.Node
	Value       yaml.Node
}

// NB: using pointers instead of values because values that
// are missing from the YAML struct would be initialized as
// their zero values (e.g. 0.0), but we don't want to use
// 0.0 if they're actually just missing.
type YamlMaterial struct {
	Color           [3]*float64
	Diffuse         *float64
	Ambient         *float64
	Specular        *float64
	Reflective      *float64
	Shininess       *float64
	RefractiveIndex *float64 `yaml:"refractive-index"`
	Transparency    *float64
}

// This returns a World as parsed from YAML, based on the format in the book.
func NewYamlSceneFile() YamlSceneFile {
	return YamlSceneFile{
		World:              NewWorld(),
		MaterialDefs:       map[string]Material{},
		TransformationDefs: map[string]Matrix{},
	}
}

func ParseYamlSceneFile(filename string) (YamlSceneFile, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	ysf := NewYamlSceneFile()
	y := []YamlInstruction{}
	err = yaml.Unmarshal([]byte(file), &y)
	if err != nil {
		return ysf, err
	}

	for _, instruction := range y {
		if instruction.Add != "" {
			switch instruction.Add {
			case "camera":
				ysf.Camera = NewCamera(instruction.Width, instruction.Height, instruction.FieldOfView)
				ysf.Camera.SetTransform(NewViewTransform(
					NewPoint(instruction.From[0], instruction.From[1], instruction.From[2]),
					NewPoint(instruction.To[0], instruction.To[1], instruction.To[2]),
					NewVector(instruction.Up[0], instruction.Up[1], instruction.Up[2]),
				))
			case "light":
				ysf.World.Lights = append(
					ysf.World.Lights,
					NewPointLight(
						NewPoint(instruction.At[0], instruction.At[1], instruction.At[2]),
						NewColor(instruction.Intensity[0], instruction.Intensity[1], instruction.Intensity[2]),
					),
				)
			case "plane", "sphere", "cube":
				var obj *Shape
				var t Matrix
				var m Material
				switch instruction.Add {
				case "plane":
					obj = NewPlane()
					obj.Label = "plane"
				case "sphere":
					obj = NewSphere()
					obj.Label = "sphere"
				case "cube":
					obj = NewCube()
					obj.Label = "cube"
				}
				if m, err = decodeMaterial(ysf.MaterialDefs, DefaultMaterial(), instruction.Material); err != nil {
					return ysf, err
				}
				obj.Material = m
				if t, err = decodeTransforms(ysf.TransformationDefs, instruction.Transform); err != nil {
					return ysf, err
				}
				obj.SetTransform(obj.Transform.Multiply(t))
				ysf.World.Objects = append(ysf.World.Objects, obj)
			default:
				return ysf, fmt.Errorf("Unknown instruction: %s\n", instruction.Add)
			}
		} else if instruction.Define != "" {
			var m Material
			var t Matrix
			if strings.HasSuffix(instruction.Define, "-material") {
				// TODO: patterns
				if instruction.Extend != "" {
					m = ysf.MaterialDefs[instruction.Extend]
				} else {
					m = DefaultMaterial()
				}
				m.Label = instruction.Define
				if m, err = decodeMaterial(ysf.MaterialDefs, m, instruction.Value); err != nil {
					return ysf, err
				}
				ysf.MaterialDefs[instruction.Define] = m
			} else if strings.HasSuffix(instruction.Define, "-transform") || strings.HasSuffix(instruction.Define, "-object") {
				if t, err = decodeTransforms(ysf.TransformationDefs, instruction.Value); err != nil {
					return ysf, err
				}
				ysf.TransformationDefs[instruction.Define] = t
			}
		}
	}

	return ysf, nil
}

func decodeFloat64(a interface{}) (float64, error) {
	f, ok := a.(float64)
	if !ok {
		i, ok := a.(int)
		if !ok {
			return f, fmt.Errorf("Cannot convert %s to float64 or int.\n", a)
		}
		f = float64(i)
	}

	return f, nil
}

func decodeFloat64Tuple(a, b, c interface{}) (float64, float64, float64, error) {
	var x, y, z float64
	var err error
	if x, err = decodeFloat64(a); err != nil {
		return 0.0, 0.0, 0.0, err
	}
	if y, err = decodeFloat64(b); err != nil {
		return 0.0, 0.0, 0.0, err
	}
	if z, err = decodeFloat64(c); err != nil {
		return 0.0, 0.0, 0.0, err
	}
	return x, y, z, nil
}

func decodeString(n yaml.Node) (string, error) {
	var val string
	if err := n.Decode(&val); err != nil {
		return "", err
	}
	return val, nil
}

func decodeMixedArray(n yaml.Node) ([]interface{}, error) {
	var val []interface{}
	if err := n.Decode(&val); err != nil {
		return val, err
	}
	return val, nil
}

func decodeYamlNodeArray(n yaml.Node) ([]yaml.Node, error) {
	val := []yaml.Node{}
	if err := n.Decode(&val); err != nil {
		return val, err
	}
	return val, nil
}

func decodeMaterial(defs map[string]Material, m Material, n yaml.Node) (Material, error) {
	// 1: try to decode as definition string
	if defKey, err := decodeString(n); err == nil {
		if def, ok := defs[defKey]; ok {
			m = def
		} else {
			return m, fmt.Errorf("Material definition not found: %s\n", defKey)
		}
		// 2: try to decode as YamlMaterial
	} else {
		var v YamlMaterial
		if err := n.Decode(&v); err != nil {
			return m, err
		}
		if v.Color[0] != nil {
			r, g, b, err := decodeFloat64Tuple(*v.Color[0], *v.Color[1], *v.Color[2])
			if err != nil {
				return m, err
			}
			m.Color = NewColor(r, g, b)
		}
		if v.Diffuse != nil {
			m.Diffuse = *v.Diffuse
		}
		if v.Ambient != nil {
			m.Ambient = *v.Ambient
		}
		if v.Specular != nil {
			m.Specular = *v.Specular
		}
		if v.Reflective != nil {
			m.Reflective = *v.Reflective
		}
		if v.Shininess != nil {
			m.Shininess = *v.Shininess
		}
		if v.Transparency != nil {
			m.Transparency = *v.Transparency
		}
		if v.RefractiveIndex != nil {
			m.RefractiveIndex = *v.RefractiveIndex
		}
	}
	return m, nil
}

func decodeTransforms(defs map[string]Matrix, n yaml.Node) (Matrix, error) {
	t := IdentityMatrix()
	nodeArray, err := decodeYamlNodeArray(n)

	if err != nil {
		return t, err
	}

	for _, transform := range nodeArray { // will be either string or array of instructions
		// 1: try to decode as definition string
		if defKey, err := decodeString(transform); err == nil {
			def, ok := defs[defKey]
			if !ok {
				return t, fmt.Errorf("Transform definition not found: %s\n", defKey)
			}
			t = t.Compose(def)
		} else {
			// 2: try to decode as YamlMaterial
			transformParts, err := decodeMixedArray(transform)
			if err != nil {
				return t, err
			}
			transformType := transformParts[0]
			switch transformType {
			case "translate":
				x, y, z, err := decodeFloat64Tuple(transformParts[1], transformParts[2], transformParts[3])
				if err != nil {
					return t, err
				}
				t = t.Compose(NewTranslation(x, y, z))
			case "scale":
				x, y, z, err := decodeFloat64Tuple(transformParts[1], transformParts[2], transformParts[3])
				if err != nil {
					return t, err
				}
				t = t.Compose(NewScale(x, y, z))
			case "shear":
				xToY, xToZ, yToZ, err := decodeFloat64Tuple(transformParts[1], transformParts[2], transformParts[3])
				if err != nil {
					return t, err
				}
				yToZ, zToX, zToY, err := decodeFloat64Tuple(transformParts[1], transformParts[2], transformParts[3])
				if err != nil {
					return t, err
				}
				t = t.Compose(NewShear(xToY, xToZ, yToZ, yToZ, zToX, zToY))
			case "rotate-x":
				deg := transformParts[1].(float64)
				t = t.Compose(NewRotateX(deg))
			case "rotate-y":
				deg := transformParts[1].(float64)
				t = t.Compose(NewRotateY(deg))
			case "rotate-z":
				deg := transformParts[1].(float64)
				t = t.Compose(NewRotateZ(deg))
			default:
				return t, fmt.Errorf("Transformation not implemented for '%s'\n", transformType)
			}
		}
	}
	return t, nil
}
