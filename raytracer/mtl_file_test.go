package raytracer

import (
	"testing"
)

func TestIgnoringUnrecognizedLines_MtlFile(t *testing.T) {
	gibberish :=
		`There was a young lady named Bright
who traveled much faster than light.
She set out one day
in a relative way,
and came back the previous night.`
	parser := ParseMtlFile(gibberish)
	assertEqualInt(t, 5, parser.IgnoredLineCount)
}

func TestNewMaterial(t *testing.T) {
	file :=
		`newmtl my_mtl
Ka 0.000000 0.000000 0.000000
Kd 0.640000 0.320000 0.160000
Ks 0.500000 0.500000 0.500000
Ns 173.1
Tr 0.4
Ni 1.45
`
	parser := ParseMtlFile(file)

	m, ok := parser.Materials["my_mtl"]
	assert(t, ok)
	assertEqualColor(t, NewColor(0, 0, 0), m.Ambient)
	assertEqualColor(t, NewColor(0.64, 0.32, 0.16), m.Diffuse)
	assertEqualColor(t, NewColor(0.5, 0.5, 0.5), m.Specular)
	assertEqualFloat64(t, 173.1, m.Shininess)
	assertEqualFloat64(t, 0.4, m.Transparency)
	assertEqualFloat64(t, 1.45, m.RefractiveIndex)
}

func TestNewMaterial_ValidName(t *testing.T) {
	file :=
		`newmtl a_regular_name
newmtl areallylongnameareallylongnameareallylongnameareallylongnameareallylongname
newmtl a-name-with-hyphens_and_underscores_and_numbers_1234567890
newmtl 
newmtl some invalid name with spaces
`
	parser := ParseMtlFile(file)
	var actualKeys []string
	for k := range parser.Materials {
		actualKeys = append(actualKeys, k)
	}

	// Re: 4th example, we could also detect these and ignore/warn on them using regexes instead.
	expectedKeys := []string{
		"a_regular_name",
		"areallylongnameareallylongnameareallylongnameareallylongnameareallylongname",
		"a-name-with-hyphens_and_underscores_and_numbers_1234567890",
		"some",
	}

	assertEqualSliceOfStrings(t, expectedKeys, actualKeys)
}

func TestNewMaterial_TransparencyFlags(t *testing.T) {
	file :=
		`newmtl semi_transparent_mtl
d 0.9
newmtl semi_transparent_mtl_2
Tr 0.1
`
	parser := ParseMtlFile(file)

	m, ok := parser.Materials["semi_transparent_mtl"]
	assert(t, ok)
	assertEqualFloat64(t, 0.1, m.Transparency)

	m, ok = parser.Materials["semi_transparent_mtl_2"]
	assert(t, ok)
	assertEqualFloat64(t, 0.1, m.Transparency)
}

func TestParseMtlFile_TextureMap(t *testing.T) {
	file :=
		`newmtl my_mtl
map_Ka my_texture_map.tga
map_Kd my_texture_map.tga
map_Ks my_texture_map.tga

map_Ns my_texture_map_spec.tga
map_d my_texture_map_alpha.tga
map_bump my_texture_map_bump.tga
disp my_texture_map_disp.tga
decal my_texture_map_stencil.tga
`
	parser := ParseMtlFile(file)

	m, ok := parser.Materials["my_mtl"]
	assert(t, ok)

	assertEqualString(t, "my_texture_map.tga", m.AmbientMapFilename)
	assertEqualString(t, "my_texture_map.tga", m.DiffuseMapFilename)
	assertEqualString(t, "my_texture_map.tga", m.SpecularMapFilename)
	assertEqualString(t, "my_texture_map_spec.tga", m.SpecularHighlightMapFilename)
	assertEqualString(t, "my_texture_map_alpha.tga", m.AlphaTextureMapFilename)
	assertEqualString(t, "my_texture_map_bump.tga", m.BumpMapFilename)
	assertEqualString(t, "my_texture_map_disp.tga", m.DisplacementMapFilename)
	assertEqualString(t, "my_texture_map_stencil.tga", m.StencilDecalMapFilename)
}
