package main

import "fmt"

func main() {
  xs := []float64{98,93,77,82,83}
  fmt.Println(average(xs))
}

func average(input []float64) float64 {
  var sum float64 = 0.0
  for _, i := range input {
    sum += i
  }
  return sum / float64(len(input))
}
