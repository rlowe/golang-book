package main

import "fmt"

func main() {
  x := []int{ 48,96,86,68, 57,82,63,70, 37,34,83,27, 19,97, 9,17, }

  var smallest int
  firstrun := true

  for _, i := range x {
    if firstrun == true {
      smallest = i
      firstrun = false
    } else if i < smallest {
      smallest = i
    }
  }

  fmt.Println(smallest)
}
