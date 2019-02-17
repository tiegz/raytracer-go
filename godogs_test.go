package main

import (
  "fmt"
  "github.com/DATA-DOG/godog"
)


func (a *Tuple) aTuple(arg1, arg2, arg3, arg4 float64) error {
  a.x = arg1
  a.y = arg2
  a.z = arg3
  a.w = arg4
  return nil
}

func (a *Tuple) ax(arg1 float64) error {
  if a.x != arg1 {
    return fmt.Errorf("expected a.x to be %f but got %f", a.x, arg1)
  } else {
    return nil
  }
}

func (a *Tuple) ay(arg1 float64) error {
  if a.y != arg1 {
    return fmt.Errorf("expected a.y to be %f but got %f", a.y, arg1)
  } else {
    return nil
  }
}

func (a *Tuple) az(arg1 float64) error {
  if a.z != arg1 {
    return fmt.Errorf("expected a.z to be %f but got %f", a.z, arg1)
  } else {
    return nil
  }
}

func (a *Tuple) aw(arg1 float64) error {
  if a.w != arg1 {
    return fmt.Errorf("expected a.w to be %f but got %f", a.w, arg1)
  } else {
    return nil
  }
}

func (a *Tuple) aIsAVector() error {
  if a.w != 0.0 {
    return fmt.Errorf("expected a to be a Vector but was a Point")
  } else {
    return nil
  }
}

func (a *Tuple) aIsNotAPoint() error {
  if a.w == 1.0 {
    return fmt.Errorf("expected a to be not a Point but was a Point")
  } else {
    return nil
  }
}

func (a *Tuple) aIsNotAVector() error {
  if a.w == 0.0 {
    return fmt.Errorf("expected a to be not a Vector but was a Vector")
  } else {
    return nil
  }
}

func (a *Tuple) aIsAPoint() error {
  if a.w != 1.0 {
    return fmt.Errorf("expected a to be a Point but was a Vector")
  } else {
    return nil
  }
}


func FeatureContext(s *godog.Suite) {
  a := Tuple{}

  s.Step(`^a tuple \(([-0-9.]+), ([-0-9.]+), ([-0-9.]+), ([-0-9.]+)\)$`, a.aTuple)
  s.Step(`^a\.x = ([-0-9.]+)$`, a.ax)
  s.Step(`^a\.y = ([-0-9.]+)$`, a.ay)
  s.Step(`^a\.z = ([-0-9.]+)$`, a.az)
  s.Step(`^a\.w = ([-0-9.]+)$`, a.aw)
  s.Step(`^a is a point$`, a.aIsAPoint)
  s.Step(`^a is not a vector$`, a.aIsNotAVector)
  s.Step(`^a is a vector$`, a.aIsAVector)
  s.Step(`^a is not a point$`, a.aIsNotAPoint)
}
