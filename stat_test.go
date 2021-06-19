package gostat

import (
	"fmt"
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
