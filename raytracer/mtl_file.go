package raytracer

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

// MtlFile represents a parsed .mtl file, which is a supplementary
// file to .obj that represents materials it can reference for objects.
// Support:
//   newmtl NAME: define a material with name
//   Ka R G B: define ambient color of the material, 0-1.0
//   Kd R G B: define diffuse color of the material, 0-1.0
//   Ks R G B: define specular color of the material, 0-1.0
//   Ns R G B: define the weighted specular exponent, 0-1000
// Unsupported lines are: d/tr, Tf, Ni,
type MtlFile struct {
	IgnoredLineCount int
	Materials        map[string]Material
}

func (m MtlFile) String() string {
	return fmt.Sprintf(
		"MtlFile(\n  IgnoredLineCount: %d,\n  Materials: %v\n)",
		m.IgnoredLineCount,
		m.Materials,
	)
}

func ParseMtlFile(s string) MtlFile {
	mf := MtlFile{}
	mf.Materials = map[string]Material{}

	scanner := bufio.NewScanner(strings.NewReader(s))
	var currentMaterial Material

	for scanner.Scan() {
		line := scanner.Text()
		var aString string
		var r, g, b float64
		var aFloat float64

		if n, err := fmt.Sscanf(line, "newmtl %s", &aString); err == nil && n == 1 {
			// "Names may be any length but cannot include blanks. Underscores may be used in material names."
			if strings.Contains(aString, " ") {
				mf.IgnoredLineCount += 1
				continue
			}

			// Store the last material, if it existed.
			if currentMaterial.Label != "" {
				mf.Materials[currentMaterial.Label] = currentMaterial
			}
			currentMaterial = Material{Label: aString}
			continue
		}

		// Scan for materials
		if n, err := fmt.Sscanf(line, "Ka %f %f %f", &r, &g, &b); err == nil && n == 3 {
			currentMaterial.Ambient = NewColor(r, g, b)
			continue
		}
		if n, err := fmt.Sscanf(line, "Kd %f %f %f", &r, &g, &b); err == nil && n == 3 {
			currentMaterial.Diffuse = NewColor(r, g, b)
			continue
		}
		if n, err := fmt.Sscanf(line, "Ks %f %f %f", &r, &g, &b); err == nil && n == 3 {
			currentMaterial.Specular = NewColor(r, g, b)
			continue
		}
		if n, err := fmt.Sscanf(line, "Ns %f", &aFloat); err == nil && n == 1 {
			currentMaterial.Shininess = aFloat
			continue
		}
		if n, err := fmt.Sscanf(line, "d %f", &aFloat); err == nil && n == 1 {
			// "d" ("dissolved") is the inverse way of declaring of "Tr" ("transparency")
			currentMaterial.Transparency = 1 - aFloat
			continue
		}
		if n, err := fmt.Sscanf(line, "Tr %f", &aFloat); err == nil && n == 1 {
			currentMaterial.Transparency = aFloat
			continue
		}
		if n, err := fmt.Sscanf(line, "Ni %f", &aFloat); err == nil && n == 1 {
			currentMaterial.RefractiveIndex = aFloat
			continue
		}
		if n, err := fmt.Sscanf(line, "illum "); err == nil && n == 1 {
			log.Fatalln(".mtl file Illumination Model not implemented yet!")
		}
		if n, err := fmt.Sscanf(line, "Tf "); err == nil && n == 1 {
			log.Fatalln(".mtl file Transmission Filter Color not implementeded yet!")
		}
		if n, err := fmt.Sscanf(line, "map_Ka %s", &aString); err == nil && n == 1 {
			// TODO implement Material.AmbientMapFilename for texture maps
			currentMaterial.AmbientMapFilename = aString
			continue
		}
		if n, err := fmt.Sscanf(line, "map_Kd %s", &aString); err == nil && n == 1 {
			// TODO implement Material.DiffuseMapFilename for texture maps
			currentMaterial.DiffuseMapFilename = aString
			continue
		}
		if n, err := fmt.Sscanf(line, "map_Ks %s", &aString); err == nil && n == 1 {
			// TODO implement Material.SpecularMapFilename for texture maps
			currentMaterial.SpecularMapFilename = aString
			continue
		}
		if n, err := fmt.Sscanf(line, "map_Ns %s", &aString); err == nil && n == 1 {
			// TODO implement Material.SpecularMapFilename for texture maps
			currentMaterial.SpecularHighlightMapFilename = aString
			continue
		}
		if n, err := fmt.Sscanf(line, "map_d %s", &aString); err == nil && n == 1 {
			// TODO implement Material.SpecularMapFilename for texture maps
			currentMaterial.AlphaTextureMapFilename = aString
			continue
		}
		if n, err := fmt.Sscanf(line, "map_bump %s", &aString); err == nil && n == 1 {
			// TODO implement Material.SpecularMapFilename for texture maps
			currentMaterial.BumpMapFilename = aString
			continue
		}
		if n, err := fmt.Sscanf(line, "disp %s", &aString); err == nil && n == 1 {
			// TODO implement Material.SpecularMapFilename for texture maps
			currentMaterial.DisplacementMapFilename = aString
			continue
		}
		if n, err := fmt.Sscanf(line, "decal %s", &aString); err == nil && n == 1 {
			// TODO implement Material.SpecularMapFilename for texture maps
			currentMaterial.StencilDecalMapFilename = aString
			continue
		}

		mf.IgnoredLineCount += 1
	}

	// Ensure that the last material parsed gets added to the map.
	mf.Materials[currentMaterial.Label] = currentMaterial

	return mf
}
