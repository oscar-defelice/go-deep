package deep

import "math"

type Mode int

const (
	ModeDefault    Mode = 0 // Use default output layer activations (i.e. sigmoid, tanh, relu)
	ModeMulti      Mode = 1 // Softmax output layer
	ModeRegression Mode = 2 // Linear output layer
)

func OutputActivation(c Mode) ActivationType {
	switch c {
	case ModeMulti:
		return ActivationSoftmax
	case ModeRegression:
		return ActivationLinear
	}
	return ActivationNone
}

func Act(act ActivationType) Activation {
	switch act {
	case ActivationSigmoid:
		return Sigmoid{}
	case ActivationTanh:
		return Tanh{}
	case ActivationReLU:
		return ReLU{}
	case ActivationLinear:
		return Linear{}
	case ActivationSoftmax:
		return Linear{}
	}
	return Linear{}
}

type ActivationType int

const (
	ActivationNone    ActivationType = 0
	ActivationSigmoid ActivationType = 1
	ActivationTanh    ActivationType = 2
	ActivationReLU    ActivationType = 3
	ActivationLinear  ActivationType = 4
	ActivationSoftmax ActivationType = 5
)

type Activation interface {
	f(float64) float64
	df(float64) float64
}

type Sigmoid struct{}

func (a Sigmoid) f(x float64) float64  { return Logistic(x, 1.0) }
func (a Sigmoid) df(y float64) float64 { return y * (1.0 - y) }

func Logistic(x, a float64) float64 {
	return 1.0 / (1.0 + math.Exp(-a*x))
}

type Tanh struct{}

func (a Tanh) f(x float64) float64  { return (1 - math.Exp(-2*x)) / (1 + math.Exp(-2*x)) }
func (a Tanh) df(y float64) float64 { return 1 - math.Pow(y, 2) }

type ReLU struct{}

func (a ReLU) f(x float64) float64 { return math.Max(x, 0) }
func (a ReLU) df(y float64) float64 {
	if y > 0 {
		return 1
	}
	return 0
}

type Linear struct{}

func (a Linear) f(x float64) float64  { return x }
func (a Linear) df(x float64) float64 { return 1 }
