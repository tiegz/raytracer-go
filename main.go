package main

import (
  "./raytracer"
  "fmt"
)

func main() {
  t := raytracer.Tuple{}
  fmt.Printf("Here's a tuple: %s", t.Type())
}
