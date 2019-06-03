package neuralnetworking

type Layer struct {
  Neurons []Neuron
}

func (l Layer) CalculateOutput (previousLayer []float64) []float64 {
  var result [len(l.Neurons)]float64
  for i := 0; i < len(l.Neurons); i++ {
    result[i] = l.Neurons[i].CalculateOutput(previousLayer)
  }
  return result
}

func BuildLayer (previousLayer int, thisLayer int) Layer {
  var neurons [thisLayer]Neurons
  for i := 0; i < len(neurons); i++ {
    neurons[i] = BuildNeuron(previousLayer)
  }
  return Layer{neurons}
}
