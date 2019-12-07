# TODOs

### Add some runnable example functions that would be useful in the docs

### Add indentation-nesting to String() methods so we can truly pretty-print them, e.g. String(indentation_level int)

### Work around method chaining problem

### Can we replace the NewX methods with just literals? Is there much constructor logic even needed?

### How to get test helpers to print out the line where the test actually failed, rather than the test helper file?

### Can we replace assertEqualXXX with a single assertEqualObject helper, using reflect.TypeOf and reflect.ValueOf?

### Can the testing methods just be instance methods on testing.T? e.g. t.assertEqual(e, a)

For example:

```go
  something := thing.DoSomething()
  something = something.DoAnotherThing()
  something = something.DoSomethingElse()
```

Would be nice to enable method chaining here. If we try that now, we get:

`"cannot call pointer method on thing.DoSomething()"`

### Can we remove the setters, since the properties are all public anyway? (e.g. shape.SetMaterial(m) -> shape.Material = m)

### Turn most method receivers into pointers?

- https://golang.org/doc/faq#methods_on_values_or_pointers
- Using value receivers is better for concurrency
- Especially Tuple, which we could then chain when manipulating, instead of creating intermediate variables

### Replace shape types with shape functions?

- you could test equality by just comparing the normalAt/intersect functions

### Fix Checker/Sphere rendering by implementing UV mapping:

> To apply a two-dimensional texture (like checkers) to the surface of an object, you need to implement something called UV mapping, which converts a three-dimensional point of intersection (x, y, z) into a two-dimensional surface coordinate (u, v). You’d then map that surface coordinate to a color. It’s fun to do, but sadly beyond the scope of this book. Tutorial-style resources are hard to find, but with a bit of reading between the lines and some experimentation, searching for topics like “spherical texture mapping” can bear fruit.

### More pattern ideas: Radial Gradient Patterns, Nested Patterns, Blended Patterns, Perturbed Patterns

### Optimizations ideas:

- create benchmarks for different levels/scopes (world, camera, shape, matrix, tuple)
- reset to non-pointers, and then try making each field-by-field a pointer (get rid of NullX types first?)
- sync.Pool
- goroutines

### Flags for main.go

- options: run examples, run a specific file, etc.
- or just bin commands?

### Fill in the glossary

### Cache the inverse transform for a speedup (convert Transform to Getter/Setter)

### Define where pointers/values should be used as a design (e.g. values for Matrix, pointers for Shapes, etc)

### Give shapes a label? For String(), etc

### Panic or return an error if NewScale() is given a zero value? (which causes a panic bc matrix is not invertible)

### Parsing OBJ files: look into storing them as triangle meshes instead of groups, which will be faster (p2123)
