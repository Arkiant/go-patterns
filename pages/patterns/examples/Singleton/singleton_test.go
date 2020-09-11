package singleton

import "testing"

func TestSingletonPattern(t *testing.T) {
	
	number := GetInstance()

	if number == nil {
		t.Errorf("expected pointer to Singleton after calling GetInstance(), not nil")
	}

	number.Sum(4)
	otherNumber := GetInstance()
	currentNumber := otherNumber.GetNumber()


	if currentNumber != 4 {
		t.Errorf("Ater calling Sum(4) the number must be 4 but it is %d\n", currentNumber)
	}
}