package raytracer

import (
	"fmt"
	"sync"
)

type Sequence struct {
	Numbers      []float64
	currentIndex int
	mutex        sync.Mutex
}

// Returns a deterministic number generator.
// TODO: break out into a SequenceInterface, with DeterministicSequence and RandomSequence?
func NewSequence(s ...float64) Sequence {
	return Sequence{s, 0, sync.Mutex{}}
}

func (s Sequence) String() string {
	return fmt.Sprintf(
		"Shape(\n  Numbers: %v\n  currentIndex: %v\n)",
		s.Numbers,
		s.currentIndex,
	)
}

func (s Sequence) IsEqualTo(s2 Sequence) bool {
	if !equalFloat64Slices(s.Numbers, s2.Numbers) {
		return false
	}
	return true
}

func (s *Sequence) Next() float64 {
	s.mutex.Lock()
	n := s.Numbers[s.currentIndex]
	s.currentIndex += 1
	if s.currentIndex >= len(s.Numbers) {
		s.currentIndex = 0
	}
	s.mutex.Unlock()
	return n
}
