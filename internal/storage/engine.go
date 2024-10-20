package storage

type Storage struct {
	data map[string]string
}

func NewStorage() *Storage {
	return &Storage{data: make(map[string]string)}
}

func (s *Storage) Get(key string) (string, bool) {
	val, ok := s.data[key]
	return val, ok
}

func (s *Storage) Set(key, value string) {
	s.data[key] = value
}

func (s *Storage) Delete(key string) {
	delete(s.data, key)
}
