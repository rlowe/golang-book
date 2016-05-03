package main

import "fmt"

func main() {
  x := [5]float64{ 98, 93, 77, 82, 83, }

  var sum float64 = 0
  for _, i := range x {
    sum += i
  }

  fmt.Println("AVG: ", sum / float64(len(x)))
}
