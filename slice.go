package gorgonia

import "gorgonia.org/tensor"

type sli struct {
	start, end, step int
}

func S(start int, opt ...int) tensor.Slice { _ = "STUB: not implemented"; return *new(tensor.Slice) }

func (s *sli) Start() int { _ = "STUB: not implemented"; return 0 }
func (s *sli) End() int   { _ = "STUB: not implemented"; return 0 }
func (s *sli) Step() int  { _ = "STUB: not implemented"; return 0 }
