package gostat

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestHisto(t *testing.T) {
	s := NewStat(200)
	for i := 0; i < 100_000; i++ {
		s.Add(rand.NormFloat64()*10.0 + 1000.)
	}
	h := s.ToHisto(40)
	fmt.Println(h)
	for _, b := range h.bars {
		if math.IsNaN(b.value) {
			t.Fatal("NaN appeared in histogram")
		}
	}
}
