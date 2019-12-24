# raytracer-go

A raytracer in Go.

This is an implementation of ["The Ray Tracer Challenge" by Jamis Buck](https://pragprog.com/book/jbtracer/the-ray-tracer-challenge), which teaches you how to write a raytracer using language-agnostic BDD/[Cucumber](https://cucumber.io/) tests.

# Testing

- Run all tests: `go test ./...`
- Run all benchmarks: `go test ./... -bench=.`
  - Benchmarking philosophy: focus on methods that are crucial to rendering and could possibly be affected by regressions. Avoid benchmarking things that are simple Go operations (+-/\*) and/or won't change much (e.g. `Tuple.Add()`).
