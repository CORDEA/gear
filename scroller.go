package main

import "errors"

type Scroller struct {
	history *History
	index   int
}

func NewScroller(history *History) Scroller {
	return Scroller{history: history, index: len(history.histories)}
}

func (s *Scroller) Reset() {
	s.index = len(s.history.histories)
}

func (s *Scroller) ScrollToBelow() (string, error) {
	if s.index >= len(s.history.histories) {
		return "", nil
	}
	history := s.history.histories[s.index]
	s.index++
	return history, nil
}

func (s *Scroller) ScrollToAbove() (string, error) {
	if s.index <= 0 {
		s.index = 1
		return "", errors.New("reached top")
	}
	s.index--
	return s.history.histories[s.index], nil
}
