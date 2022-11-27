package student

type Student struct {
	ra   int
	name string
}

func New() *Student {
	return &Student{
		ra:   -1,
		name: " ",
	}
}

func NewWithParams(ra int, name string) *Student {
	return &Student{
		ra:   ra,
		name: name,
	}
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) GetRA() int {
	return s.ra
}
