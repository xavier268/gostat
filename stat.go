package gostat

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Stat struct {
	bkts buckets // bucket currently filled, always ORDERED. Buckets are never empty. Their limit must NEVER touch and they should never overlap.
	nbkt int     // expected number of buckets
}

// New with specified number of internal buckets
func New(precision int) *Stat {
	s := new(Stat)
	s.nbkt = precision
	s.bkts = make([]bucket, 0, s.nbkt+1) // prepare capacity for performance
	return s
}

// String implements Stringer interface, for debugging.
func (s *Stat) String() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "Bucket dump (%d / %dbuckets)\n", s.bkts.Len(), s.nbkt)
	c, m, v := s.CountMeanVar()
	fmt.Fprintf(&sb, "Count\t%d\nMean\t%f \nVar\t%f \n", c, m, v)
	fmt.Fprintf(&sb, "Min\t%f\nMax\t%f \n", s.Min(), s.Max())
	fmt.Fprintf(&sb, "\t%s\n", bucket{}.Header())
	for i, b := range s.bkts {
		fmt.Fprintf(&sb, "%d\t%s\n", i, b.String())
	}

	return sb.String()
}

// Add any scalar value (in, uint, float, ..)
func (s *Stat) Add(data interface{}) {

	switch v := data.(type) {
	case int:
		s.add(float64(v))
	case int8:
		s.add(float64(v))
	case int16:
		s.add(float64(v))
	case int32:
		s.add(float64(v))
	case int64:
		s.add(float64(v))
	case uint:
		s.add(float64(v))
	case uint8:
		s.add(float64(v))
	case uint16:
		s.add(float64(v))
	case uint32:
		s.add(float64(v))
	case uint64:
		s.add(float64(v))
	case float32:
		s.add(float64(v))
	case float64:
		s.add(float64(v))
	default:
		panic("Invalid type added to Stat object")
	}
}

// add a float64
func (s *Stat) add(d float64) {

	// try to put data in an existing bucket
	for i, b := range s.bkts {
		if b.high() > d {
			break // will never fit in any existing bucket
		}
		if b.contains(d) { // found suitable bucket
			s.bkts[i].add(d)
			return
		}
	}

	// create a dedicated bucket for this data that could not fit anywhere
	b := bucket{}
	b.c = d
	b.add(d)
	s.bkts = append(s.bkts, b)
	sort.Sort(s.bkts)

	// if bucket count is still reasonably low, we're done !
	if s.bkts.Len() <= s.nbkt {
		return
	}

	//fmt.Println(s)
	//fmt.Println("Merge required")

	// Here, we have too many buckets - we select the most attractive merge move and do it !
	bi := 0            // best move index so far
	bc := math.Inf(+1) // best cost so far
	for i := 0; i < s.bkts.Len()-1; i++ {
		if c := s.bkts.eval(i); c < bc {
			bc = c
			bi = i
		}
	}
	// do the "best move"
	// fmt.Println("Merging buckets :", bi, " and ", bi+1)
	s.bkts = s.bkts.merge(bi)
	//fmt.Println(s)
	// Done !
}

// CountMeanVar provides exact values for count, mean, variance - more efficient by calculating all values at once.
func (s *Stat) CountMeanVar() (int, float64, float64) {
	var m, v float64
	var n int

	for _, b := range s.bkts {
		m += b.sum
		v += b.sum2
		n += b.n
	}
	m = m / float64(n)
	v = v/float64(n-1) - m*m
	return n, m, v
}

// Count provides exact values for count
func (s *Stat) Count() int {

	var n int
	for _, b := range s.bkts {
		n += b.n
	}
	return n
}

// Min provides exact minimum value
func (s *Stat) Min() float64 {
	return s.bkts[0].low()
}

// Min provides exact maximum value
func (s *Stat) Max() float64 {
	return s.bkts[len(s.bkts)-1].high()
}

// NRepart gives an estimate f the number of data points that are below x (special rounding  for x = c), assuming  GAUSSIAN law.
func (s *Stat) NRepart(x float64) float64 {
	var res float64

	for _, b := range s.bkts {
		res += b.NRepart(x)
	}
	return res
}

// Percentile provides the value such that p percent of the value are below, and 1-P and above.
func (s *Stat) Percentile(p float64) float64 {

	if p < .0 || p > 1.0 {
		panic("Invalid input")
	}

	// TOD0 - idea is to first locate the bucket that contanins the value,
	// then solve using NPart to find the exact percentile.

	panic("Not implemented")
}
