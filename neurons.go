package neuralnetworking

import "math"

type Neuron struct {
  Weights []float64
  Bias float64
}

func activationFuction (z float64) float64 {
  //return 1 / (1 + math.Pow(math.E, -1 * z))
  //return math.Tanh(z)
  return ((1.0 / (1.0 + math.Pow(math.E, -1.0 * z)))*2.0) - 1
}

func (n Neuron) CalculateOutput (previousLayer []float64) float64 {
  if (len(n.Weights) != len(previousLayer)) {
    panic("The amount of weights is not equal to the previous layer.")
  }
  acum := 0.0
  for i := 0; i < len(n.Weights); i++ {
    acum += n.Weights[i] * previousLayer[i]
  }
  acum += n.Bias
  return activationFuction(acum)
}

func BuildNeuron (previousLayer int) Neuron {
  weights := make([]float64, previousLayer)
  return Neuron{weights, 0.0}
}
