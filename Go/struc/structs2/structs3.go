package main

import "fmt"

type Student struct {
	name   string
	grades []float32
	age    int
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAverageGrade() float32 {
	var sum float32 = 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s *Student) getMaxGrade() float32 {
	var maxGrade float32 = 0
	for _, v := range s.grades {
		if v > maxGrade {
			maxGrade = v
		}
	}
	return maxGrade
}

func main() {
	s1 := Student{"Sebas", []float32{3.1, 2.3}, 21}
	s2 := Student{"Pablo", []float32{5.0, 5.0}, 21}
	s1.setAge(2)
	fmt.Println(s1.getAge(), s1.age) //I don't know why, but I can see both

	fmt.Println("Sebas: ", s1.getAverageGrade())
	fmt.Println("Pablo: ", s2.getAverageGrade())
	fmt.Println("The Sebas's max grade is: ", s1.getMaxGrade())
}
