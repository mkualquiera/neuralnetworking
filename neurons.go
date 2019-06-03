package neuralnetworking

import "math"

type Neuron struct {
  Weights []float64
  Bias float64
}

func activationFuction (z float64) float64 {
  return 1 / (1 + math.Pow(math.E, -1 * z))
}

func (n Neuron) CalculateOutput (previousLayer []float64) float64 {
  if (len(n.Weights) != len(previousLayer)) {
    panic(println("The amount of weights is not equal to the previous layer."))
  }
  acum := 0.0
  for i := 0; i < len(n.Weights); i++ {
    acum += n.Weights[i] * previousLayer[i]
  }
  acum += n.Bias
  return activationFuction(acum)
}

func BuildNeuron (previousLayer int) Neuron {
  var weights [previousLayer]int
  return Neuron{weights, 0.0}
}
