package neuralnetworking

type Network struct {
  Layers []Layer
}

func (n Network) CalculateOutput (initialLayer []float64) []float64 {
  previousLayer := initialLayer
  for i := 0; i < len(n.Layers); i++ {
    previousLayer = n.Layers[i].CalculateOutput(previousLayer)
  }
  return previousLayer
}

func BuildNetwork (descriptor []int) Network {
  var layers [len(descriptor) - 1]Layer
  for i := 1; i < len(descriptor); i++ {
    layers[i-1] = BuildLayer(descriptor[i-1], descriptor[i])
  }
  return Network{layers}
}
