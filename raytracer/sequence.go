package raytracer

type Sequence struct {
	Numbers      []float64
	currentIndex int
}

// Returns a deterministic number generator.
// TODO: break out into a SequenceInterface, with DeterministicSequence and RandomSequence?
func NewSequence(s ...float64) Sequence {
	return Sequence{s, 0}
}

func (s Sequence) IsEqualTo(s2 Sequence) bool {
	if !equalFloat64Slices(s.Numbers, s2.Numbers) {
		return false
	}
	return true
}

func (s *Sequence) Next() float64 {
	n := s.Numbers[s.currentIndex]
	s.currentIndex += 1
	if s.currentIndex >= len(s.Numbers) {
		s.currentIndex = 0
	}
	return n
}
