package main

import (
  "github.com/tiegz/raytracer-go/raytracer"
  "fmt"
)

func main() {
  t := raytracer.Tuple{}
  fmt.Printf("Here's a tuple: %s\n", t.Type())
}
