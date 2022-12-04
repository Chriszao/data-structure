package student

type Student struct {
	name string
	ra   int
}

func New() *Student {
	return &Student{
		name: " ",
		ra:   -1,
	}
}

func NewWithParams(ra int, name string) *Student {
	return &Student{
		name: name,
		ra:   ra,
	}
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) GetRa() int {
	return s.ra
}
