package neuralnetworking

import "math/rand"

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

func (n *Network) Randomize () {
  for layerIndex := 0; layerIndex < len(n.Layers); layerIndex++ {
    for neuronIndex := 0; neuronIndex < len(n.Layers[layerIndex].Neurons); neuronIndex++ {
      for weightIndex := 0; weightIndex < len(n.Layers[layerIndex].Neurons[neuronIndex].Weights); weightIndex++ {
        n.Layers[layerIndex].Neurons[neuronIndex].Weights[weightIndex] = (rand.Float64() * 2) - 1
      }
      n.Layers[layerIndex].Neurons[neuronIndex].Bias = (rand.Float64() * 2) - 1
    }
  }
}

func (n *Network) Mutate (biasMutationFactor float64, biasMutationProbability float64, weightMutationFactor float64, weightMutationProbability float64) {
  for layerIndex := 0; layerIndex < len(n.Layers); layerIndex++ {
    for neuronIndex := 0; neuronIndex < len(n.Layers[layerIndex].Neurons); neuronIndex++ {
      for weightIndex := 0; weightIndex < len(n.Layers[layerIndex].Neurons[neuronIndex].Weights); weightIndex++ {
        if (rand.Float64() < weightMutationProbability) {
          mutation := ((rand.Float64() * 2) - 1) * weightMutationFactor
          n.Layers[layerIndex].Neurons[neuronIndex].Weights[weightIndex] += mutation
        }
      }
      if (rand.Float64() < biasMutationProbability) {
        mutation := ((rand.Float64() * 2) - 1) * biasMutationFactor
        n.Layers[layerIndex].Neurons[neuronIndex].Bias += mutation
      }
    }
  }
}

func BuildRandomNetwork (descriptor []int) Network {
  network := BuildNetwork(descriptor)
  network.Randomize()
  return network
}

func BuildNetwork (descriptor []int) Network {
  layers := make([]Layer, len(descriptor) - 1)
  for i := 1; i < len(descriptor); i++ {
    layers[i-1] = BuildLayer(descriptor[i-1], descriptor[i])
  }
  return Network{layers}
}
