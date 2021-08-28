package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance("number 5")
	if instance1 == nil {
		//Test 1 failed
		t.Error("A new connection object must have been made")
	}
	expectedCounter := instance1

	currentCount := instance1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}

	instance2 := GetInstance("is alive") // should return same instance already created with first GetInstance
	if instance2 != expectedCounter {
		//Test 2 failed
		t.Error("Singleton instances must be different")
	}

	if instance2.name == "is alive" {
		t.Error("I should be 'number 5', but I'm not!!!")
	}

	if instance2.name != "number 5" {
		t.Error("Singleton is not being created correct- I HAVE TO BE 'number 5'!!!!")
	}
	
	currentCount = instance2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the current count must be 2 but was %d\n", currentCount)
	}
}
