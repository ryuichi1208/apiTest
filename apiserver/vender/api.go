func (s *Stream) Count() int {
	return len(s.b) + s.stream.count()
}

func (s *Stream) flush() {
	s.maybeSort()
	s.stream.merge(s.b)
	s.b = s.b[:0]
}

func (s *Stream) maybeSort() {
	if !s.sorted {
		s.sorted = true
		sort.Sort(s.b)
	}
}

func (s *Stream) flushed() bool {
	return len(s.stream.l) > 0
}

type stream struct {
	n float64
	l []Sample
	ƒ invariant
}

func (s *stream) reset() {
	s.l = s.l[:0]
	s.n = 0
}

func (s *stream) insert(v float64) {
	s.merge(Samples{{v, 1, 0}})
}
