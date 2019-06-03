package neuralnetworking

type Layer struct {
  Neurons []Neuron
}

func (l Layer) CalculateOutput (previousLayer []float64) []float64 {
  result := make([]float64, len(l.Neurons))
  for i := 0; i < len(l.Neurons); i++ {
    result[i] = l.Neurons[i].CalculateOutput(previousLayer)
  }
  return result
}

func BuildLayer (previousLayer int, thisLayer int) Layer {
  neurons := make([]Neuron, thisLayer)
  for i := 0; i < len(neurons); i++ {
    neurons[i] = BuildNeuron(previousLayer)
  }
  return Layer{neurons}
}
