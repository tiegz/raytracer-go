package raytracer

type Tuple struct {
  X float64
  Y float64
  Z float64
  W float64 // 0: point, 1: vector
}

func (t *Tuple) Type() string {
  if t.W == 1.0 {
    return "Point"
  } else {
    return "Vector"
  }
}
