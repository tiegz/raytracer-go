package main

import (
  "./raytracer"
  "fmt"
  "github.com/DATA-DOG/godog"
)


type TestSubject struct {
  ATuple raytracer.Tuple
}

func (subject *TestSubject) aTuple(arg1, arg2, arg3, arg4 float64) error {
  subject.ATuple.X = arg1
  subject.ATuple.Y = arg2
  subject.ATuple.Z = arg3
  subject.ATuple.W = arg4
  return nil
}

func (subject *TestSubject) ax(arg1 float64) error {
  if subject.ATuple.X != arg1 {
    return fmt.Errorf("expected X to be %f but got %f", subject.ATuple.X, arg1)
  } else {
    return nil
  }
}

func (subject *TestSubject) ay(arg1 float64) error {
  if subject.ATuple.Y != arg1 {
    return fmt.Errorf("expected Y to be %f but got %f", subject.ATuple.Y, arg1)
  } else {
    return nil
  }
}

func (subject *TestSubject) az(arg1 float64) error {
  if subject.ATuple.Z != arg1 {
    return fmt.Errorf("expected Z to be %f but got %f", subject.ATuple.Z, arg1)
  } else {
    return nil
  }
}

func (subject *TestSubject) aw(arg1 float64) error {
  if subject.ATuple.W != arg1 {
    return fmt.Errorf("expected W to be %f but got %f", subject.ATuple.W, arg1)
  } else {
    return nil
  }
}

func (subject *TestSubject) aIsAVector() error {
  if subject.ATuple.Type() != "Vector" {
    return fmt.Errorf("expected a to be a Vector but was a Point")
  } else {
    return nil
  }
}

func (subject *TestSubject) aIsNotAPoint() error {
  if subject.ATuple.Type() == "Point" {
    return fmt.Errorf("expected a to be not a Point but was a Point")
  } else {
    return nil
  }
}

func (subject *TestSubject) aIsNotAVector() error {
  if subject.ATuple.Type() == "Vector" {
    return fmt.Errorf("expected a to be not a Vector but was a Vector")
  } else {
    return nil
  }
}

func (subject *TestSubject) aIsAPoint() error {
  if subject.ATuple.Type() != "Point" {
    return fmt.Errorf("expected a to be a Point but was a Vector")
  } else {
    return nil
  }
}

func FeatureContext(s *godog.Suite) {
  testSubject := TestSubject{}

  s.Step(`^a tuple \(([-0-9.]+), ([-0-9.]+), ([-0-9.]+), ([-0-9.]+)\)$`, testSubject.aTuple)
  s.Step(`^a\.x = ([-0-9.]+)$`, testSubject.ax)
  s.Step(`^a\.y = ([-0-9.]+)$`, testSubject.ay)
  s.Step(`^a\.z = ([-0-9.]+)$`, testSubject.az)
  s.Step(`^a\.w = ([-0-9.]+)$`, testSubject.aw)
  s.Step(`^a is a point$`, testSubject.aIsAPoint)
  s.Step(`^a is not a vector$`, testSubject.aIsNotAVector)
  s.Step(`^a is a vector$`, testSubject.aIsAVector)
  s.Step(`^a is not a point$`, testSubject.aIsNotAPoint)
}

