package singleton

type singleton struct {
	count int
	name  string
}

var instance *singleton

func GetInstance(name string) *singleton {
	if instance == nil {
		instance = new(singleton)
		instance.name = name 
	}

	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

func (s *singleton) GetCount() int {
	return s.count
}
