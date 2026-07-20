package gorgonia

import (
	"gorgonia.org/tensor"
)

type Solver interface {
	Step([]ValueGrad) error
}

type ValueGrad interface {
	Valuer
	Grad() (Value, error)
}

type Namer interface {
	Name() string
}

func newCachedDV(n ValueGrad, weights, grad Value, zero bool) (cached *dualValue, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractWeightGrad(n ValueGrad) (weights, grad Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), *new(Value), nil
}

type SolverOpt func(s Solver)

func WithL2Reg(l2reg float64) SolverOpt { _ = "STUB: not implemented"; return *new(SolverOpt) }

func WithL1Reg(l1reg float64) SolverOpt { _ = "STUB: not implemented"; return *new(SolverOpt) }

func WithBatchSize(batch float64) SolverOpt { _ = "STUB: not implemented"; return *new(SolverOpt) }

func WithEps(eps float64) SolverOpt { _ = "STUB: not implemented"; return *new(SolverOpt) }

func WithClip(clip float64) SolverOpt { _ = "STUB: not implemented"; return *new(SolverOpt) }

func WithLearnRate(eta float64) SolverOpt { _ = "STUB: not implemented"; return *new(SolverOpt) }

func WithBeta1(beta1 float64) SolverOpt { _ = "STUB: not implemented"; return *new(SolverOpt) }

func WithBeta2(beta2 float64) SolverOpt { _ = "STUB: not implemented"; return *new(SolverOpt) }

func WithRho(rho float64) SolverOpt { _ = "STUB: not implemented"; return *new(SolverOpt) }

func WithMomentum(momentum float64) SolverOpt { _ = "STUB: not implemented"; return *new(SolverOpt) }

type RMSPropSolver struct {
	decay float64
	eps   float64
	l2reg float64
	clip  float64
	eta   float64

	useClip, useL2Reg bool

	cache []*dualValue
}

func NewRMSPropSolver(opts ...SolverOpt) *RMSPropSolver { _ = "STUB: not implemented"; return nil }

func (s *RMSPropSolver) Step(model []ValueGrad) (err error) { _ = "STUB: not implemented"; return nil }

type AdamSolver struct {
	eta   float64
	eps   float64
	beta1 float64
	beta2 float64
	clip  float64
	l1reg float64
	l2reg float64
	batch float64

	useClip, useL1Reg, useL2Reg bool

	iter  int
	cache []*dualValue
}

func NewAdamSolver(opts ...SolverOpt) *AdamSolver { _ = "STUB: not implemented"; return nil }

func (s *AdamSolver) Step(model []ValueGrad) (err error) { _ = "STUB: not implemented"; return nil }

type VanillaSolver struct {
	eta   float64
	clip  float64
	l1reg float64
	l2reg float64
	batch float64

	useClip, useL1Reg, useL2Reg bool
}

func NewVanillaSolver(opts ...SolverOpt) *VanillaSolver { _ = "STUB: not implemented"; return nil }

func (s *VanillaSolver) Step(model []ValueGrad) (err error) { _ = "STUB: not implemented"; return nil }

type Momentum struct {
	eta      float64
	momentum float64
	clip     float64
	l1reg    float64
	l2reg    float64
	batch    float64

	useClip, useL1Reg, useL2Reg bool

	cache []*dualValue
}

func NewMomentum(opts ...SolverOpt) *Momentum { _ = "STUB: not implemented"; return nil }

func (s *Momentum) Step(model []ValueGrad) (err error) { _ = "STUB: not implemented"; return nil }

type AdaGradSolver struct {
	eta   float64
	eps   float64
	l1Reg float64
	l2reg float64
	clip  float64

	useL2Reg, useClip bool

	cache []*dualValue
}

func NewAdaGradSolver(opts ...SolverOpt) *AdaGradSolver { _ = "STUB: not implemented"; return nil }

func (s *AdaGradSolver) Step(model []ValueGrad) (err error) { _ = "STUB: not implemented"; return nil }

type BarzilaiBorweinSolver struct {
	eta     float64
	clip    float64
	useClip bool
	prevDV  []*dualValue
}

func NewBarzilaiBorweinSolver(opts ...SolverOpt) *BarzilaiBorweinSolver {
	_ = "STUB: not implemented"
	return nil
}

func (s *BarzilaiBorweinSolver) Step(model []ValueGrad) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type adamwState struct {
	expMA   tensor.Tensor
	expMASq tensor.Tensor
	denom   tensor.Tensor
}

type AdamW struct {
	η     float64
	ε     float64
	λ     float64
	ɛ     float64
	β1    float64
	β2    float64
	clip  float64
	l1reg float64
	l2reg float64
	batch float64

	useL1Reg, useL2Reg, useClip bool

	iter   float64
	states map[*Node]*adamwState
}

func NewAdamW(opts ...SolverOpt) *AdamW { _ = "STUB: not implemented"; return nil }

func (a *AdamW) Step(model []ValueGrad) (err error) { _ = "STUB: not implemented"; return nil }
