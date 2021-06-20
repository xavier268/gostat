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

	s := New(10)
	// fmt.Println(s)
	for _, i := range data {
		s.Add(i)
		// fmt.Println(s)
	}
	fmt.Println(s)

}

func TestDist1(t *testing.T) {
	fmt.Println("Uniform distribution")
	s := New(20)
	for i := 0; i < 10_000; i++ {
		s.Add(rand.Float64() * 1000.)
	}
	fmt.Println(s)
}

func TestDist2(t *testing.T) {
	fmt.Println("Two separate uniform distribution")
	s := New(30)
	for i := 0; i < 10_000; i++ {
		s.Add(rand.Float64()*1000. - 5000.)
		s.Add(rand.Float64() * 10000.)
	}
	fmt.Println(s)
}

func TestDist3(t *testing.T) {
	fmt.Println("Skewd distribution")
	s := New(30)
	for i := 0; i < 100_000; i++ {
		s.Add(-math.Log(1 - rand.Float64()))

	}
	fmt.Println(s)
}
