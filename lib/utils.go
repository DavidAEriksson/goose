package lib

import (
  "math/rand"
)

// Generates a random state sting
func GenerateState () string {
  var state string
  for i := 0; i < 32; i++ {
    state += string(rand.Intn(26) + 97)
  }
  return state
}
