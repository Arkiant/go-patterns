package singleton

/*
Singleton interface with all available methods
*/
type Singleton interface {
	Sum(number int) int
	GetNumber() int
}

type singleton struct {
	number int
}

var instance *singleton

/*
GetInstance create a new instance or get the current instance
*/
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

/*
Sum a given number
*/
func (s *singleton) Sum(number int) int {
	s.number += number
	return s.number
}

/*
GetNumber get a current number
*/
func (s *singleton) GetNumber() int {
	return s.number
}