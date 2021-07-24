package gostat

import (
	"fmt"
	"math"
	"strings"
)

// bar represent a single bar of the histogram.
type bar struct {
	start, end, value float64
}

type Histo struct {
	bars []bar
	nob  int
	vmax float64 // max bar values

}

// ToLinearHisto creates an histogram with the provided number of bars (nob) of equal size,
// from the minimum to the maximum
func (s *Stat) ToHisto(nob int) *Histo {

	if nob < 5 {
		nob = 5
	}

	h := new(Histo)
	h.nob = nob

	h.bars = make([]bar, h.nob)
	h.bars[0].start = s.Min()
	h.bars[nob-1].end = s.Max()
	step := (s.Max() - s.Min()) / float64(nob)
	cur := s.Min() + step
	for i := 1; i < h.nob; i++ {
		h.bars[i-1].end, h.bars[i].start = cur, cur
		cur += step
	}
	for i, b := range h.bars {
		h.bars[i].value = s.NRepart(b.end) - s.NRepart(b.start)
		h.vmax = math.Max(h.bars[i].value, h.vmax)
	}
	return h
}

func (h *Histo) String() string {

	var sb strings.Builder
	for _, b := range h.bars {
		fmt.Fprintf(&sb, "[%#5.f\t\t%#5.f]\t\t%#5.f\t", b.start, b.end, b.value)
		for i := 0.0; i < b.value; i += h.vmax / 80. {
			fmt.Fprintf(&sb, "*")
		}

		fmt.Fprintln(&sb)
	}
	return sb.String()
}
