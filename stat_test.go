package gostat

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestDump(t *testing.T) {

	data := []int{2, 5, 4, 7, 1, 6, 1, 4, 4, 4, 2, -20, 40, -50, 100, -70}
	fmt.Println("Input is : ", data)

	s := NewStat(10)
	// fmt.Println(s)
	for _, i := range data {
		s.Add(i)
		// fmt.Println(s)
	}
	fmt.Println(s)
}

func BenchmarkAdd0(b *testing.B) {
	s := NewStat(20)
	b.ResetTimer()
	for i := 0.; i < float64(b.N); i++ {
		s.add(i)
	}
}

func BenchmarkAdd1(b *testing.B) {
	s := NewStat(20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}

func TestDist1(t *testing.T) {
	fmt.Println("Uniform distribution")
	s := NewStat(20)
	for i := 0; i < 10_000; i++ {
		s.Add(rand.Float64())
	}
	fmt.Println(s)
}

func TestDist2(t *testing.T) {
	fmt.Println("Two separate uniform distribution")
	s := NewStat(30)
	for i := 0; i < 10_000; i++ {
		s.Add(rand.Float64()*1000. - 5000.)
		s.Add(rand.Float64() * 10000.)
	}
	fmt.Println(s)
}

func TestDist3(t *testing.T) {
	fmt.Println("Skewd distribution")
	s := NewStat(30)
	for i := 0; i < 100_000; i++ {
		s.Add(-math.Log(1 - rand.Float64()))

	}
	fmt.Println(s)
}

func TestRepartNormal(t *testing.T) {
	s := NewStat(10)
	for i := 0; i < 100; i++ {
		s.Add(200. + 60.*rand.NormFloat64())
	}
	fmt.Println(s)

	for i := -1000.; i < 1000.; i += 20 {
		fmt.Printf("%5f\t==REPART==>\t%5f\n", i, s.NRepart(i))
	}
}
func TestRepartUniform(t *testing.T) {
	s := NewStat(10)
	for i := 0; i < 100; i++ {
		s.Add(200. + 60.*float64(i))
	}
	fmt.Println(s)

	for i := -1000.; i < 1000.; i += 20 {
		fmt.Printf("%5f\t==REPART==>\t%5f\n", i, s.NRepart(i))
	}
}
