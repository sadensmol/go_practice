package main

import "fmt"

type MinStack struct {
	data    []int // data on stack
	lastMin *int  // last min value
	mins    []int // mins
	minats  []int // positions where mins were changed in sync with mins
	//todo add locks
}

func (s *MinStack) Min() int {
	if s.lastMin == nil {
		panic("no min yet")
	}
	return *s.lastMin
}

func (s *MinStack) Pop() int {
	l := len(s.data)
	v := s.data[l-1]
	s.data = s.data[:l-1]

	if s.lastMin != nil {
		mal := len(s.minats)

		if mal > 0 {
			ma := s.minats[mal-1] // last min at position
			// fmt.Printf("last min at position: %d, cur len: %d\n", ma, l)
			if ma == l {
				ml := len(s.mins)

				mv := s.mins[ml-1]
				// fmt.Printf("last min value: %d\n", mv)
				s.lastMin = &mv
				s.mins = s.mins[:ml-1]
				s.minats = s.minats[:mal-1]
			}
		}
	}
	return v
}

func (s *MinStack) Push(v int) {
	s.data = append(s.data, v)

	if s.lastMin == nil || v < *s.lastMin {
		l := len(s.data)
		s.lastMin = &v
		s.mins = append(s.mins, v)
		s.minats = append(s.minats, l)
	}
}

func main() {

	ms := &MinStack{}

	// fmt.Printf("min %v\n", ms.Min())

	ms.Push(10) //10, min 10
	fmt.Printf("min %d\n", ms.Min())
	ms.Push(5) // 10, 5, min 5
	fmt.Printf("min %d\n", ms.Min())
	ms.Push(3) // 10, 5 ,3, min 3
	fmt.Printf("min %d\n", ms.Min())
	ms.Push(20) // 10, 5, 3, 20 min 3
	fmt.Printf("min %d\n", ms.Min())
	v := ms.Pop() // 20 -> 10, 5, 3, min 3
	fmt.Printf("v = %d, min %d\n", v, ms.Min())
	v = ms.Pop() // 3 -> 10, 5, min 5
	fmt.Printf("v = %d, min %d\n", v, ms.Min())
	v = ms.Pop() // 5 -> 10, min 5
	fmt.Printf("v = %d, min %d\n", v, ms.Min())
	v = ms.Pop() // 10 -> min 10 ??
	fmt.Printf("v = %d, min %d\n", v, ms.Min())
}
